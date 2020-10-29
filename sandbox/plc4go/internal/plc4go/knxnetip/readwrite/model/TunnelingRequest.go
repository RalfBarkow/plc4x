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
    "encoding/xml"
    "errors"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
    "reflect"
)

// The data-structure of this message
type TunnelingRequest struct {
    TunnelingRequestDataBlock ITunnelingRequestDataBlock
    Cemi ICEMI
    KNXNetIPMessage
}

// The corresponding interface
type ITunnelingRequest interface {
    IKNXNetIPMessage
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m TunnelingRequest) MsgType() uint16 {
    return 0x0420
}

func (m TunnelingRequest) initialize() spi.Message {
    return m
}

func NewTunnelingRequest(tunnelingRequestDataBlock ITunnelingRequestDataBlock, cemi ICEMI) KNXNetIPMessageInitializer {
    return &TunnelingRequest{TunnelingRequestDataBlock: tunnelingRequestDataBlock, Cemi: cemi}
}

func CastITunnelingRequest(structType interface{}) ITunnelingRequest {
    castFunc := func(typ interface{}) ITunnelingRequest {
        if iTunnelingRequest, ok := typ.(ITunnelingRequest); ok {
            return iTunnelingRequest
        }
        return nil
    }
    return castFunc(structType)
}

func CastTunnelingRequest(structType interface{}) TunnelingRequest {
    castFunc := func(typ interface{}) TunnelingRequest {
        if sTunnelingRequest, ok := typ.(TunnelingRequest); ok {
            return sTunnelingRequest
        }
        if sTunnelingRequest, ok := typ.(*TunnelingRequest); ok {
            return *sTunnelingRequest
        }
        return TunnelingRequest{}
    }
    return castFunc(structType)
}

func (m TunnelingRequest) LengthInBits() uint16 {
    var lengthInBits uint16 = m.KNXNetIPMessage.LengthInBits()

    // Simple field (tunnelingRequestDataBlock)
    lengthInBits += m.TunnelingRequestDataBlock.LengthInBits()

    // Simple field (cemi)
    lengthInBits += m.Cemi.LengthInBits()

    return lengthInBits
}

func (m TunnelingRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func TunnelingRequestParse(io *utils.ReadBuffer, totalLength uint16) (KNXNetIPMessageInitializer, error) {

    // Simple Field (tunnelingRequestDataBlock)
    _tunnelingRequestDataBlockMessage, _err := TunnelingRequestDataBlockParse(io)
    if _err != nil {
        return nil, errors.New("Error parsing simple field 'tunnelingRequestDataBlock'. " + _err.Error())
    }
    var tunnelingRequestDataBlock ITunnelingRequestDataBlock
    tunnelingRequestDataBlock, _tunnelingRequestDataBlockOk := _tunnelingRequestDataBlockMessage.(ITunnelingRequestDataBlock)
    if !_tunnelingRequestDataBlockOk {
        return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_tunnelingRequestDataBlockMessage).Name() + " to ITunnelingRequestDataBlock")
    }

    // Simple Field (cemi)
    _cemiMessage, _err := CEMIParse(io, uint8(totalLength) - uint8(uint8(uint8(uint8(6)) + uint8(tunnelingRequestDataBlock.LengthInBytes()))))
    if _err != nil {
        return nil, errors.New("Error parsing simple field 'cemi'. " + _err.Error())
    }
    var cemi ICEMI
    cemi, _cemiOk := _cemiMessage.(ICEMI)
    if !_cemiOk {
        return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_cemiMessage).Name() + " to ICEMI")
    }

    // Create the instance
    return NewTunnelingRequest(tunnelingRequestDataBlock, cemi), nil
}

func (m TunnelingRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (tunnelingRequestDataBlock)
    tunnelingRequestDataBlock := CastITunnelingRequestDataBlock(m.TunnelingRequestDataBlock)
    _tunnelingRequestDataBlockErr := tunnelingRequestDataBlock.Serialize(io)
    if _tunnelingRequestDataBlockErr != nil {
        return errors.New("Error serializing 'tunnelingRequestDataBlock' field " + _tunnelingRequestDataBlockErr.Error())
    }

    // Simple Field (cemi)
    cemi := CastICEMI(m.Cemi)
    _cemiErr := cemi.Serialize(io)
    if _cemiErr != nil {
        return errors.New("Error serializing 'cemi' field " + _cemiErr.Error())
    }

        return nil
    }
    return KNXNetIPMessageSerialize(io, m.KNXNetIPMessage, CastIKNXNetIPMessage(m), ser)
}

func (m *TunnelingRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    for {
        token, err := d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "tunnelingRequestDataBlock":
                var data *TunnelingRequestDataBlock
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.TunnelingRequestDataBlock = CastITunnelingRequestDataBlock(data)
            case "cemi":
                switch tok.Attr[0].Value {
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIDataReq":
                        var dt *CEMIDataReq
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.Cemi = dt
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIDataCon":
                        var dt *CEMIDataCon
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.Cemi = dt
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIDataInd":
                        var dt *CEMIDataInd
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.Cemi = dt
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIRawReq":
                        var dt *CEMIRawReq
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.Cemi = dt
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIRawCon":
                        var dt *CEMIRawCon
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.Cemi = dt
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIRawInd":
                        var dt *CEMIRawInd
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.Cemi = dt
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIPollDataReq":
                        var dt *CEMIPollDataReq
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.Cemi = dt
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIPollDataCon":
                        var dt *CEMIPollDataCon
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.Cemi = dt
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIBusmonInd":
                        var dt *CEMIBusmonInd
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.Cemi = dt
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIMPropReadReq":
                        var dt *CEMIMPropReadReq
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.Cemi = dt
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIMPropReadCon":
                        var dt *CEMIMPropReadCon
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.Cemi = dt
                    }
            }
        }
    }
}

func (m TunnelingRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.knxnetip.readwrite.TunnelingRequest"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.TunnelingRequestDataBlock, xml.StartElement{Name: xml.Name{Local: "tunnelingRequestDataBlock"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Cemi, xml.StartElement{Name: xml.Name{Local: "cemi"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}
