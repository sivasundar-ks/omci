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

// Dot1XPortExtensionPackageClassID is the 16-bit ID for the OMCI
// Managed entity Dot1X port extension package
const Dot1XPortExtensionPackageClassID ClassID = ClassID(290)

var dot1xportextensionpackageBME *ManagedEntityDefinition

// Dot1XPortExtensionPackage (class ID #290)
//	An instance of this ME represents a set of attributes that control a port's IEEE 802.1X
//	operation. It is created and deleted autonomously by the ONU upon the creation or deletion of a
//	PPTP that supports [IEEE 802.1X] authentication of customer premises equipment (CPE).
//
//	Relationships
//		An instance of this ME is associated with a PPTP that performs IEEE 802.1X authentication of CPE
//		(e.g., Ethernet or DSL).
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute provides a unique number for each instance of this ME. Its
//			value is the same as that of its associated PPTP (i.e., slot and port number). (R) (mandatory)
//			(2-bytes)
//
//		Dot1X Enable
//			Dot1x enable: If true, this Boolean attribute forces the associated port to authenticate via
//			[IEEE 802.1X] as a precondition of normal service. The default value false does not impose IEEE
//			802.1X authentication on the associated port. (R,-W) (mandatory) (1-byte)
//
//		Action Register
//			(W) (mandatory) (1-byte)
//
//		Authenticator Pae State
//			(R) (optional) (1-byte)
//
//		Backend Authentication State
//			(R) (optional) (1-byte)
//
//		Admin Controlled Directions
//			Admin controlled directions: This attribute controls the directionality of the port's
//			authentication requirement. The default value 0 indicates that control is imposed in both
//			directions. The value 1 indicates that control is imposed only on traffic from the subscriber
//			towards the network. (R,-W) (optional) (1-byte)
//
//		Operational Controlled Directions
//			Operational controlled directions: This attribute indicates the directionality of the port's
//			current authentication state. The value 0 indicates that control is imposed in both directions.
//			The value 1 indicates that control is imposed only on traffic from the subscriber towards the
//			network. (R) (optional) (1-byte)
//
//		Authenticator Controlled Port Status
//			Authenticator controlled port status: This attribute indicates whether the controlled port is
//			currently authorized (1) or unauthorized (2). (R) (optional) (1-byte)
//
//		Quiet Period
//			Quiet period: This attribute specifies the interval between EAP request/identity invitations
//			sent to the peer. Other events such as carrier present or EAPOL start frames from the peer may
//			trigger an EAP request/identity frame from the ONU at any time; this attribute controls the
//			ONU's periodic behaviour in the absence of these other inputs. It is expressed in seconds.
//			(R,-W) (optional) (2-bytes)
//
//		Server Timeout Period
//			Server timeout period: This attribute specifies the time the ONU will wait for a response from
//			the radius server before timing out. Within this maximum interval, the ONU may initiate several
//			retransmissions with exponentially increasing delay. Upon timeout, the ONU may try another
//			radius server if there is one, or invoke the fallback policy, if no alternate radius servers are
//			available. Server timeout is expressed in seconds, with a default value of 30 and a maximum
//			value of 65535. (R,-W) (optional) (2-bytes)
//
//		Re_Authentication Period
//			Re-authentication period: This attribute records the re-authentication interval specified by the
//			radius authentication server. It is expressed in seconds. The attribute is only meaningful after
//			a port has been authenticated. (R) (optional) (2-bytes)
//
//		Re_Authentication Enabled
//			Re-authentication enabled: This Boolean attribute records whether the radius authentication
//			server has enabled re-authentication on this service (true) or not (false). The attribute is
//			only meaningful after a port has been authenticated. (R) (optional) (1-byte)
//
//		Key Transmission Enabled
//			Key transmission enabled: This Boolean attribute indicates whether key transmission is enabled
//			(true) or not (false). This feature is not required; the parameter is listed here for
//			completeness vis-`a-vis [IEEE 802.1X]. (R,-W) (optional) (1-byte)
//
type Dot1XPortExtensionPackage struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	dot1xportextensionpackageBME = &ManagedEntityDefinition{
		Name:    "Dot1XPortExtensionPackage",
		ClassID: 290,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0xfff0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1:  ByteField("Dot1XEnable", UnsignedIntegerAttributeType, 0x8000, 0, mapset.NewSetWith(Read, Write), false, false, false, 1),
			2:  ByteField("ActionRegister", UnsignedIntegerAttributeType, 0x4000, 0, mapset.NewSetWith(Write), false, false, false, 2),
			3:  ByteField("AuthenticatorPaeState", UnsignedIntegerAttributeType, 0x2000, 0, mapset.NewSetWith(Read), false, true, false, 3),
			4:  ByteField("BackendAuthenticationState", UnsignedIntegerAttributeType, 0x1000, 0, mapset.NewSetWith(Read), false, true, false, 4),
			5:  ByteField("AdminControlledDirections", UnsignedIntegerAttributeType, 0x0800, 0, mapset.NewSetWith(Read, Write), false, true, false, 5),
			6:  ByteField("OperationalControlledDirections", UnsignedIntegerAttributeType, 0x0400, 0, mapset.NewSetWith(Read), false, true, false, 6),
			7:  ByteField("AuthenticatorControlledPortStatus", UnsignedIntegerAttributeType, 0x0200, 0, mapset.NewSetWith(Read), false, true, false, 7),
			8:  Uint16Field("QuietPeriod", UnsignedIntegerAttributeType, 0x0100, 0, mapset.NewSetWith(Read, Write), false, true, false, 8),
			9:  Uint16Field("ServerTimeoutPeriod", UnsignedIntegerAttributeType, 0x0080, 0, mapset.NewSetWith(Read, Write), false, true, false, 9),
			10: Uint16Field("ReAuthenticationPeriod", UnsignedIntegerAttributeType, 0x0040, 0, mapset.NewSetWith(Read), false, true, false, 10),
			11: ByteField("ReAuthenticationEnabled", UnsignedIntegerAttributeType, 0x0020, 0, mapset.NewSetWith(Read), false, true, false, 11),
			12: ByteField("KeyTransmissionEnabled", UnsignedIntegerAttributeType, 0x0010, 0, mapset.NewSetWith(Read, Write), false, true, false, 12),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
		Alarms: AlarmMap{
			0: "dot1x local authentication - allowed",
			1: "dot1x local authentication - denied",
		},
	}
}

// NewDot1XPortExtensionPackage (class ID 290) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewDot1XPortExtensionPackage(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*dot1xportextensionpackageBME, params...)
}
