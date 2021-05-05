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
type ApduDataExtGroupPropertyValueInfoReport struct {
	Parent *ApduDataExt
}

// The corresponding interface
type IApduDataExtGroupPropertyValueInfoReport interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtGroupPropertyValueInfoReport) ExtApciType() uint8 {
	return 0x2B
}

func (m *ApduDataExtGroupPropertyValueInfoReport) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtGroupPropertyValueInfoReport() *ApduDataExt {
	child := &ApduDataExtGroupPropertyValueInfoReport{
		Parent: NewApduDataExt(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataExtGroupPropertyValueInfoReport(structType interface{}) *ApduDataExtGroupPropertyValueInfoReport {
	castFunc := func(typ interface{}) *ApduDataExtGroupPropertyValueInfoReport {
		if casted, ok := typ.(ApduDataExtGroupPropertyValueInfoReport); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtGroupPropertyValueInfoReport); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtGroupPropertyValueInfoReport(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtGroupPropertyValueInfoReport(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtGroupPropertyValueInfoReport) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueInfoReport"
}

func (m *ApduDataExtGroupPropertyValueInfoReport) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataExtGroupPropertyValueInfoReport) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtGroupPropertyValueInfoReport) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtGroupPropertyValueInfoReportParse(readBuffer utils.ReadBuffer) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueInfoReport"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueInfoReport"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtGroupPropertyValueInfoReport{
		Parent: &ApduDataExt{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataExtGroupPropertyValueInfoReport) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueInfoReport"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueInfoReport"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
