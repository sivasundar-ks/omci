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
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/google/gopacket"
)

// ManagedEntity provides a complete instance of a Managed Entity
type ManagedEntity struct {
	definition             ManagedEntityDefinition
	attributeMask          uint16
	requestedAttributeMask uint16
	attributes             AttributeValueMap
}

// String provides a simple string that describes this struct
func (entity *ManagedEntity) String() string {
	return fmt.Sprintf("ManagedEntity: %v, EntityID: (%d/%#x): Attributes: %v",
		entity.GetClassID(), entity.GetEntityID(), entity.GetEntityID(), entity.attributes)
}

// NewManagedEntity creates a ManagedEntity given an ME Definition and parameter/attribute data
func NewManagedEntity(definition ManagedEntityDefinition, params ...ParamData) (*ManagedEntity, OmciErrors) {
	entity := &ManagedEntity{
		definition: definition,
		attributes: make(map[string]interface{}),
	}
	if params != nil {
		err := entity.setAttributes(params...)
		if err != nil {
			return nil, err
		}
	}
	return entity, nil
}

// GetManagedEntityDefinition provides the ME definition of a Managed Entity
func (entity *ManagedEntity) GetManagedEntityDefinition() ManagedEntityDefinition {
	return entity.definition
}

// GetName provides the ME Name of a Managed Entity
func (entity ManagedEntity) GetName() string {
	return entity.definition.GetName()
}

// GetClassID returns the 16-bit class ID of a Managed Entity
func (entity ManagedEntity) GetClassID() ClassID {
	return entity.definition.GetClassID()
}

// GetMessageTypes returns the OMCI message types that a Managed Entity supports
func (entity ManagedEntity) GetMessageTypes() mapset.Set {
	return entity.definition.GetMessageTypes()
}

// GetAllowedAttributeMask returns the 16-bit bitmask of attributes a Managed Entity supports
func (entity ManagedEntity) GetAllowedAttributeMask() uint16 {
	return entity.definition.GetAllowedAttributeMask()
}

// GetAttributeDefinitions returns the attribute definition map for a Managed Entity
func (entity ManagedEntity) GetAttributeDefinitions() AttributeDefinitionMap {
	return entity.definition.GetAttributeDefinitions()
}

// DecodeAttributes will decode the attributes portion of a Managed Entity frame/packet
func (entity *ManagedEntity) DecodeAttributes(mask uint16, data []byte, p gopacket.PacketBuilder, msgType byte) (AttributeValueMap, error) {
	return entity.definition.DecodeAttributes(mask, data, p, msgType)
}

// SerializeAttributes will serialize the attributes of a Managed Entity type
func (entity *ManagedEntity) SerializeAttributes(attr AttributeValueMap, mask uint16,
	b gopacket.SerializeBuffer, msgType byte, bytesAvailable int) error {
	return entity.definition.SerializeAttributes(attr, mask, b, msgType, bytesAvailable)
}

// GetEntityID will return the Entity/Instance ID for a Managed Entity
func (entity *ManagedEntity) GetEntityID() uint16 {
	if eid, err := entity.GetAttributeByIndex(0); err == nil {
		return eid.(uint16)
	}
	return 0
}

// SetEntityID will set the Entity/Instance ID for a Managed Entity
func (entity *ManagedEntity) SetEntityID(eid uint16) error {
	return entity.SetAttributeByIndex(0, eid)
}

// GetAttributeMask will return the 16-bit attribute mask of a Managed Entity
func (entity *ManagedEntity) GetAttributeMask() uint16 {
	return entity.attributeMask
}

// SetRequestedAttributeMask is used to initialize the requested attribute mask to a specific
// value. This should only be done on "Get" type operations that need to fetch and attribute
// and store it in the entity. For other operations (create, set, ...) you should specify
// the attributes and values in the Params initialization or use the SetAttribute
func (entity *ManagedEntity) SetRequestedAttributeMask(mask uint16) {
	entity.requestedAttributeMask = mask
}

// GetRequestedAttributeMask will return the 16-bit requested attribute mask of a Managed Entity.
// This is only specified for requests that perform a Get operation
func (entity *ManagedEntity) GetRequestedAttributeMask() uint16 {
	return entity.requestedAttributeMask
}

// GetAttributeValueMap will return the map of attributes of a Managed Entity
func (entity *ManagedEntity) GetAttributeValueMap() AttributeValueMap {
	return entity.attributes
}

// GetAttribute will return the value of a specific attribute for the specified attribute by name
func (entity *ManagedEntity) GetAttribute(name string) (interface{}, error) {
	value, ok := entity.attributes[name]
	if !ok {
		return 0, fmt.Errorf("attribute '%v' not found", name)
	}
	return value, nil
}

// GetAttributeByIndex will return the value of a specific attribute for the specified attribute by index
func (entity *ManagedEntity) GetAttributeByIndex(index uint) (interface{}, error) {
	if len(entity.attributes) == 0 {
		return nil, errors.New("attributes have already been set")
	}
	if _, ok := entity.definition.AttributeDefinitions[index]; !ok {
		return nil, fmt.Errorf("invalid attribute index: %d, should be 0..%d",
			index, len(entity.definition.AttributeDefinitions)-1)
	}
	return entity.GetAttribute(entity.definition.AttributeDefinitions[index].Name)
}

