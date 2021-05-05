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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

type SupportedPhysicalMedia uint8

type ISupportedPhysicalMedia interface {
	KnxSupport() bool
	Description() string
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	SupportedPhysicalMedia_OTHER                     SupportedPhysicalMedia = 0x00
	SupportedPhysicalMedia_OIL_METER                 SupportedPhysicalMedia = 0x01
	SupportedPhysicalMedia_ELECTRICITY_METER         SupportedPhysicalMedia = 0x02
	SupportedPhysicalMedia_GAS_METER                 SupportedPhysicalMedia = 0x03
	SupportedPhysicalMedia_HEAT_METER                SupportedPhysicalMedia = 0x04
	SupportedPhysicalMedia_STEAM_METER               SupportedPhysicalMedia = 0x05
	SupportedPhysicalMedia_WARM_WATER_METER          SupportedPhysicalMedia = 0x06
	SupportedPhysicalMedia_WATER_METER               SupportedPhysicalMedia = 0x07
	SupportedPhysicalMedia_HEAT_COST_ALLOCATOR       SupportedPhysicalMedia = 0x08
	SupportedPhysicalMedia_COMPRESSED_AIR            SupportedPhysicalMedia = 0x09
	SupportedPhysicalMedia_COOLING_LOAD_METER_INLET  SupportedPhysicalMedia = 0x0A
	SupportedPhysicalMedia_COOLING_LOAD_METER_OUTLET SupportedPhysicalMedia = 0x0B
	SupportedPhysicalMedia_HEAT_INLET                SupportedPhysicalMedia = 0x0C
	SupportedPhysicalMedia_HEAT_AND_COOL             SupportedPhysicalMedia = 0x0D
	SupportedPhysicalMedia_BUS_OR_SYSTEM             SupportedPhysicalMedia = 0x0E
	SupportedPhysicalMedia_UNKNOWN_DEVICE_TYPE       SupportedPhysicalMedia = 0x0F
	SupportedPhysicalMedia_BREAKER                   SupportedPhysicalMedia = 0x20
	SupportedPhysicalMedia_VALVE                     SupportedPhysicalMedia = 0x21
	SupportedPhysicalMedia_WASTE_WATER_METER         SupportedPhysicalMedia = 0x28
	SupportedPhysicalMedia_GARBAGE                   SupportedPhysicalMedia = 0x29
	SupportedPhysicalMedia_RADIO_CONVERTER           SupportedPhysicalMedia = 0x37
)

var SupportedPhysicalMediaValues []SupportedPhysicalMedia

func init() {
	SupportedPhysicalMediaValues = []SupportedPhysicalMedia{
		SupportedPhysicalMedia_OTHER,
		SupportedPhysicalMedia_OIL_METER,
		SupportedPhysicalMedia_ELECTRICITY_METER,
		SupportedPhysicalMedia_GAS_METER,
		SupportedPhysicalMedia_HEAT_METER,
		SupportedPhysicalMedia_STEAM_METER,
		SupportedPhysicalMedia_WARM_WATER_METER,
		SupportedPhysicalMedia_WATER_METER,
		SupportedPhysicalMedia_HEAT_COST_ALLOCATOR,
		SupportedPhysicalMedia_COMPRESSED_AIR,
		SupportedPhysicalMedia_COOLING_LOAD_METER_INLET,
		SupportedPhysicalMedia_COOLING_LOAD_METER_OUTLET,
		SupportedPhysicalMedia_HEAT_INLET,
		SupportedPhysicalMedia_HEAT_AND_COOL,
		SupportedPhysicalMedia_BUS_OR_SYSTEM,
		SupportedPhysicalMedia_UNKNOWN_DEVICE_TYPE,
		SupportedPhysicalMedia_BREAKER,
		SupportedPhysicalMedia_VALVE,
		SupportedPhysicalMedia_WASTE_WATER_METER,
		SupportedPhysicalMedia_GARBAGE,
		SupportedPhysicalMedia_RADIO_CONVERTER,
	}
}

