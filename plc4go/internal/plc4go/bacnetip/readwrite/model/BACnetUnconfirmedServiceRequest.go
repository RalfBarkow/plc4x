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
type BACnetUnconfirmedServiceRequest struct {
	Child IBACnetUnconfirmedServiceRequestChild
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequest interface {
	ServiceChoice() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetUnconfirmedServiceRequestParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetUnconfirmedServiceRequest, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetUnconfirmedServiceRequestChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetUnconfirmedServiceRequest)
	GetTypeName() string
	IBACnetUnconfirmedServiceRequest
}

func NewBACnetUnconfirmedServiceRequest() *BACnetUnconfirmedServiceRequest {
	return &BACnetUnconfirmedServiceRequest{}
}

func CastBACnetUnconfirmedServiceRequest(structType interface{}) *BACnetUnconfirmedServiceRequest {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequest {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequest) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequest"
}

func (m *BACnetUnconfirmedServiceRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequest) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetUnconfirmedServiceRequest) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice, _serviceChoiceErr := readBuffer.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetUnconfirmedServiceRequest
	var typeSwitchError error
	switch {
	case serviceChoice == 0x00: // BACnetUnconfirmedServiceRequestIAm
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestIAmParse(readBuffer)
	case serviceChoice == 0x01: // BACnetUnconfirmedServiceRequestIHave
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestIHaveParse(readBuffer)
	case serviceChoice == 0x02: // BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationParse(readBuffer)
	case serviceChoice == 0x03: // BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationParse(readBuffer)
	case serviceChoice == 0x04: // BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferParse(readBuffer, len)
	case serviceChoice == 0x05: // BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedTextMessageParse(readBuffer)
	case serviceChoice == 0x06: // BACnetUnconfirmedServiceRequestTimeSynchronization
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestTimeSynchronizationParse(readBuffer)
	case serviceChoice == 0x07: // BACnetUnconfirmedServiceRequestWhoHas
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestWhoHasParse(readBuffer)
	case serviceChoice == 0x08: // BACnetUnconfirmedServiceRequestWhoIs
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestWhoIsParse(readBuffer)
	case serviceChoice == 0x09: // BACnetUnconfirmedServiceRequestUTCTimeSynchronization
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUTCTimeSynchronizationParse(readBuffer)
	case serviceChoice == 0x0A: // BACnetUnconfirmedServiceRequestWriteGroup
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestWriteGroupParse(readBuffer)
	case serviceChoice == 0x0B: // BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *BACnetUnconfirmedServiceRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetUnconfirmedServiceRequest) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetUnconfirmedServiceRequest, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequest"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := uint8(child.ServiceChoice())
	_serviceChoiceErr := writeBuffer.WriteUint8("serviceChoice", 8, (serviceChoice))

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequest"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
