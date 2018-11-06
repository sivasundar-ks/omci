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

import (
	"../../omci"
)

type XdslChannelConfigurationProfile struct {
	BaseManagedEntity
}

func NewXdslChannelConfigurationProfile(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		name:     "XdslChannelConfigurationProfile",
		classID:  107,
		entityID: eid,
		msgTypes: []omci.MsgType{
			omci.Set,
			omci.Get,
			omci.Create,
			omci.Delete,
		},
		attributeList: []omci.IAttribute{
			omci.NewUint16Field("ManagedEntityId", 0, omci.Read|omci.SetByCreate),
			omci.NewUint32Field("MinimumDataRate", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint32Field("MaximumDataRate", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewByteField("RateAdaptationRatio", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewByteField("MaximumInterleavingDelay", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint32Field("DataRateThresholdUpshift", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint32Field("DataRateThresholdDownshift", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint32Field("MinimumReservedDataRate", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint32Field("MinimumDataRateInLowPowerState", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewByteField("MinimumImpulseNoiseProtection", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewByteField("MaximumBitErrorRatio", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewByteField("MinimumImpulseNoiseProtection8Khz", 0, omci.Read|omci.Write),
			omci.NewByteField("MaximumDelayVariation", 0, omci.Read|omci.Write),
			omci.NewByteField("ChannelInitializationPolicySelection", 0, omci.Read|omci.Write),
			omci.NewUint32Field("MinimumSosBitRateDownstream", 0, omci.Read|omci.Write),
			omci.NewUint32Field("MinimumSosBitRateUpstream", 0, omci.Read|omci.Write),
		},
	}
	entity.computeAttributeMask()
	return &XdslChannelConfigurationProfile{entity}, nil
}
