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

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AmsPacket struct {
	TargetAmsNetId *AmsNetId
	TargetAmsPort  uint16
	SourceAmsNetId *AmsNetId
	SourceAmsPort  uint16
	CommandId      CommandId
	State          *State
	ErrorCode      uint32
	InvokeId       uint32
	Data           *AdsData
}

// The corresponding interface
type IAmsPacket interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewAmsPacket(targetAmsNetId *AmsNetId, targetAmsPort uint16, sourceAmsNetId *AmsNetId, sourceAmsPort uint16, commandId CommandId, state *State, errorCode uint32, invokeId uint32, data *AdsData) *AmsPacket {
	return &AmsPacket{TargetAmsNetId: targetAmsNetId, TargetAmsPort: targetAmsPort, SourceAmsNetId: sourceAmsNetId, SourceAmsPort: sourceAmsPort, CommandId: commandId, State: state, ErrorCode: errorCode, InvokeId: invokeId, Data: data}
}

func CastAmsPacket(structType interface{}) *AmsPacket {
	castFunc := func(typ interface{}) *AmsPacket {
		if casted, ok := typ.(AmsPacket); ok {
			return &casted
		}
		if casted, ok := typ.(*AmsPacket); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AmsPacket) GetTypeName() string {
	return "AmsPacket"
}

func (m *AmsPacket) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AmsPacket) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (targetAmsNetId)
	lengthInBits += m.TargetAmsNetId.LengthInBits()

	// Simple field (targetAmsPort)
	lengthInBits += 16

	// Simple field (sourceAmsNetId)
	lengthInBits += m.SourceAmsNetId.LengthInBits()

	// Simple field (sourceAmsPort)
	lengthInBits += 16

	// Simple field (commandId)
	lengthInBits += 16

	// Simple field (state)
	lengthInBits += m.State.LengthInBits()

	// Implicit Field (length)
	lengthInBits += 32

	// Simple field (errorCode)
	lengthInBits += 32

	// Simple field (invokeId)
	lengthInBits += 32

	// Simple field (data)
	lengthInBits += m.Data.LengthInBits()

	return lengthInBits
}

