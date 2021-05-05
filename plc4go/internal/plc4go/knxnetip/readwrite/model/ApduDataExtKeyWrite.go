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
type ApduDataExtKeyWrite struct {
	Parent *ApduDataExt
}

// The corresponding interface
type IApduDataExtKeyWrite interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtKeyWrite) ExtApciType() uint8 {
	return 0x13
}

func (m *ApduDataExtKeyWrite) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtKeyWrite() *ApduDataExt {
	child := &ApduDataExtKeyWrite{
		Parent: NewApduDataExt(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataExtKeyWrite(structType interface{}) *ApduDataExtKeyWrite {
	castFunc := func(typ interface{}) *ApduDataExtKeyWrite {
		if casted, ok := typ.(ApduDataExtKeyWrite); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtKeyWrite); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtKeyWrite(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtKeyWrite(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtKeyWrite) GetTypeName() string {
	return "ApduDataExtKeyWrite"
}

func (m *ApduDataExtKeyWrite) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataExtKeyWrite) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtKeyWrite) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtKeyWriteParse(readBuffer utils.ReadBuffer) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtKeyWrite"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtKeyWrite"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtKeyWrite{
		Parent: &ApduDataExt{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataExtKeyWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtKeyWrite"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtKeyWrite"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
