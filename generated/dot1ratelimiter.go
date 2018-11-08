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

type Dot1RateLimiter struct {
	BaseManagedEntity
}

func NewDot1RateLimiter(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		Name:     "Dot1RateLimiter",
		ClassID:  298,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
			Create,
			Delete,
		},
		AttributeMask: 0,
		Attributes: []IAttribute{
			NewUint16Field("ManagedEntityId", 0, Read|SetByCreate),
			NewUint16Field("ParentMePointer", 0, Read|Write|SetByCreate),
			NewByteField("TpType", 0, Read|Write|SetByCreate),
			NewUint16Field("UpstreamUnicastFloodRatePointer", 0, Read|Write|SetByCreate),
			NewUint16Field("UpstreamBroadcastRatePointer", 0, Read|Write|SetByCreate),
			NewUint16Field("UpstreamMulticastPayloadRatePointer", 0, Read|Write|SetByCreate),
		},
	}
	entity.computeAttributeMask()
	return &Dot1RateLimiter{entity}, nil
}