func (m *AmsPacket) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AmsPacketParse(readBuffer utils.ReadBuffer) (*AmsPacket, error) {
	if pullErr := readBuffer.PullContext("AmsPacket"); pullErr != nil {
		return nil, pullErr
	}

	if pullErr := readBuffer.PullContext("targetAmsNetId"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (targetAmsNetId)
	targetAmsNetId, _targetAmsNetIdErr := AmsNetIdParse(readBuffer)
	if _targetAmsNetIdErr != nil {
		return nil, errors.Wrap(_targetAmsNetIdErr, "Error parsing 'targetAmsNetId' field")
	}
	if closeErr := readBuffer.CloseContext("targetAmsNetId"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (targetAmsPort)
	targetAmsPort, _targetAmsPortErr := readBuffer.ReadUint16("targetAmsPort", 16)
	if _targetAmsPortErr != nil {
		return nil, errors.Wrap(_targetAmsPortErr, "Error parsing 'targetAmsPort' field")
	}

	if pullErr := readBuffer.PullContext("sourceAmsNetId"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (sourceAmsNetId)
	sourceAmsNetId, _sourceAmsNetIdErr := AmsNetIdParse(readBuffer)
	if _sourceAmsNetIdErr != nil {
		return nil, errors.Wrap(_sourceAmsNetIdErr, "Error parsing 'sourceAmsNetId' field")
	}
	if closeErr := readBuffer.CloseContext("sourceAmsNetId"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (sourceAmsPort)
	sourceAmsPort, _sourceAmsPortErr := readBuffer.ReadUint16("sourceAmsPort", 16)
	if _sourceAmsPortErr != nil {
		return nil, errors.Wrap(_sourceAmsPortErr, "Error parsing 'sourceAmsPort' field")
	}

	if pullErr := readBuffer.PullContext("commandId"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (commandId)
	commandId, _commandIdErr := CommandIdParse(readBuffer)
	if _commandIdErr != nil {
		return nil, errors.Wrap(_commandIdErr, "Error parsing 'commandId' field")
	}
	if closeErr := readBuffer.CloseContext("commandId"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("state"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (state)
	state, _stateErr := StateParse(readBuffer)
	if _stateErr != nil {
		return nil, errors.Wrap(_stateErr, "Error parsing 'state' field")
	}
	if closeErr := readBuffer.CloseContext("state"); closeErr != nil {
		return nil, closeErr
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := readBuffer.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	// Simple Field (errorCode)
	errorCode, _errorCodeErr := readBuffer.ReadUint32("errorCode", 32)
	if _errorCodeErr != nil {
		return nil, errors.Wrap(_errorCodeErr, "Error parsing 'errorCode' field")
	}

	// Simple Field (invokeId)
	invokeId, _invokeIdErr := readBuffer.ReadUint32("invokeId", 32)
	if _invokeIdErr != nil {
		return nil, errors.Wrap(_invokeIdErr, "Error parsing 'invokeId' field")
	}

	if pullErr := readBuffer.PullContext("data"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (data)
	data, _dataErr := AdsDataParse(readBuffer, &commandId, state.Response)
	if _dataErr != nil {
		return nil, errors.Wrap(_dataErr, "Error parsing 'data' field")
	}
	if closeErr := readBuffer.CloseContext("data"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AmsPacket"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, commandId, state, errorCode, invokeId, data), nil
}

func (m *AmsPacket) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AmsPacket"); pushErr != nil {
		return pushErr
	}

	// Simple Field (targetAmsNetId)
	if pushErr := writeBuffer.PushContext("targetAmsNetId"); pushErr != nil {
		return pushErr
	}
	_targetAmsNetIdErr := m.TargetAmsNetId.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("targetAmsNetId"); popErr != nil {
		return popErr
	}
	if _targetAmsNetIdErr != nil {
		return errors.Wrap(_targetAmsNetIdErr, "Error serializing 'targetAmsNetId' field")
	}

	// Simple Field (targetAmsPort)
	targetAmsPort := uint16(m.TargetAmsPort)
	_targetAmsPortErr := writeBuffer.WriteUint16("targetAmsPort", 16, (targetAmsPort))
	if _targetAmsPortErr != nil {
		return errors.Wrap(_targetAmsPortErr, "Error serializing 'targetAmsPort' field")
	}

	// Simple Field (sourceAmsNetId)
	if pushErr := writeBuffer.PushContext("sourceAmsNetId"); pushErr != nil {
		return pushErr
	}
	_sourceAmsNetIdErr := m.SourceAmsNetId.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("sourceAmsNetId"); popErr != nil {
		return popErr
	}
	if _sourceAmsNetIdErr != nil {
		return errors.Wrap(_sourceAmsNetIdErr, "Error serializing 'sourceAmsNetId' field")
	}

	// Simple Field (sourceAmsPort)
	sourceAmsPort := uint16(m.SourceAmsPort)
	_sourceAmsPortErr := writeBuffer.WriteUint16("sourceAmsPort", 16, (sourceAmsPort))
	if _sourceAmsPortErr != nil {
		return errors.Wrap(_sourceAmsPortErr, "Error serializing 'sourceAmsPort' field")
	}

	// Simple Field (commandId)
	if pushErr := writeBuffer.PushContext("commandId"); pushErr != nil {
		return pushErr
	}
	_commandIdErr := m.CommandId.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("commandId"); popErr != nil {
		return popErr
	}
	if _commandIdErr != nil {
		return errors.Wrap(_commandIdErr, "Error serializing 'commandId' field")
	}

	// Simple Field (state)
	if pushErr := writeBuffer.PushContext("state"); pushErr != nil {
		return pushErr
	}
	_stateErr := m.State.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("state"); popErr != nil {
		return popErr
	}
	if _stateErr != nil {
		return errors.Wrap(_stateErr, "Error serializing 'state' field")
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length := uint32(m.Data.LengthInBytes())
	_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (errorCode)
	errorCode := uint32(m.ErrorCode)
	_errorCodeErr := writeBuffer.WriteUint32("errorCode", 32, (errorCode))
	if _errorCodeErr != nil {
		return errors.Wrap(_errorCodeErr, "Error serializing 'errorCode' field")
	}

	// Simple Field (invokeId)
	invokeId := uint32(m.InvokeId)
	_invokeIdErr := writeBuffer.WriteUint32("invokeId", 32, (invokeId))
	if _invokeIdErr != nil {
		return errors.Wrap(_invokeIdErr, "Error serializing 'invokeId' field")
	}

	// Simple Field (data)
	if pushErr := writeBuffer.PushContext("data"); pushErr != nil {
		return pushErr
	}
	_dataErr := m.Data.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("data"); popErr != nil {
		return popErr
	}
	if _dataErr != nil {
		return errors.Wrap(_dataErr, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("AmsPacket"); popErr != nil {
		return popErr
	}
	return nil
}
