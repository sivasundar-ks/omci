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

// NetworkDialPlanTableClassID is the 16-bit ID for the OMCI
// Managed entity Network dial plan table
const NetworkDialPlanTableClassID ClassID = ClassID(145)

var networkdialplantableBME *ManagedEntityDefinition

// NetworkDialPlanTable (class ID #145)
//	The network dial plan table ME is optional for ONUs providing VoIP services. This ME is used to
//	provision dial plans from the OLT. Instances of this ME are created and deleted by the OLT. If a
//	non-OMCI interface is used to manage SIP for VoIP, this ME is unnecessary.
//
//	Relationships
//		An instance of this ME may be associated with one or more instances of the SIP user data ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Dial Plan Number
//			Dial plan number: This attribute indicates the current number of dial plans in the dial plan
//			table. (R) (mandatory) (2-bytes)
//
//		Dial Plan Table Max Size
//			Dial plan table max size: This attribute defines the maximum number of dial plans that can be
//			stored in the dial plan table. (R, setbycreate) (mandatory) (2-bytes)
//
//		Critical Dial Timeout
//			Critical dial timeout: This attribute defines the critical dial timeout for digit map
//			processing, in milliseconds. The recommended default value is 4000-ms. (R,-W, setbycreate)
//			(mandatory) (2-bytes)
//
//		Partial Dial Timeout
//			Partial dial timeout: This attribute defines the partial dial timeout for digit map processing,
//			in milliseconds. The recommended default value is 16000-ms. (R,-W, setbycreate) (mandatory)
//			(2-bytes)
//
//		Dial Plan Format
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Dial Plan Table
//			(R,-W) (mandatory) (30 * N bytes, where N is the number of dial plans)
//
type NetworkDialPlanTable struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	networkdialplantableBME = &ManagedEntityDefinition{
		Name:    "NetworkDialPlanTable",
		ClassID: 145,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetNext,
			Set,
		),
		AllowedAttributeMask: 0xfc00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 0),
			1: Uint16Field("DialPlanNumber", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, 1),
			2: Uint16Field("DialPlanTableMaxSize", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 2),
			3: Uint16Field("CriticalDialTimeout", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 3),
			4: Uint16Field("PartialDialTimeout", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 4),
			5: ByteField("DialPlanFormat", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 5),
			6: TableField("DialPlanTable", TableAttributeType, TableInfo{nil, 30}, mapset.NewSetWith(Read, Write), false, false, 6),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewNetworkDialPlanTable (class ID 145) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewNetworkDialPlanTable(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*networkdialplantableBME, params...)
}
