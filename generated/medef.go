/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */
package generated

import (
	"errors"
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/google/gopacket"
	"math/bits"
)

type ManagedEntityDefinition struct {
	Name                 string
	ClassID              ClassID
	MessageTypes         mapset.Set
	AllowedAttributeMask uint16
	AttributeDefinitions AttributeDefinitionMap
}

func (bme *ManagedEntityDefinition) String() string {
	return fmt.Sprintf("Definition: %s: CID: %v, Attributes: %v",
		bme.Name, bme.ClassID, bme.AttributeDefinitions)
}

func (bme *ManagedEntityDefinition) GetName() string {
	return bme.Name
}
func (bme *ManagedEntityDefinition) GetClassID() ClassID {
	return bme.ClassID
}
func (bme *ManagedEntityDefinition) GetMessageTypes() mapset.Set {
	return bme.MessageTypes
}
func (bme *ManagedEntityDefinition) GetAllowedAttributeMask() uint16 {
	return bme.AllowedAttributeMask
}
func (bme *ManagedEntityDefinition) GetAttributeDefinitions() *AttributeDefinitionMap {
	return &bme.AttributeDefinitions
}

func (bme *ManagedEntityDefinition) DecodeAttributes(mask uint16, data []byte, p gopacket.PacketBuilder, msgType byte) (AttributeValueMap, error) {
	if (mask | bme.GetAllowedAttributeMask()) != bme.GetAllowedAttributeMask() {
		// TODO: Provide custom error code so a response 'result' can properly be coded
		return nil, errors.New("unsupported attribute mask")
	}
	keyList := GetAttributeDefinitionMapKeys(bme.AttributeDefinitions)

	attrMap := make(AttributeValueMap, bits.OnesCount16(mask))
	for _, index := range keyList {
		if index == 0 {
			continue // Skip Entity ID
		}
		attrDef := bme.AttributeDefinitions[index]
		name := attrDef.GetName()

		if mask&(1<<(16-uint(index))) != 0 {
			value, err := attrDef.Decode(data, p, msgType)
			if err != nil {
				return nil, err
			}
			if attrDef.IsTableAttribute() {
				switch msgType {
				default:
					return nil, errors.New(fmt.Sprintf("unsupported Message Type '%v' for table serialization", msgType))

				case byte(Get) | AK: // Get Response
					attrMap[name] = value
					data = data[4:]

				case byte(GetNext) | AK: // Get Next Response
					// Value is a partial octet buffer we need to collect and at
					// the end (last segment) pull it up into more appropriate table
					// rows
					valueBuffer, ok := value.([]byte)
					if !ok {
						panic("unexpected type already returned as get-next-response attribute data")
					}
					if existing, found := attrMap[name]; found {
						prev, ok := existing.([]byte)
						if !ok {
							panic("unexpected type already in attribute value map")
						}
						attrMap[name] = append(prev, valueBuffer...)
					} else {
						attrMap[name] = valueBuffer
					}
					if size := attrDef.GetSize(); size != 0 && size != len(valueBuffer) {
						panic("unexpected size difference")
					}
					data = data[len(valueBuffer):]

				case byte(Set) | AR: // Set Request
					fmt.Println("TODO")

				case byte(SetTable) | AR: // Set Table Request
					// TODO: Only baseline supported at this time
					return nil, errors.New("attribute encode for set-table-request not yet supported")
				}
			} else {
				attrMap[name] = value
				data = data[attrDef.GetSize():]
			}
		}
	}
	return attrMap, nil
}

func (bme *ManagedEntityDefinition) SerializeAttributes(attr AttributeValueMap, mask uint16,
	b gopacket.SerializeBuffer, msgType byte, bytesAvailable int) error {
	if (mask | bme.GetAllowedAttributeMask()) != bme.GetAllowedAttributeMask() {
		// TODO: Provide custom error code so a response 'result' can properly be coded
		return errors.New("unsupported attribute mask")
	}
	// TODO: Need to limit number of bytes appended to not exceed packet size
	// Is there space/metadata info in 'b' parameter to allow this?
	keyList := GetAttributeDefinitionMapKeys(bme.AttributeDefinitions)

	for _, index := range keyList {
		if index == 0 {
			continue // Skip Entity ID
		}
		attrDef := bme.AttributeDefinitions[index]

		if mask&(1<<(16-uint(index))) != 0 {
			value, ok := attr[attrDef.GetName()]
			if !ok {
				msg := fmt.Sprintf("attribute not found: '%v'", attrDef.GetName())
				return errors.New(msg)
			}
			size, err := attrDef.SerializeTo(value, b, msgType, bytesAvailable)
			if err != nil {
				return err
			}
			bytesAvailable -= size
		}
	}
	return nil
}
