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

// ThresholdData2ClassID is the 16-bit ID for the OMCI
// Managed entity Threshold data 2
const ThresholdData2ClassID ClassID = ClassID(274)

var thresholddata2BME *ManagedEntityDefinition

// ThresholdData2 (class ID #274)
//	Together with an instance of the threshold data 1 ME, an instance of this ME contains threshold
//	values for counters maintained in one or more instances of PM history data MEs.
//
//	For a complete discussion of generic PM architecture, refer to clause-I.4.
//
//	Instances of this ME are created and deleted by the OLT.
//
//	Relationships
//		Refer to the relationships of the threshold data 1 ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Its value is the
//			same as that of the paired threshold data-1-instance. (R, setbycreate) (mandatory) (2-bytes)
//
//		Threshold Value_8
//			Threshold value-8: (R,-W, setbycreate) (mandatory) (4-bytes)
//
//		Threshold Value_9
//			Threshold value-9: (R,-W, setbycreate) (mandatory) (4-bytes)
//
//		Threshold Value_10
//			Threshold value-10: (R,-W, setbycreate) (mandatory) (4-bytes)
//
//		Threshold Value_11
//			Threshold value-11: (R,-W, setbycreate) (mandatory) (4-bytes)
//
//		Threshold Value_12
//			Threshold value-12: (R,-W, setbycreate) (mandatory) (4-bytes)
//
//		Threshold Value_13
//			Threshold value-13: (R,-W, setbycreate) (mandatory) (4-bytes)
//
//		Threshold Value_14
//			Threshold value-14: (R,-W, setbycreate) (mandatory) (4-bytes)
//
type ThresholdData2 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	thresholddata2BME = &ManagedEntityDefinition{
		Name:    "ThresholdData2",
		ClassID: 274,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xfe00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 0),
			1: Uint32Field("ThresholdValue8", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 1),
			2: Uint32Field("ThresholdValue9", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 2),
			3: Uint32Field("ThresholdValue10", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 3),
			4: Uint32Field("ThresholdValue11", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 4),
			5: Uint32Field("ThresholdValue12", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 5),
			6: Uint32Field("ThresholdValue13", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 6),
			7: Uint32Field("ThresholdValue14", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 7),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewThresholdData2 (class ID 274) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewThresholdData2(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*thresholddata2BME, params...)
}
