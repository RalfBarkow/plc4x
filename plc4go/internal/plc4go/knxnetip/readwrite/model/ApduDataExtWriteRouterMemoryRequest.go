//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ApduDataExtWriteRouterMemoryRequest struct {
	Parent *ApduDataExt
}

// The corresponding interface
type IApduDataExtWriteRouterMemoryRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtWriteRouterMemoryRequest) ExtApciType() uint8 {
	return 0x0A
}

func (m *ApduDataExtWriteRouterMemoryRequest) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtWriteRouterMemoryRequest() *ApduDataExt {
	child := &ApduDataExtWriteRouterMemoryRequest{
		Parent: NewApduDataExt(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataExtWriteRouterMemoryRequest(structType interface{}) *ApduDataExtWriteRouterMemoryRequest {
	castFunc := func(typ interface{}) *ApduDataExtWriteRouterMemoryRequest {
		if casted, ok := typ.(ApduDataExtWriteRouterMemoryRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtWriteRouterMemoryRequest); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtWriteRouterMemoryRequest(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtWriteRouterMemoryRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtWriteRouterMemoryRequest) GetTypeName() string {
	return "ApduDataExtWriteRouterMemoryRequest"
}

func (m *ApduDataExtWriteRouterMemoryRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataExtWriteRouterMemoryRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtWriteRouterMemoryRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtWriteRouterMemoryRequestParse(readBuffer utils.ReadBuffer) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtWriteRouterMemoryRequest"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtWriteRouterMemoryRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtWriteRouterMemoryRequest{
		Parent: &ApduDataExt{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataExtWriteRouterMemoryRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtWriteRouterMemoryRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtWriteRouterMemoryRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
