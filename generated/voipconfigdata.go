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

type VoipConfigData struct {
	BaseManagedEntity
}

func NewVoipConfigData(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		Name:     "VoipConfigData",
		ClassID:  138,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
		},
		AttributeMask: 0,
		Attributes: []IAttribute{
			NewUint16Field("ManagedEntityId", 0, Read),
			NewByteField("AvailableSignallingProtocols", 0, Read),
			NewByteField("SignallingProtocolUsed", 0, Read|Write),
			NewUint32Field("AvailableVoipConfigurationMethods", 0, Read),
			NewByteField("VoipConfigurationMethodUsed", 0, Read|Write),
			NewUint16Field("VoipConfigurationAddressPointer", 0, Read|Write),
			NewByteField("VoipConfigurationState", 0, Read),
			NewByteField("RetrieveProfile", 0, Write),
			NewUnknownField("ProfileVersion", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &VoipConfigData{entity}, nil
}
