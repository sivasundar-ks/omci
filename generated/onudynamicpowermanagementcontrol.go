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
 *              https://github.com/cboling/OMCI-parser
 */
package generated

type OnuDynamicPowerManagementControl struct {
	BaseManagedEntity
}

func NewOnuDynamicPowerManagementControl(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		Name:     "OnuDynamicPowerManagementControl",
		ClassID:  336,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
		},
		AttributeMask: 0,
		Attributes: []IAttribute{
			NewUint16Field("ManagedEntityId", 0, Read),
			NewByteField("PowerReductionManagementCapability", 0, Read),
			NewByteField("PowerReductionManagementMode", 0, Read|Write),
			NewUint16Field("Itransinit", 0, Read),
			NewUint16Field("Itxinit", 0, Read),
			NewUint32Field("MaximumSleepInterval", 0, Read|Write),
			NewUint32Field("MaximumReceiverOffInterval", 0, Read|Write),
			NewUint32Field("MinimumAwareInterval", 0, Read|Write),
			NewUint16Field("MinimumActiveHeldInterval", 0, Read|Write),
			NewUint64Field("MaximumSleepIntervalExtension", 0, Read|Write),
			NewByteField("EthernetPassiveOpticalNetworkEponCapabilityExtension", 0, Read),
			NewByteField("EponSetupExtension", 0, Read|Write),
			NewUint32Field("MissingConsecutiveBurstsThreshold", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &OnuDynamicPowerManagementControl{entity}, nil
}
