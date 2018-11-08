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

type Dot1XPortExtensionPackage struct {
	BaseManagedEntity
}

func NewDot1XPortExtensionPackage(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		Name:     "Dot1XPortExtensionPackage",
		ClassID:  290,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
		},
		AttributeMask: 0,
		Attributes: []IAttribute{
			NewUint16Field("ManagedEntityId:", 0, Read),
			NewByteField("Dot1XEnable", 0, Read|Write),
			NewByteField("ActionRegister", 0, Write),
			NewByteField("AuthenticatorPaeState", 0, Read),
			NewByteField("BackendAuthenticationState", 0, Read),
			NewByteField("AdminControlledDirections", 0, Read|Write),
			NewByteField("OperationalControlledDirections", 0, Read),
			NewByteField("AuthenticatorControlledPortStatus", 0, Read),
			NewUint16Field("QuietPeriod", 0, Read|Write),
			NewUint16Field("ServerTimeoutPeriod", 0, Read|Write),
			NewUint16Field("ReAuthenticationPeriod", 0, Read),
			NewByteField("ReAuthenticationEnabled", 0, Read),
			NewByteField("KeyTransmissionEnabled", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &Dot1XPortExtensionPackage{entity}, nil
}
