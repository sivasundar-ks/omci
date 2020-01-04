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

// AuthenticationSecurityMethodClassID is the 16-bit ID for the OMCI
// Managed entity Authentication security method
const AuthenticationSecurityMethodClassID ClassID = ClassID(148)

var authenticationsecuritymethodBME *ManagedEntityDefinition

// AuthenticationSecurityMethod (class ID #148)
//	The authentication security method defines the user ID and password configuration to establish a
//	session between a client and a server. This object may be used in the role of the client or
//	server. An instance of this ME is created by the OLT if authenticated communication is
//	necessary.
//
//	Relationships
//		One instance of this management entity may be associated with a network address ME. This ME may
//		also be cited by other MEs that require authentication parameter management.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0xFFFF
//			is reserved. (R, setbycreate) (mandatory) (2-bytes)
//
//		Validation Scheme
//			(R,-W) (mandatory) (1-byte)
//
//		Username 1
//			Username 1:	This string attribute is the user name. If the string is shorter than 25-bytes, it
//			must be null terminated (Note). (R,-W) (mandatory) (25-bytes)
//
//		Password
//			Password:	This string attribute is the password. If the string is shorter than 25-bytes, it must
//			be null terminated. (R,-W) (mandatory) (25-bytes)
//
//		Realm
//			Realm:	This string attribute specifies the realm used in digest authentication. If the string is
//			shorter than 25-bytes, it must be null terminated. (R,-W) (mandatory) (25-bytes)
//
//		Username 2
//			NOTE - The total username is the concatenation of the username 1 and username 2 attributes if
//			and only if: a) username 1 comprises 25 non-null characters; b) username 2 is supported by the
//			ONU; and c) username 2 contains a leading non-null character string. Otherwise, the total
//			username is simply the value of the username 1 attribute.
//
type AuthenticationSecurityMethod struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	authenticationsecuritymethodBME = &ManagedEntityDefinition{
		Name:    "AuthenticationSecurityMethod",
		ClassID: 148,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xf800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 0),
			1: ByteField("ValidationScheme", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, 1),
			2: MultiByteField("Username1", OctetsAttributeType, 25, nil, mapset.NewSetWith(Read, Write), false, false, 2),
			3: MultiByteField("Password", OctetsAttributeType, 25, nil, mapset.NewSetWith(Read, Write), false, false, 3),
			4: MultiByteField("Realm", OctetsAttributeType, 25, nil, mapset.NewSetWith(Read, Write), false, false, 4),
			5: MultiByteField("Username2", OctetsAttributeType, 25, nil, mapset.NewSetWith(Read, Write), false, true, 5),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewAuthenticationSecurityMethod (class ID 148) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewAuthenticationSecurityMethod(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*authenticationsecuritymethodBME, params...)
}