func (e SupportedPhysicalMedia) KnxSupport() bool {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return true
		}
	case 0x01:
		{ /* '0x01' */
			return true
		}
	case 0x02:
		{ /* '0x02' */
			return true
		}
	case 0x03:
		{ /* '0x03' */
			return true
		}
	case 0x04:
		{ /* '0x04' */
			return true
		}
	case 0x05:
		{ /* '0x05' */
			return true
		}
	case 0x06:
		{ /* '0x06' */
			return true
		}
	case 0x07:
		{ /* '0x07' */
			return true
		}
	case 0x08:
		{ /* '0x08' */
			return true
		}
	case 0x09:
		{ /* '0x09' */
			return false
		}
	case 0x0A:
		{ /* '0x0A' */
			return true
		}
	case 0x0B:
		{ /* '0x0B' */
			return true
		}
	case 0x0C:
		{ /* '0x0C' */
			return true
		}
	case 0x0D:
		{ /* '0x0D' */
			return true
		}
	case 0x0E:
		{ /* '0x0E' */
			return false
		}
	case 0x0F:
		{ /* '0x0F' */
			return false
		}
	case 0x20:
		{ /* '0x20' */
			return true
		}
	case 0x21:
		{ /* '0x21' */
			return true
		}
	case 0x28:
		{ /* '0x28' */
			return true
		}
	case 0x29:
		{ /* '0x29' */
			return true
		}
	case 0x37:
		{ /* '0x37' */
			return false
		}
	default:
		{
			return false
		}
	}
}

func (e SupportedPhysicalMedia) Description() string {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return "used_for_undefined_physical_medium"
		}
	case 0x01:
		{ /* '0x01' */
			return "measures_volume_of_oil"
		}
	case 0x02:
		{ /* '0x02' */
			return "measures_electric_energy"
		}
	case 0x03:
		{ /* '0x03' */
			return "measures_volume_of_gaseous_energy"
		}
	case 0x04:
		{ /* '0x04' */
			return "heat_energy_measured_in_outlet_pipe"
		}
	case 0x05:
		{ /* '0x05' */
			return "measures_weight_of_hot_steam"
		}
	case 0x06:
		{ /* '0x06' */
			return "measured_heated_water_volume"
		}
	case 0x07:
		{ /* '0x07' */
			return "measured_water_volume"
		}
	case 0x08:
		{ /* '0x08' */
			return "measured_relative_cumulated_heat_consumption"
		}
	case 0x09:
		{ /* '0x09' */
			return "measures_weight_of_compressed_air"
		}
	case 0x0A:
		{ /* '0x0A' */
			return "cooling_energy_measured_in_inlet_pipe"
		}
	case 0x0B:
		{ /* '0x0B' */
			return "cooling_energy_measured_in_outlet_pipe"
		}
	case 0x0C:
		{ /* '0x0C' */
			return "heat_energy_measured_in_inlet_pipe"
		}
	case 0x0D:
		{ /* '0x0D' */
			return "measures_both_heat_and_cool"
		}
	case 0x0E:
		{ /* '0x0E' */
			return "no_meter"
		}
	case 0x0F:
		{ /* '0x0F' */
			return "used_for_undefined_physical_medium"
		}
	case 0x20:
		{ /* '0x20' */
			return "status_of_electric_energy_supply"
		}
	case 0x21:
		{ /* '0x21' */
			return "status_of_supply_of_Gas_or_water"
		}
	case 0x28:
		{ /* '0x28' */
			return "measured_volume_of_disposed_water"
		}
	case 0x29:
		{ /* '0x29' */
			return "measured_weight_of_disposed_rubbish"
		}
	case 0x37:
		{ /* '0x37' */
			return "enables_the_radio_transmission_of_a_meter_without_a_radio_interface"
		}
	default:
		{
			return ""
		}
	}
}
func SupportedPhysicalMediaByValue(value uint8) SupportedPhysicalMedia {
	switch value {
	case 0x00:
		return SupportedPhysicalMedia_OTHER
	case 0x01:
		return SupportedPhysicalMedia_OIL_METER
	case 0x02:
		return SupportedPhysicalMedia_ELECTRICITY_METER
	case 0x03:
		return SupportedPhysicalMedia_GAS_METER
	case 0x04:
		return SupportedPhysicalMedia_HEAT_METER
	case 0x05:
		return SupportedPhysicalMedia_STEAM_METER
	case 0x06:
		return SupportedPhysicalMedia_WARM_WATER_METER
	case 0x07:
		return SupportedPhysicalMedia_WATER_METER
	case 0x08:
		return SupportedPhysicalMedia_HEAT_COST_ALLOCATOR
	case 0x09:
		return SupportedPhysicalMedia_COMPRESSED_AIR
	case 0x0A:
		return SupportedPhysicalMedia_COOLING_LOAD_METER_INLET
	case 0x0B:
		return SupportedPhysicalMedia_COOLING_LOAD_METER_OUTLET
	case 0x0C:
		return SupportedPhysicalMedia_HEAT_INLET
	case 0x0D:
		return SupportedPhysicalMedia_HEAT_AND_COOL
	case 0x0E:
		return SupportedPhysicalMedia_BUS_OR_SYSTEM
	case 0x0F:
		return SupportedPhysicalMedia_UNKNOWN_DEVICE_TYPE
	case 0x20:
		return SupportedPhysicalMedia_BREAKER
	case 0x21:
		return SupportedPhysicalMedia_VALVE
	case 0x28:
		return SupportedPhysicalMedia_WASTE_WATER_METER
	case 0x29:
		return SupportedPhysicalMedia_GARBAGE
	case 0x37:
		return SupportedPhysicalMedia_RADIO_CONVERTER
	}
	return 0
}