func (entity *ManagedEntity) setAttributes(params ...ParamData) OmciErrors {
	if entity.attributes == nil {
		entity.attributes = make(map[string]interface{})
	} else if len(entity.attributes) > 0 {
		return NewNonStatusError("attributes have already been set")
	}
	eidName := entity.definition.AttributeDefinitions[0].Name
	if len(params) == 0 {
		entity.attributes[eidName] = uint16(0)
		return nil
	}
	entity.attributes[eidName] = params[0].EntityID

	for name, value := range params[0].Attributes {
		if name == eidName {
			continue
		}
		if err := entity.SetAttribute(name, value); err != nil {
			return err
		}
	}
	return nil
}

// SetAttribute can be uses to set the value of a specific attribute by name
func (entity *ManagedEntity) SetAttribute(name string, value interface{}) OmciErrors {
	attrDef, err := GetAttributeDefinitionByName(entity.definition.GetAttributeDefinitions(), name)
	if err != nil {
		// Not found, which means not in the attribute definition map
		return NewProcessingError(err.Error())
	} else if entity.attributes == nil {
		entity.attributes = make(map[string]interface{})
	}
	mask := uint16(1 << (16 - attrDef.GetIndex()))
	// check any constraints
	if constraintCheck := attrDef.GetConstraints(); constraintCheck != nil {
		err = constraintCheck(value)
		if err != nil {
			return NewParameterError(mask, entity.GetAttributeDefinitions(), err)
		}
	}
	entity.attributes[name] = value
	entity.attributeMask |= mask
	return nil
}

// SetAttributeByIndex can be uses to set the value of a specific attribute by attribute index (0..15)
func (entity *ManagedEntity) SetAttributeByIndex(index uint, value interface{}) error {
	attrDef, ok := entity.definition.AttributeDefinitions[index]
	if !ok {
		return fmt.Errorf("invalid attribute index: %d, should be 0..%d",
			index, len(entity.definition.AttributeDefinitions)-1)
	} else if entity.attributes == nil {
		entity.attributes = make(map[string]interface{})
	}
	mask := uint16(1 << (16 - attrDef.GetIndex()))
	// check any constraints
	if constraintCheck := attrDef.GetConstraints(); constraintCheck != nil {
		err := constraintCheck(value)
		if err != nil {
			return NewParameterError(mask, entity.GetAttributeDefinitions(), err)
		}
	}
	entity.attributes[attrDef.Name] = value
	entity.attributeMask |= mask
	return nil
}

// DeleteAttribute is used to remove a specific attribute from a Managed Index by name
func (entity *ManagedEntity) DeleteAttribute(name string) error {
	attrDef, err := GetAttributeDefinitionByName(entity.definition.GetAttributeDefinitions(), name)
	if err != nil {
		return err
	}
	if entity.attributes != nil {
		delete(entity.attributes, name)
		entity.attributeMask &= ^uint16(1 << (16 - attrDef.GetIndex()))
	}
	return nil
}

// DeleteAttributeByIndex is used to remove a specific attribute from a Managed Index by attribute index (0..15)
func (entity *ManagedEntity) DeleteAttributeByIndex(index uint) error {
	attrDef, ok := entity.definition.AttributeDefinitions[index]
	if !ok {
		return fmt.Errorf("invalid attribute index: %d, should be 0..%d",
			index, len(entity.definition.AttributeDefinitions)-1)
	}
	if entity.attributes != nil {
		delete(entity.attributes, attrDef.Name)
		entity.attributeMask &= ^uint16(1 << (16 - attrDef.GetIndex()))
	}
	return nil
}

// DecodeFromBytes decodes a Managed Entity give an octet stream pointing to the ME within a frame
func (entity *ManagedEntity) DecodeFromBytes(data []byte, p gopacket.PacketBuilder, msgType byte) error {
	if len(data) < 6 {
		p.SetTruncated()
		return errors.New("frame too small")
	}
	classID := ClassID(binary.BigEndian.Uint16(data[0:2]))
	entityID := binary.BigEndian.Uint16(data[2:4])
	parameters := ParamData{EntityID: entityID}

	meDefinition, omciErr := LoadManagedEntityDefinition(classID, parameters)
	if omciErr != nil {
		return omciErr.GetError()
	}
	entity.definition = meDefinition.definition
	entity.attributeMask = binary.BigEndian.Uint16(data[4:6])
	entity.attributes = make(map[string]interface{})
	entity.SetEntityID(entityID)
	packetAttributes, err := entity.DecodeAttributes(entity.GetAttributeMask(), data[6:], p, msgType)
	if err != nil {
		return err
	}
	for name, value := range packetAttributes {
		entity.attributes[name] = value
	}
	return nil
}

// SerializeTo serializes a Managed Entity into an octet stream
func (entity *ManagedEntity) SerializeTo(b gopacket.SerializeBuffer, msgType byte, bytesAvailable int) error {
	// Add class ID and entity ID
	bytes, err := b.AppendBytes(6)
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint16(bytes, uint16(entity.GetClassID()))
	binary.BigEndian.PutUint16(bytes[2:], entity.GetEntityID())
	binary.BigEndian.PutUint16(bytes[4:], entity.GetAttributeMask())

	// TODO: Need to limit number of bytes appended to not exceed packet size
	// Is there space/metadata info in 'b' parameter to allow this?
	err = entity.SerializeAttributes(entity.attributes, entity.GetAttributeMask(), b, msgType, bytesAvailable)
	return err
}
