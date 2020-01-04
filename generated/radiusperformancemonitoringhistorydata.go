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

// RadiusPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity Radius performance monitoring history data
const RadiusPerformanceMonitoringHistoryDataClassID ClassID = ClassID(293)

var radiusperformancemonitoringhistorydataBME *ManagedEntityDefinition

// RadiusPerformanceMonitoringHistoryData (class ID #293)
//	This ME collects performance statistics on an ONU's radius client, particularly as related to
//	its IEEE-802.1X operation.
//
//	Instances of this ME are created and deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an ONU.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID (namely 0), this ME is implicitly linked to an instance of a dot1X configuration
//			profile. (R, setbycreate) (mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 ME that
//			contains PM threshold values. Since no threshold value attribute number exceeds 7, a threshold
//			data 2 ME is optional. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Access_Request Packets Transmitted
//			Access-request packets transmitted: This attribute counts transmitted radius access-request
//			messages, including retransmissions. (R) (mandatory) (4-bytes)
//
//		Access_Request Retransmission Count
//			Access-request retransmission count: This attribute counts radius access-request
//			retransmissions. (R) (mandatory) (4-bytes)
//
//		Access_Challenge Packets Received
//			Access-challenge packets received: This attribute counts received radius access-challenge
//			messages. (R) (mandatory) (4-bytes)
//
//		Access_Accept Packets Received
//			Access-accept packets received: This attribute counts received radius access-accept messages.
//			(R) (mandatory) (4-bytes)
//
//		Access_Reject Packets Received
//			Access-reject packets received: This attribute counts received radius access-reject messages.
//			(R) (mandatory) (4-bytes)
//
//		Invalid Radius Packets Received
//			Invalid radius packets received: This attribute counts received invalid radius messages. (R)
//			(mandatory) (4-bytes)
//
type RadiusPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	radiusperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "RadiusPerformanceMonitoringHistoryData",
		ClassID: 293,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xff00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 0),
			1: ByteField("IntervalEndTime", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, 1),
			2: Uint16Field("ThresholdData12Id", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 2),
			3: Uint32Field("AccessRequestPacketsTransmitted", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 3),
			4: Uint32Field("AccessRequestRetransmissionCount", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 4),
			5: Uint32Field("AccessChallengePacketsReceived", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 5),
			6: Uint32Field("AccessAcceptPacketsReceived", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 6),
			7: Uint32Field("AccessRejectPacketsReceived", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 7),
			8: Uint32Field("InvalidRadiusPacketsReceived", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 8),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewRadiusPerformanceMonitoringHistoryData (class ID 293) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewRadiusPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*radiusperformancemonitoringhistorydataBME, params...)
}
