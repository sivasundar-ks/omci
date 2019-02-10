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

import "github.com/deckarep/golang-set"

const Vdsl2LineInventoryAndStatusDataPart4ClassId uint16 = 415

var vdsl2lineinventoryandstatusdatapart4BME *ManagedEntityDefinition

// Vdsl2LineInventoryAndStatusDataPart4 (class ID #415) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type Vdsl2LineInventoryAndStatusDataPart4 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	vdsl2lineinventoryandstatusdatapart4BME = &ManagedEntityDefinition{
		Name:    "Vdsl2LineInventoryAndStatusDataPart4",
		ClassID: 415,
		MessageTypes: mapset.NewSetWith(
			Get,
		),
		AllowedAttributeMask: 0x0,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read, false, false, false),
		},
	}
}

// NewVdsl2LineInventoryAndStatusDataPart4 (class ID 415 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewVdsl2LineInventoryAndStatusDataPart4(params ...ParamData) (*ManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: vdsl2lineinventoryandstatusdatapart4BME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
