/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package thrift

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type THttpClient struct {
	*TBaseTransport
	client             *http.Client
	response           *http.Response
	url                *url.URL
	requestBuffer      *bytes.Buffer
	header             http.Header
	nsecConnectTimeout int64
	nsecReadTimeout    int64
}

type THttpClientTransportFactory struct {
	url string

	config *TConfiguration
}

func (p *THttpClientTransportFactory) GetTransport(trans TTransport) (TTransport, error) {
	if trans != nil {
		t, ok := trans.(*THttpClient)
		if ok && t.url != nil {
			return NewTHttpClientWithConfiguration(t.url.String(), p.config)
		}
	}
	return NewTHttpClientWithConfiguration(p.url, p.config)
}

func (p *THttpClientTransportFactory) GetConfiguration() *TConfiguration {
	return p.config
}

// Deprecated: Use TConfiguration instead
type THttpClientOptions struct {
	// If nil, defaultConfiguration.httpClient is used
	Client *http.Client
}

func NewTHttpClientTransportFactory(url string) *THttpClientTransportFactory {
	return NewTHttpClientTransportFactoryWithConfiguration(url, defaultConfiguration)
}

// Deprecated: Use NewTHttpClientTransportFactoryWithConfiguration instead
func NewTHttpClientTransportFactoryWithOptions(url string, options THttpClientOptions) *THttpClientTransportFactory {
	f := NewTHttpClientTransportFactory(url)
	if options.Client != nil {
		f.config.httpClient = options.Client
	}
	return f
}

func NewTHttpClientTransportFactoryWithConfiguration(url string, config *TConfiguration) *THttpClientTransportFactory {
	return &THttpClientTransportFactory{url: url, config: config}
}

func NewTHttpClientWithConfiguration(urlstr string, config *TConfiguration) (TTransport, error) {
	parsedURL, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 0, 1024)
	httpHeader := map[string][]string{"Content-Type": {"application/x-thrift"}}
	return &THttpClient{TBaseTransport: NewTBaseTransport(config), client: config.httpClient, url: parsedURL, requestBuffer: bytes.NewBuffer(buf), header: httpHeader}, nil
}

// Deprecated: Use NewTHttpClientWithConfiguration instead
func NewTHttpClientWithOptions(urlstr string, options THttpClientOptions) (TTransport, error) {
	f, err := NewTHttpClient(urlstr)
	if err != nil {
		return nil, err
	}

	if options.Client != nil {
		f.(*THttpClient).client = options.Client
	}
	return f, nil
}

func NewTHttpClient(urlstr string) (TTransport, error) {
	return NewTHttpClientWithConfiguration(urlstr, defaultConfiguration)
}

// Set the HTTP Header for this specific Thrift Transport
// It is important that you first assert the TTransport as a THttpClient type
// like so:
//
// httpTrans := trans.(THttpClient)
// httpTrans.SetHeader("User-Agent","Thrift Client 1.0")
func (p *THttpClient) SetHeader(key string, value string) {
	p.header.Add(key, value)
}

// Get the HTTP Header represented by the supplied Header Key for this specific Thrift Transport
// It is important that you first assert the TTransport as a THttpClient type
// like so:
//
// httpTrans := trans.(THttpClient)
// hdrValue := httpTrans.GetHeader("User-Agent")
func (p *THttpClient) GetHeader(key string) string {
	return p.header.Get(key)
}

// Deletes the HTTP Header given a Header Key for this specific Thrift Transport
// It is important that you first assert the TTransport as a THttpClient type
// like so:
//
// httpTrans := trans.(THttpClient)
// httpTrans.DelHeader("User-Agent")
func (p *THttpClient) DelHeader(key string) {
	p.header.Del(key)
}

func (p *THttpClient) Open() error {
	// do nothing
	return nil
}

func (p *THttpClient) IsOpen() bool {
	return p.response != nil || p.requestBuffer != nil
}

func (p *THttpClient) closeResponse() error {
	var err error
	if p.response != nil && p.response.Body != nil {
		// The docs specify that if keepalive is enabled and the response body is not
		// read to completion the connection will never be returned to the pool and
		// reused. Errors are being ignored here because if the connection is invalid
		// and this fails for some reason, the Close() method will do any remaining
		// cleanup.
		io.Copy(ioutil.Discard, p.response.Body)

		err = p.response.Body.Close()
	}

	p.response = nil
	return err
}

func (p *THttpClient) Close() error {
	if p.requestBuffer != nil {
		p.requestBuffer.Reset()
		p.requestBuffer = nil
	}
	return p.closeResponse()
}

func (p *THttpClient) Read(buf []byte) (int, error) {
	if p.response == nil {
		return 0, NewTTransportException(NOT_OPEN, "Response buffer is empty, no request.")
	}
	n, err := p.response.Body.Read(buf)
	if n > 0 && (err == nil || err == io.EOF) {
		return n, nil
	}
	return n, NewTTransportExceptionFromError(err)
}

func (p *THttpClient) ReadByte() (c byte, err error) {
	return readByte(p.response.Body)
}

func (p *THttpClient) Write(buf []byte) (int, error) {
	n, err := p.requestBuffer.Write(buf)
	return n, err
}

func (p *THttpClient) WriteByte(c byte) error {
	return p.requestBuffer.WriteByte(c)
}

func (p *THttpClient) WriteString(s string) (n int, err error) {
	return p.requestBuffer.WriteString(s)
}

func (p *THttpClient) Flush(ctx context.Context) error {
	// Close any previous response body to avoid leaking connections.
	p.closeResponse()

	req, err := http.NewRequest("POST", p.url.String(), p.requestBuffer)
	if err != nil {
		return NewTTransportExceptionFromError(err)
	}
	req.Header = p.header
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	response, err := p.client.Do(req)
	if err != nil {
		return NewTTransportExceptionFromError(err)
	}
	if response.StatusCode != http.StatusOK {
		// Close the response to avoid leaking file descriptors. closeResponse does
		// more than just call Close(), so temporarily assign it and reuse the logic.
		p.response = response
		p.closeResponse()

		// TODO(pomack) log bad response
		return NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, "HTTP Response code: "+strconv.Itoa(response.StatusCode))
	}
	p.response = response
	return nil
}

func (p *THttpClient) RemainingBytes() (num_bytes uint64) {
	len := p.response.ContentLength
	if len >= 0 {
		return uint64(len)
	}

	const maxSize = ^uint64(0)
	return maxSize // the truth is, we just don't know unless framed is used
}
