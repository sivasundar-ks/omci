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

// SoftwareImage (class ID #7) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type SoftwareImage struct {
	BaseManagedEntityDefinition
}

// NewSoftwareImage (class ID 7 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewSoftwareImage(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "SoftwareImage",
		ClassID:  7,
		EntityID: eid,
		MessageTypes: []MsgType{
			Get,
			DownloadSection,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: []*AttributeDefinition{
			Uint16Field("ManagedEntityId", 0, Read),
			UnknownField("Version", 0, Read),
			ByteField("IsCommitted", 0, Read),
			ByteField("IsActive", 0, Read),
			ByteField("IsValid", 0, Read),
			UnknownField("ProductCode", 0, Read),
			UnknownField("ImageHash", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &SoftwareImage{entity}, nil
}
