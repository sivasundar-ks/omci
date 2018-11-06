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

type SipUserData struct {
	BaseManagedEntity
}

func NewSipUserData(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		name:     "SipUserData",
		classID:  153,
		entityID: eid,
		msgTypes: []omci.MsgType{
			omci.Set,
			omci.Get,
			omci.Create,
			omci.Delete,
		},
		attributeList: []omci.IAttribute{
			omci.NewUint16Field("ManagedEntityId", 0, omci.Read|omci.SetByCreate),
			omci.NewUint16Field("SipAgentPointer", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint16Field("UserPartAor", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUnknownField("SipDisplayName", 0, omci.Read|omci.Write),
			omci.NewUint16Field("UsernameAndPassword", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint16Field("VoicemailServerSipUri", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint32Field("VoicemailSubscriptionExpirationTime", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint16Field("NetworkDialPlanPointer", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint16Field("ApplicationServicesProfilePointer", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint16Field("FeatureCodePointer", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint16Field("PptpPointer", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewByteField("ReleaseTimer", 0, omci.Read|omci.Write),
			omci.NewByteField("ReceiverOffHookRohTimer", 0, omci.Read|omci.Write),
		},
	}
	entity.computeAttributeMask()
	return &SipUserData{entity}, nil
}