func SupportedPhysicalMediaByName(value string) SupportedPhysicalMedia {
	switch value {
	case "OTHER":
		return SupportedPhysicalMedia_OTHER
	case "OIL_METER":
		return SupportedPhysicalMedia_OIL_METER
	case "ELECTRICITY_METER":
		return SupportedPhysicalMedia_ELECTRICITY_METER
	case "GAS_METER":
		return SupportedPhysicalMedia_GAS_METER
	case "HEAT_METER":
		return SupportedPhysicalMedia_HEAT_METER
	case "STEAM_METER":
		return SupportedPhysicalMedia_STEAM_METER
	case "WARM_WATER_METER":
		return SupportedPhysicalMedia_WARM_WATER_METER
	case "WATER_METER":
		return SupportedPhysicalMedia_WATER_METER
	case "HEAT_COST_ALLOCATOR":
		return SupportedPhysicalMedia_HEAT_COST_ALLOCATOR
	case "COMPRESSED_AIR":
		return SupportedPhysicalMedia_COMPRESSED_AIR
	case "COOLING_LOAD_METER_INLET":
		return SupportedPhysicalMedia_COOLING_LOAD_METER_INLET
	case "COOLING_LOAD_METER_OUTLET":
		return SupportedPhysicalMedia_COOLING_LOAD_METER_OUTLET
	case "HEAT_INLET":
		return SupportedPhysicalMedia_HEAT_INLET
	case "HEAT_AND_COOL":
		return SupportedPhysicalMedia_HEAT_AND_COOL
	case "BUS_OR_SYSTEM":
		return SupportedPhysicalMedia_BUS_OR_SYSTEM
	case "UNKNOWN_DEVICE_TYPE":
		return SupportedPhysicalMedia_UNKNOWN_DEVICE_TYPE
	case "BREAKER":
		return SupportedPhysicalMedia_BREAKER
	case "VALVE":
		return SupportedPhysicalMedia_VALVE
	case "WASTE_WATER_METER":
		return SupportedPhysicalMedia_WASTE_WATER_METER
	case "GARBAGE":
		return SupportedPhysicalMedia_GARBAGE
	case "RADIO_CONVERTER":
		return SupportedPhysicalMedia_RADIO_CONVERTER
	}
	return 0
}

