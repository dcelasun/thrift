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
)

// Memory buffer-based implementation of the TTransport interface.
type TMemoryBuffer struct {
	*TBaseTransport
	*bytes.Buffer
	size int
}

type TMemoryBufferTransportFactory struct {
	size int

	config *TConfiguration
}

func (p *TMemoryBufferTransportFactory) GetTransport(trans TTransport) (TTransport, error) {
	if trans != nil {
		t, ok := trans.(*TMemoryBuffer)
		if ok && t.size > 0 {
			return NewTMemoryBufferLenWithConfiguration(t.size, p.config), nil
		}
	}
	return NewTMemoryBufferLenWithConfiguration(p.size, p.config), nil
}

func (p *TMemoryBufferTransportFactory) GetConfiguration() *TConfiguration {
	return p.config
}

func NewTMemoryBufferTransportFactory(size int) *TMemoryBufferTransportFactory {
	return NewTMemoryBufferTransportFactoryWithConfiguration(size, defaultConfiguration)
}

func NewTMemoryBufferTransportFactoryWithConfiguration(size int, config *TConfiguration) *TMemoryBufferTransportFactory {
	return &TMemoryBufferTransportFactory{size: size, config: config}
}

func NewTMemoryBuffer() *TMemoryBuffer {
	return &TMemoryBuffer{Buffer: &bytes.Buffer{}, size: 0}
}

func NewTMemoryBufferLen(size int) *TMemoryBuffer {
	return NewTMemoryBufferLenWithConfiguration(size, defaultConfiguration)
}

func NewTMemoryBufferLenWithConfiguration(size int, config *TConfiguration) *TMemoryBuffer {
	buf := make([]byte, 0, size)
	return &TMemoryBuffer{TBaseTransport: NewTBaseTransport(config), Buffer: bytes.NewBuffer(buf), size: size}
}

func (p *TMemoryBuffer) IsOpen() bool {
	return true
}

func (p *TMemoryBuffer) Open() error {
	return nil
}

func (p *TMemoryBuffer) Close() error {
	p.Buffer.Reset()
	return nil
}

// Flushing a memory buffer is a no-op
func (p *TMemoryBuffer) Flush(ctx context.Context) error {
	return nil
}

func (p *TMemoryBuffer) RemainingBytes() (num_bytes uint64) {
	return uint64(p.Buffer.Len())
}
