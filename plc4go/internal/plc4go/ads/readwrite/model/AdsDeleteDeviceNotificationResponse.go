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
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type AdsDeleteDeviceNotificationResponse struct {
	Result ReturnCode
	Parent *AdsData
}

// The corresponding interface
type IAdsDeleteDeviceNotificationResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsDeleteDeviceNotificationResponse) CommandId() CommandId {
	return CommandId_ADS_DELETE_DEVICE_NOTIFICATION
}

func (m *AdsDeleteDeviceNotificationResponse) Response() bool {
	return true
}

func (m *AdsDeleteDeviceNotificationResponse) InitializeParent(parent *AdsData) {
}

func NewAdsDeleteDeviceNotificationResponse(result ReturnCode) *AdsData {
	child := &AdsDeleteDeviceNotificationResponse{
		Result: result,
		Parent: NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsDeleteDeviceNotificationResponse(structType interface{}) *AdsDeleteDeviceNotificationResponse {
	castFunc := func(typ interface{}) *AdsDeleteDeviceNotificationResponse {
		if casted, ok := typ.(AdsDeleteDeviceNotificationResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsDeleteDeviceNotificationResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsDeleteDeviceNotificationResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsDeleteDeviceNotificationResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsDeleteDeviceNotificationResponse) GetTypeName() string {
	return "AdsDeleteDeviceNotificationResponse"
}

func (m *AdsDeleteDeviceNotificationResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsDeleteDeviceNotificationResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsDeleteDeviceNotificationResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsDeleteDeviceNotificationResponseParse(readBuffer utils.ReadBuffer) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsDeleteDeviceNotificationResponse"); pullErr != nil {
		return nil, pullErr
	}

	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (result)
	result, _resultErr := ReturnCodeParse(readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AdsDeleteDeviceNotificationResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsDeleteDeviceNotificationResponse{
		Result: result,
		Parent: &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsDeleteDeviceNotificationResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDeleteDeviceNotificationResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return pushErr
		}
		_resultErr := m.Result.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return popErr
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		if popErr := writeBuffer.PopContext("AdsDeleteDeviceNotificationResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsDeleteDeviceNotificationResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