func CastSupportedPhysicalMedia(structType interface{}) SupportedPhysicalMedia {
	castFunc := func(typ interface{}) SupportedPhysicalMedia {
		if sSupportedPhysicalMedia, ok := typ.(SupportedPhysicalMedia); ok {
			return sSupportedPhysicalMedia
		}
		return 0
	}
	return castFunc(structType)
}

func (m SupportedPhysicalMedia) LengthInBits() uint16 {
	return 8
}

func (m SupportedPhysicalMedia) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SupportedPhysicalMediaParse(readBuffer utils.ReadBuffer) (SupportedPhysicalMedia, error) {
	val, err := readBuffer.ReadUint8("SupportedPhysicalMedia", 8)
	if err != nil {
		return 0, nil
	}
	return SupportedPhysicalMediaByValue(val), nil
}

func (e SupportedPhysicalMedia) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("SupportedPhysicalMedia", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (m *SupportedPhysicalMedia) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.CharData:
			tok := token.(xml.CharData)
			*m = SupportedPhysicalMediaByName(string(tok))
		}
	}
}

func (m SupportedPhysicalMedia) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.String(), start); err != nil {
		return err
	}
	return nil
}

func (e SupportedPhysicalMedia) name() string {
	switch e {
	case SupportedPhysicalMedia_OTHER:
		return "OTHER"
	case SupportedPhysicalMedia_OIL_METER:
		return "OIL_METER"
	case SupportedPhysicalMedia_ELECTRICITY_METER:
		return "ELECTRICITY_METER"
	case SupportedPhysicalMedia_GAS_METER:
		return "GAS_METER"
	case SupportedPhysicalMedia_HEAT_METER:
		return "HEAT_METER"
	case SupportedPhysicalMedia_STEAM_METER:
		return "STEAM_METER"
	case SupportedPhysicalMedia_WARM_WATER_METER:
		return "WARM_WATER_METER"
	case SupportedPhysicalMedia_WATER_METER:
		return "WATER_METER"
	case SupportedPhysicalMedia_HEAT_COST_ALLOCATOR:
		return "HEAT_COST_ALLOCATOR"
	case SupportedPhysicalMedia_COMPRESSED_AIR:
		return "COMPRESSED_AIR"
	case SupportedPhysicalMedia_COOLING_LOAD_METER_INLET:
		return "COOLING_LOAD_METER_INLET"
	case SupportedPhysicalMedia_COOLING_LOAD_METER_OUTLET:
		return "COOLING_LOAD_METER_OUTLET"
	case SupportedPhysicalMedia_HEAT_INLET:
		return "HEAT_INLET"
	case SupportedPhysicalMedia_HEAT_AND_COOL:
		return "HEAT_AND_COOL"
	case SupportedPhysicalMedia_BUS_OR_SYSTEM:
		return "BUS_OR_SYSTEM"
	case SupportedPhysicalMedia_UNKNOWN_DEVICE_TYPE:
		return "UNKNOWN_DEVICE_TYPE"
	case SupportedPhysicalMedia_BREAKER:
		return "BREAKER"
	case SupportedPhysicalMedia_VALVE:
		return "VALVE"
	case SupportedPhysicalMedia_WASTE_WATER_METER:
		return "WASTE_WATER_METER"
	case SupportedPhysicalMedia_GARBAGE:
		return "GARBAGE"
	case SupportedPhysicalMedia_RADIO_CONVERTER:
		return "RADIO_CONVERTER"
	}
	return ""
}

func (e SupportedPhysicalMedia) String() string {
	return e.name()
}
