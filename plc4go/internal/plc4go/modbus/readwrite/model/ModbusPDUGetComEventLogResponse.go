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
type ModbusPDUGetComEventLogResponse struct {
	Status       uint16
	EventCount   uint16
	MessageCount uint16
	Events       []int8
	Parent       *ModbusPDU
}

// The corresponding interface
type IModbusPDUGetComEventLogResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUGetComEventLogResponse) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUGetComEventLogResponse) FunctionFlag() uint8 {
	return 0x0C
}

func (m *ModbusPDUGetComEventLogResponse) Response() bool {
	return true
}

func (m *ModbusPDUGetComEventLogResponse) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUGetComEventLogResponse(status uint16, eventCount uint16, messageCount uint16, events []int8) *ModbusPDU {
	child := &ModbusPDUGetComEventLogResponse{
		Status:       status,
		EventCount:   eventCount,
		MessageCount: messageCount,
		Events:       events,
		Parent:       NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUGetComEventLogResponse(structType interface{}) *ModbusPDUGetComEventLogResponse {
	castFunc := func(typ interface{}) *ModbusPDUGetComEventLogResponse {
		if casted, ok := typ.(ModbusPDUGetComEventLogResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUGetComEventLogResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUGetComEventLogResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUGetComEventLogResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUGetComEventLogResponse) GetTypeName() string {
	return "ModbusPDUGetComEventLogResponse"
}

func (m *ModbusPDUGetComEventLogResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUGetComEventLogResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 16

	// Simple field (eventCount)
	lengthInBits += 16

	// Simple field (messageCount)
	lengthInBits += 16

	// Array field
	if len(m.Events) > 0 {
		lengthInBits += 8 * uint16(len(m.Events))
	}

	return lengthInBits
}

func (m *ModbusPDUGetComEventLogResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUGetComEventLogResponseParse(readBuffer utils.ReadBuffer) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventLogResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Simple Field (status)
	status, _statusErr := readBuffer.ReadUint16("status", 16)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}

	// Simple Field (eventCount)
	eventCount, _eventCountErr := readBuffer.ReadUint16("eventCount", 16)
	if _eventCountErr != nil {
		return nil, errors.Wrap(_eventCountErr, "Error parsing 'eventCount' field")
	}

	// Simple Field (messageCount)
	messageCount, _messageCountErr := readBuffer.ReadUint16("messageCount", 16)
	if _messageCountErr != nil {
		return nil, errors.Wrap(_messageCountErr, "Error parsing 'messageCount' field")
	}

	// Array field (events)
	if pullErr := readBuffer.PullContext("events", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	events := make([]int8, uint16(byteCount)-uint16(uint16(6)))
	for curItem := uint16(0); curItem < uint16(uint16(byteCount)-uint16(uint16(6))); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'events' field")
		}
		events[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("events", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventLogResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUGetComEventLogResponse{
		Status:       status,
		EventCount:   eventCount,
		MessageCount: messageCount,
		Events:       events,
		Parent:       &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUGetComEventLogResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventLogResponse"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(uint8(len(m.Events))) + uint8(uint8(6)))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Simple Field (status)
		status := uint16(m.Status)
		_statusErr := writeBuffer.WriteUint16("status", 16, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (eventCount)
		eventCount := uint16(m.EventCount)
		_eventCountErr := writeBuffer.WriteUint16("eventCount", 16, (eventCount))
		if _eventCountErr != nil {
			return errors.Wrap(_eventCountErr, "Error serializing 'eventCount' field")
		}

		// Simple Field (messageCount)
		messageCount := uint16(m.MessageCount)
		_messageCountErr := writeBuffer.WriteUint16("messageCount", 16, (messageCount))
		if _messageCountErr != nil {
			return errors.Wrap(_messageCountErr, "Error serializing 'messageCount' field")
		}

		// Array Field (events)
		if m.Events != nil {
			if pushErr := writeBuffer.PushContext("events", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Events {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'events' field")
				}
			}
			if popErr := writeBuffer.PopContext("events", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventLogResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUGetComEventLogResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
