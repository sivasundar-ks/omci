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

// XdslSubcarrierMaskingDownstreamProfileClassID is the 16-bit ID for the OMCI
// Managed entity xDSL subcarrier masking downstream profile
const XdslSubcarrierMaskingDownstreamProfileClassID ClassID = ClassID(108)

var xdslsubcarriermaskingdownstreamprofileBME *ManagedEntityDefinition

// XdslSubcarrierMaskingDownstreamProfile (class ID #108)
//	This ME contains the subcarrier masking downstream profile for an xDSL UNI. Instances of this ME
//	are created and deleted by the OLT.
//
//	Relationships
//		An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI part
//		1.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0 is
//			reserved. (R, set-by-create) (mandatory) (2-bytes)
//
//		Downstream Subcarrier Mask 1
//			Downstream subcarrier mask 1: Subcarriers 1 to 128. (R,-W, set-by-create) (mandatory) (16-bytes)
//
//		Downstream Subcarrier Mask 2
//			Downstream subcarrier mask 2: Subcarriers 129 to 256. (R,-W) (mandatory for modems that support
//			NSCds-> 128) (16-bytes)
//
//		Downstream Subcarrier Mask 3
//			Downstream subcarrier mask 3: Subcarriers 257 to 384. (R,-W) (mandatory for modems that support
//			NSCds-> 256) (16-bytes)
//
//		Downstream Subcarrier Mask 4
//			Downstream subcarrier mask 4: Subcarriers 385 to 512. (R,-W) (mandatory for modems that support
//			NSCds-> 384) (16-bytes)
//
//		Mask Valid
//			(R,-W) (mandatory) (1-byte)
//
type XdslSubcarrierMaskingDownstreamProfile struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdslsubcarriermaskingdownstreamprofileBME = &ManagedEntityDefinition{
		Name:    "XdslSubcarrierMaskingDownstreamProfile",
		ClassID: 108,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xf800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 0),
			1: MultiByteField("DownstreamSubcarrierMask1", OctetsAttributeType, 16, nil, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 1),
			2: MultiByteField("DownstreamSubcarrierMask2", OctetsAttributeType, 16, nil, mapset.NewSetWith(Read, Write), false, false, 2),
			3: MultiByteField("DownstreamSubcarrierMask3", OctetsAttributeType, 16, nil, mapset.NewSetWith(Read, Write), false, false, 3),
			4: MultiByteField("DownstreamSubcarrierMask4", OctetsAttributeType, 16, nil, mapset.NewSetWith(Read, Write), false, false, 4),
			5: ByteField("MaskValid", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, 5),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewXdslSubcarrierMaskingDownstreamProfile (class ID 108) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewXdslSubcarrierMaskingDownstreamProfile(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdslsubcarriermaskingdownstreamprofileBME, params...)
}
