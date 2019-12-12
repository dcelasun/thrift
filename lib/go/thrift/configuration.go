package thrift

import "net/http"

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

var (
	defaultConfiguration = &TConfiguration{
		recursionDepth: 64,
		maxFrameLength: 16384000,
		maxMessageSize: 100 * 1 << 20, // 100MB
		httpClient:     http.DefaultClient,
	}
)

// Holds configuration across protocols and transports.
type TConfiguration struct {
	// The maximum recursive depth protocol.go/Skip() will traverse.
	recursionDepth int
	// Maximum frame length in bytes for TFramedTransport.
	maxFrameLength uint32
	// Maximum message size in bytes. Requests with larger payloads will be rejected.
	maxMessageSize int64
	// HTTP client to be used.
	httpClient *http.Client
}

func NewTConfiguration() *TConfiguration {
	tc := *defaultConfiguration
	return &tc
}

func (c *TConfiguration) WithRecursionDepth(n int) *TConfiguration {
	c.recursionDepth = n
	return c
}

func (c *TConfiguration) WithMaxFrameLength(n uint32) *TConfiguration {
	c.maxFrameLength = n
	return c
}

func (c *TConfiguration) WithMaxMessageSize(n int64) *TConfiguration {
	c.maxMessageSize = n
	return c
}

func (c *TConfiguration) WithHttpClient(h *http.Client) *TConfiguration {
	c.httpClient = h
	return c
}
