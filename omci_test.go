/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package omci

import (
	me "./generated"
	"encoding/hex"
	"fmt"
	"github.com/google/gopacket"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func stringToPacket(input string) ([]byte, error) {
	var p []byte

	p, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return p, nil
}

func packetToString(input []byte) string {
	return strings.ToLower(hex.EncodeToString(input))
}

func getSbcMask(meDefinition me.IManagedEntityDefinition) uint16 {
	var sbcMask uint16

	for index, attr := range meDefinition.GetAttributeDefinitions() {
		if me.SupportsAttributeAccess(attr, me.SetByCreate) {
			if index == 0 {
				continue	// Skip Entity ID
			}
			sbcMask |= 1 << (15 - uint(index - 1))
		}
	}
	return sbcMask
}

// MibResetRequestTest tests decode/encode of a MIB Reset Request
func TestMibResetRequest(t *testing.T) {
	mibResetRequest := "00014F0A000200000000000000000000" +
		"00000000000000000000000000000000" +
		"000000000000000000000028"

	data, err := stringToPacket(mibResetRequest)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.TransactionID, uint16(1))
	assert.Equal(t, omciMsg.MessageType, byte(me.MibReset)|me.AR)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeMibResetRequest)
	assert.NotNil(t, msgLayer)

	omciMsg2, ok2 := msgLayer.(*MibResetRequest)
	assert.True(t, ok2)
	assert.Equal(t, omciMsg2.EntityClass, me.OnuDataClassId)
	assert.Equal(t, omciMsg2.EntityInstance, uint16(0))

	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, options, omciMsg, omciMsg2)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t, strings.ToLower(mibResetRequest), reconstituted)
}

func TestCreateGalEthernetProfile(t *testing.T) {
	createGalEthernetProfile := "0002440A011000010030000000000000" +
		"00000000000000000000000000000000" +
		"000000000000000000000028"

	data, err := stringToPacket(createGalEthernetProfile)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.TransactionID, uint16(2))
	assert.Equal(t, omciMsg.MessageType, byte(me.Create)|me.AR)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeCreateRequest)
	assert.NotNil(t, msgLayer)

	omciMsg2, ok2 := msgLayer.(*CreateRequest)
	assert.True(t, ok2)
	assert.Equal(t, omciMsg2.EntityClass, me.GalEthernetProfileClassId)
	assert.Equal(t, omciMsg2.EntityInstance, uint16(1))

	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, options, omciMsg, omciMsg2)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t,  strings.ToLower(createGalEthernetProfile), reconstituted)
}

func TestSetTCont(t *testing.T) {
	setTCont := "0003480A010680008000040000000000" +
		"00000000000000000000000000000000" +
		"000000000000000000000028"

	data, err := stringToPacket(setTCont)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.TransactionID, uint16(3))
	assert.Equal(t, omciMsg.MessageType, byte(me.Set)|me.AR)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeSetRequest)
	assert.NotNil(t, msgLayer)

	omciMsg2, ok2 := msgLayer.(*SetRequest)
	assert.True(t, ok2)
	assert.Equal(t, omciMsg2.EntityClass, me.TContClassId)
	assert.Equal(t, omciMsg2.EntityInstance, uint16(0x8000))

	attributes := omciMsg2.Attributes
	assert.Equal(t, len(attributes), 1)

	// TODO: Create generic test to look up the name from definition
	// Here 1 is the index in the attribute definition map of a TCONT that points
	// to the AllocID attribute.
	assert.Equal(t, attributes[1].Name, "AllocId")
	assert.Equal(t, attributes[1].Value, uint16(1024))

	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, options, omciMsg, omciMsg2)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t,  strings.ToLower(setTCont), reconstituted)
}

func TestCreate8021pMapperService_profile(t *testing.T) {
	create8021pMapperServiceProfile := "0007440A00828000ffffffffffffffff" +
		"ffffffffffffffffffff000000000000" +
		"000000000000000000000028"

	data, err := stringToPacket(create8021pMapperServiceProfile)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.TransactionID, uint16(7))
	assert.Equal(t, omciMsg.MessageType, byte(me.Create)|me.AR)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeCreateRequest)
	assert.NotNil(t, msgLayer)

	createRequest, ok2 := msgLayer.(*CreateRequest)
	assert.True(t, ok2)
	assert.Equal(t, createRequest.EntityClass, me.Ieee8021PMapperServiceProfileClassId)
	assert.Equal(t, createRequest.EntityInstance, uint16(0x8000))

	attributes := createRequest.Attributes
	assert.NotNil(t, attributes)
	assert.Equal(t, len(attributes), 12)

	for index := uint(1); index <= uint(9); index++ {
		value, err2 := attributes[index].GetValue()
		assert.Nil(t, err2)

		value16, ok3 := value.(uint16)
		assert.True(t, ok3)
		assert.Equal(t, value16, uint16(0xffff))
	}
	// As this is a create request, gather up all set-by-create attributes
	// make sure we got them all, and nothing else
	meDefinition, err := me.LoadManagedEntityDefinition(createRequest.EntityClass)
	assert.Nil(t, err)

	sbcMask := getSbcMask(meDefinition)
	for index := 1; index <= 16; index++ {
		if sbcMask & uint16(1 << (uint)(16 - index)) != 0 {
			_, ok3 := attributes[uint(index)]
			assert.True(t, ok3)
		} else {
			_, ok3 := attributes[uint(index)]
			assert.False(t, ok3)
		}
	}
	// TODO: Individual attribute tests here if needed
	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, options, omciMsg, createRequest)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t,  strings.ToLower(create8021pMapperServiceProfile), reconstituted)
}

func TestCreate_macBridgeService_profile(t *testing.T) {
	var createMacBridgeServiceProfile = "000B440A002D02010001008000140002" +
		"000f0001000000000000000000000000" +
		"000000000000000000000028"

	data, err := stringToPacket(createMacBridgeServiceProfile)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.TransactionID, uint16(0xb))
	assert.Equal(t, omciMsg.MessageType, byte(me.Create)|me.AR)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeCreateRequest)
	assert.NotNil(t, msgLayer)

	createRequest, ok2 := msgLayer.(*CreateRequest)
	assert.True(t, ok2)
	assert.Equal(t, createRequest.EntityClass, me.MacBridgeServiceProfileClassId)
	assert.Equal(t, createRequest.EntityInstance, uint16(0x201))

	attributes := createRequest.Attributes
	assert.NotNil(t, attributes)

	// As this is a create request, gather up all set-by-create attributes
	// make sure we got them all, and nothing else
	meDefinition, err := me.LoadManagedEntityDefinition(createRequest.EntityClass)
	assert.Nil(t, err)

	sbcMask := getSbcMask(meDefinition)
	for index := 1; index <= 16; index++ {
		if sbcMask & uint16(1 << (uint)(16 - index)) != 0 {
			_, ok3 := attributes[uint(index)]
			assert.True(t, ok3)
		} else {
			_, ok3 := attributes[uint(index)]
			assert.False(t, ok3)
		}
	}
	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, options, omciMsg, createRequest)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t,  strings.ToLower(createMacBridgeServiceProfile), reconstituted)
}

func TestCreateGemPortNetworkCtp(t *testing.T) {
	createGemPortNetworkCtp := "000C440A010C01000400800003010000" +
		"00000000000000000000000000000000" +
		"000000000000000000000028"

	data, err := stringToPacket(createGemPortNetworkCtp)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.TransactionID, uint16(0xc))
	assert.Equal(t, omciMsg.MessageType, byte(me.Create)|me.AR)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeCreateRequest)
	assert.NotNil(t, msgLayer)

	createRequest, ok2 := msgLayer.(*CreateRequest)
	assert.True(t, ok2)
	assert.Equal(t, createRequest.EntityClass, me.GemPortNetworkCtpClassId)
	assert.Equal(t, createRequest.EntityInstance, uint16(0x100))

	attributes := createRequest.Attributes
	assert.NotNil(t, attributes)

	// As this is a create request, gather up all set-by-create attributes
	// make sure we got them all, and nothing else
	meDefinition, err := me.LoadManagedEntityDefinition(createRequest.EntityClass)
	assert.Nil(t, err)

	sbcMask := getSbcMask(meDefinition)
	for index := 1; index <= 16; index++ {
		if sbcMask & uint16(1 << (uint)(16 - index)) != 0 {
			_, ok3 := attributes[uint(index)]
			assert.True(t, ok3)
		} else {
			_, ok3 := attributes[uint(index)]
			assert.False(t, ok3)
		}
	}
	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, options, omciMsg, createRequest)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t,  strings.ToLower(createGemPortNetworkCtp), reconstituted)
}

// TODO: Uncomment as encode/decode supported
//func TestMulticastGemInterworkingTp(t *testing.T) {
//
//	multicastGemInterworkingTp := "0011440A011900060104000001000000" +
//		"00000000000000000000000000000000" +
//		"000000000000000000000028"
//
//	data, err := stringToPacket(multicastGemInterworkingTp)
//	assert.NoError(t, err)
//
//	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
//	fmt.Println(packet)
//
//	customLayer := packet.Layer(LayerTypeOMCI)
//	assert.NotNil(t, customLayer)
//}
//
//func TestCreateGemInteworkingTp(t *testing.T) {
//
//	createGemInteworkingTp := "0012440A010A80010100058000000000" +
//		"01000000000000000000000000000000" +
//		"000000000000000000000028"
//
//	data, err := stringToPacket(createGemInteworkingTp)
//	assert.NoError(t, err)
//
//	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
//	fmt.Println(packet)
//
//	customLayer := packet.Layer(LayerTypeOMCI)
//	assert.NotNil(t, customLayer)
//}

func TestSet8021pMapperServiceProfile(t *testing.T) {
	set8021pMapperServiceProfile := "0016480A008280004000800100000000" +
		"00000000000000000000000000000000" +
		"000000000000000000000028"

	data, err := stringToPacket(set8021pMapperServiceProfile)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.TransactionID, uint16(0x16))
	assert.Equal(t, omciMsg.MessageType, byte(me.Set)|me.AR)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeSetRequest)
	assert.NotNil(t, msgLayer)

	setRequest, ok2 := msgLayer.(*SetRequest)
	assert.True(t, ok2)
	assert.Equal(t, setRequest.EntityClass, me.Ieee8021PMapperServiceProfileClassId)
	assert.Equal(t, setRequest.EntityInstance, uint16(0x8000))

	attributes := setRequest.Attributes
	assert.NotNil(t, attributes)
	assert.Equal(t, len(attributes), 1)

	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, options, omciMsg, setRequest)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t,  strings.ToLower(set8021pMapperServiceProfile), reconstituted)
}
// TODO: Uncomment as encode/decode supported
//func TestCreateMacBridgePortConfigurationData(t *testing.T) {
//
//	createMacBridgePortConfigurationData := "001A440A002F21010201020380000000" +
//		"00000000000000000000000000000000" +
//		"000000000000000000000028"
//
//	data, err := stringToPacket(createMacBridgePortConfigurationData)
//	assert.NoError(t, err)
//
//	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
//	fmt.Println(packet)
//
//	customLayer := packet.Layer(LayerTypeOMCI)
//	assert.NotNil(t, customLayer)
//}
//
//func TestCreateVlanTaggingFilterData(t *testing.T) {
//
//	createVlanTaggingFilterData := "001F440A005421010400000000000000" +
//		"00000000000000000000000000000000" +
//		"100100000000000000000028"
//
//	data, err := stringToPacket(createVlanTaggingFilterData)
//	assert.NoError(t, err)
//
//	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
//	fmt.Println(packet)
//
//	customLayer := packet.Layer(LayerTypeOMCI)
//	assert.NotNil(t, customLayer)
//}
//
//func TestCreateExtendedVlanTaggingOperationConfigurationData(t *testing.T) {
//
//	createExtendedVlanTaggingOperationConfigurationData := "0023440A00AB02020A04010000000000" +
//		"00000000000000000000000000000000" +
//		"000000000000000000000028"
//
//	data, err := stringToPacket(createExtendedVlanTaggingOperationConfigurationData)
//	assert.NoError(t, err)
//
//	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
//	fmt.Println(packet)
//}
//
//func TestSetExtendedVlanTagging_operationConfigurationData(t *testing.T) {
//
//	setExtendedVlanTaggingOperationConfigurationData := "0024480A00AB02023800810081000000" +
//		"00000000000000000000000000000000" +
//		"000000000000000000000028"
//
//	data, err := stringToPacket(setExtendedVlanTaggingOperationConfigurationData)
//	assert.NoError(t, err)
//
//	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
//	fmt.Println(packet)
//
//	customLayer := packet.Layer(LayerTypeOMCI)
//	assert.NotNil(t, customLayer)
//}
//
//func TestSetExtendedVlanTagging1(t *testing.T) {
//
//	setExtendedVlanTagging1 := "0025480A00AB02020400f00000008200" +
//		"5000402f000000082004000000000000" +
//		"000000000000000000000028"
//
//	data, err := stringToPacket(setExtendedVlanTagging1)
//	assert.NoError(t, err)
//
//	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
//	fmt.Println(packet)
//
//	customLayer := packet.Layer(LayerTypeOMCI)
//	assert.NotNil(t, customLayer)
//}
//
//func TestSetExtendedVlanTagging2(t *testing.T) {
//
//	setExtendedVlanTagging2 := "0026480A00AB02020400F00000008200" +
//		"d000402f00000008200c000000000000" +
//		"000000000000000000000028"
//
//	data, err := stringToPacket(setExtendedVlanTagging2)
//	assert.NoError(t, err)
//
//	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
//	fmt.Println(packet)
//
//	customLayer := packet.Layer(LayerTypeOMCI)
//	assert.NotNil(t, customLayer)
//}
//
//func TestCreateMacBridgePortConfigurationData2(t *testing.T) {
//
//	createMacBridgePortConfigurationData2 := "0029440A002F02010201010b04010000" +
//		"00000000000000000000000000000000" +
//		"000000000000000000000028"
//
//	data, err := stringToPacket(createMacBridgePortConfigurationData2)
//	assert.NoError(t, err)
//
//	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
//	fmt.Println(packet)
//
//	customLayer := packet.Layer(LayerTypeOMCI)
//	assert.NotNil(t, customLayer)
//}

func TestMibUpload(t *testing.T) {
	mibUpload := "00304D0A000200000000000000000000" +
		"00000000000000000000000000000000" +
		"000000000000000000000028"

	data, err := stringToPacket(mibUpload)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.TransactionID, uint16(0x30))
	assert.Equal(t, omciMsg.MessageType, byte(me.MibUpload)|me.AR)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeMibUploadRequest)
	assert.NotNil(t, msgLayer)

	uploadRequest, ok2 := msgLayer.(*MibUploadRequest)
	assert.True(t, ok2)
	assert.Equal(t, uploadRequest.EntityClass, me.OnuDataClassId)
	assert.Equal(t, uploadRequest.EntityInstance, uint16(0))

	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, options, omciMsg, uploadRequest)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t,  strings.ToLower(mibUpload), reconstituted)
}

// TODO: Uncomment as encode/decode supported
//func TestEnhSecurityAvc(t *testing.T) {
//
//	enhSecurityAvc := "0000110a014c0000008000202020202020202020202020202020202020202020" +
//		"2020202020202020000000280be43cf4"
//
//	data, err := stringToPacket(enhSecurityAvc)
//	assert.NoError(t, err)
//
//	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
//	fmt.Println(packet)
//
//	customLayer := packet.Layer(LayerTypeOMCI)
//	assert.NotNil(t, customLayer)
//}

func TestAlarmMessage(t *testing.T) {
	alarmMessage := "0000100a00050101000000000000000000000000000000000000000000000000" +
		"0000000220000000000000280be43cf4"

	data, err := stringToPacket(alarmMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.TransactionID, uint16(0))
	assert.Equal(t, omciMsg.MessageType, byte(me.AlarmNotification))
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeAlarmNotification)
	assert.Nil(t, msgLayer)		// TODO: Fix decode

	//assert.NotNil(t, msgLayer)
	//
	//alarmNotification, ok2 := msgLayer.(*AlarmNotificationMsg)
	//assert.True(t, ok2)
	//// TODO: Repace with actual entity class
	//assert.Equal(t, alarmNotification.EntityClass, uint16(0x0005))
	//assert.Equal(t, alarmNotification.EntityInstance, uint16(0x101))
	// TODO: Decode alarm bits

	// TODO: Serialize frame and test with original
}

func TestOnuRebootRequest(t *testing.T) {
	onuRebootRequest := "0016590a01000000000000000000000000000" +
		"0000000000000000000000000000000000000" +
		"00000000000028"

	data, err := stringToPacket(onuRebootRequest)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.TransactionID, uint16(0x16))
	assert.Equal(t, omciMsg.MessageType, byte(me.Reboot)|me.AR)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeRebootRequest)
	assert.Nil(t, msgLayer)		// TODO: Fix decode

	//assert.NotNil(t, msgLayer)
	//
	//rebootRequest, ok2 := msgLayer.(*RebootRequest)
	//assert.True(t, ok2)
	//assert.Equal(t, rebootRequest.EntityClass, me.OnuDataClassId)
	//assert.Equal(t, rebootRequest.EntityInstance, uint16(0x8000))
	
	// TODO: Test Decoded flags

	// TODO: Serialize frame and test with original
}

func TestMibUploadNextSequence(t *testing.T) {
	mibUploadNextSequence := [...]string{
		"00032e0a0002000000020000800000000000000000000000000000000000000000000000000000000000002828ce00e2",
		"00042e0a0002000000050101f0002f2f05202020202020202020202020202020202020202000000000000028d4eb4bdf",
		"00052e0a00020000000501010f802020202020202020202020202020202020202020000000000000000000282dbe4b44",
		"00062e0a0002000000050104f000303001202020202020202020202020202020202020202000000000000028ef1b035b",
		"00072e0a00020000000501040f80202020202020202020202020202020202020202000000000000000000028fec29135",
		"00082e0a0002000000050180f000f8f801202020202020202020202020202020202020202000000000000028fd4e0b07",
		"00092e0a00020000000501800f802020202020202020202020202020202020202020000000000000000000283306b3c0",
		"000a2e0a0002000000060101f0002f054252434d12345678000000000000000000000000000c000000000028585c2083",
		"000b2e0a00020000000601010f004252434d00000000000000000000000000000000000000000000000000284f0e82b9",
		"000c2e0a000200000006010100f8202020202020202020202020202020202020202000000000000000000028e68bdb63",
		"000d2e0a0002000000060101000400000000000000000000000000000000000000000000000000000000002857bc2730",
		"000e2e0a0002000000060104f00030014252434d12345678000000000000000000000000000c000000000028afe656f5",
		"000f2e0a00020000000601040f004252434d0000000000000000000000000000000000000000000000000028f8f6db74",
		"00102e0a000200000006010400f8202020202020202020202020202020202020202000000800000000000028064fc177",
		"00112e0a000200000006010400040000000000000000000000000000000000000000000000000000000000285a5c0841",
		"00122e0a0002000000060180f000f8014252434d12345678000000000000000000000000000c0000000000286826eafe",
		"00132e0a00020000000601800f004252434d00000000000000000000000000000000000000000000000000281c4b7033",
		"00142e0a000200000006018000f8202020202020202020202020202020202020202000084040000000000028ac144eb3",
		"00152e0a000200000006018000040000000000000000000000000000000000000000000000000000000000280a81a9a7",
		"00162e0a0002000000070000f0003530323247574f32363632303033010101000000000000000000000000287ea42d51",
		"00172e0a0002000000070001f0003530323247574f3236363230303300000100000000000000000000000028b17f567f",
		"00182e0a0002000000830000c0002020202020202020202020202020202020202020202020200000000000280e7eebaa",
		"00192e0a00020000008300002000202020202020202020202020202000000000000000000000000000000028a95c03b3",
		"001a2e0a00020000008300001000000000000000000000000000000000000000000000000000000000000028f30515a1",
		"001b2e0a0002000000850000ffe0000000000000000000000000000000000000000000000000000000000028764c18de",
		"001c2e0a0002000000860001c00000001018aaaa000000000000000000000000000000000000000000000028ea220ce0",
		"001d2e0a00020000008600012000000000000000000000000000000000000000000000000000000000000028fbdb571a",
		"001e2e0a00020000008600011f80000000000000000000000000000000000000000000000000000000000028c2682282",
		"001f2e0a000200000086000100780000000000000000000000000000000000000000000000000000000000289c4809b1",
		"00202e0a00020000008600010004000000000000000000000000000000000000000000000000000000000028d174a7d6",
		"00212e0a000200000086000100020000000000000000000000000000000000000000000000000000000000288f353976",
		"00222e0a0002000001000000e0004252434d00000000000000000000000000004252434d123456780000002803bbceb6",
		"00232e0a00020000010000001f800000000000000000000000000000000000000000000000000000000000281b9674db",
		"00242e0a00020000010000000040000000000000000000000000000000000000000000000000000000000028b1050b9b",
		"00252e0a000200000100000000380000000000000000000000000000030000000000000000000000000000288266645e",
		"00262e0a0002000001010000f80042564d344b3030425241303931352d3030383300b3000001010000000028837d624f",
		"00272e0a000200000101000007f8000000010020027c8563001600003000000000000000000000000000002896c707e1",
		"00282e0a0002000001068000e00000ff0101000000000000000000000000000000000000000000000000002811acb324",
		"00292e0a0002000001068001e00000ff0101000000000000000000000000000000000000000000000000002823ad6aa9",
		"002a2e0a0002000001068002e00000ff01010000000000000000000000000000000000000000000000000028a290efd9",
		"002b2e0a0002000001068003e00000ff01010000000000000000000000000000000000000000000000000028af893357",
		"002c2e0a0002000001068004e00000ff01010000000000000000000000000000000000000000000000000028901141a3",
		"002d2e0a0002000001068005e00000ff01010000000000000000000000000000000000000000000000000028c4398bcc",
		"002e2e0a0002000001068006e00000ff01010000000000000000000000000000000000000000000000000028e60acd99",
		"002f2e0a0002000001068007e00000ff010100000000000000000000000000000000000000000000000000284b5faf23",
		"00302e0a0002000001078001ffff01000800300000050900000000ffff000000008181000000000000000028bef89455",
		"00312e0a0002000001080401f0000000000004010000000000000000000000000000000000000000000000287dc5183d",
		"00322e0a0002000001150401fff0000080008000000000040100000000010000000000000000000000000028cc0a46a9",
		"00332e0a0002000001150401000f0200020002000200ffff09000000000000000000000000000000000000288c42acdd",
		"00342e0a0002000001150402fff0000080008000000000040100010000010000000000000000000000000028de9f625a",
		"00352e0a0002000001150402000f0200020002000200ffff09000000000000000000000000000000000000280587860b",
		"00362e0a0002000001150403fff0000080008000000000040100020000010000000000000000000000000028a49cc820",
		"00372e0a0002000001150403000f0200020002000200ffff0900000000000000000000000000000000000028b4e4a2b9",
		"00382e0a0002000001150404fff00000800080000000000401000300000100000000000000000000000000288233147b",
		"00392e0a0002000001150404000f0200020002000200ffff090000000000000000000000000000000000002881b706b0",
		"003a2e0a0002000001150405fff0000080008000000000040100040000010000000000000000000000000028be8efc9f",
		"003b2e0a0002000001150405000f0200020002000200ffff0900000000000000000000000000000000000028d944804b",
		"003c2e0a0002000001150406fff0000080008000000000040100050000010000000000000000000000000028725c3864",
		"003d2e0a0002000001150406000f0200020002000200ffff09000000000000000000000000000000000000284e2d5cd2",
		"003e2e0a0002000001150407fff0000080008000000000040100060000010000000000000000000000000028464b03ba",
		"003f2e0a0002000001150407000f0200020002000200ffff09000000000000000000000000000000000000287006cfd0",
		"00402e0a0002000001150408fff0000080008000000000040100070000010000000000000000000000000028cd88ebeb",
		"00412e0a0002000001150408000f0200020002000200ffff09000000000000000000000000000000000000285a5905e2",
		"00422e0a0002000001158000fff0000100010000000000800000000000010000000000000000000000000028e61b19d1",
		"00432e0a0002000001158000000f0200020002000200ffff0900000000000000000000000000000000000028b0cc5937",
		"00442e0a0002000001158001fff00001000100000000008000000100000100000000000000000000000000285386bbf2",
		"00452e0a0002000001158001000f0200020002000200ffff0900000000000000000000000000000000000028c06723ab",
		"00462e0a0002000001158002fff0000100010000000000800000020000010000000000000000000000000028ab49704a",
		"00472e0a0002000001158002000f0200020002000200ffff090000000000000000000000000000000000002857432f25",
		"00482e0a0002000001158003fff0000100010000000000800000030000010000000000000000000000000028b383c057",
		"00492e0a0002000001158003000f0200020002000200ffff0900000000000000000000000000000000000028dca40d66",
		"004a2e0a0002000001158004fff00001000100000000008000000400000100000000000000000000000000286b7ba0e2",
		"004b2e0a0002000001158004000f0200020002000200ffff0900000000000000000000000000000000000028fd442363",
		"004c2e0a0002000001158005fff00001000100000000008000000500000100000000000000000000000000280ee9a0b8",
		"004d2e0a0002000001158005000f0200020002000200ffff0900000000000000000000000000000000000028bc1b9843",
		"004e2e0a0002000001158006fff00001000100000000008000000600000100000000000000000000000000280c535114",
		"004f2e0a0002000001158006000f0200020002000200ffff090000000000000000000000000000000000002887032f2b",
		"00502e0a0002000001158007fff0000100010000000000800000070000010000000000000000000000000028a77d7f61",
		"00512e0a0002000001158007000f0200020002000200ffff090000000000000000000000000000000000002835e9f567",
		"00522e0a0002000001158008fff0000100010000000000800100000000010000000000000000000000000028ff4ca94b",
		"00532e0a0002000001158008000f0200020002000200ffff09000000000000000000000000000000000000281e2f1e33",
		"00542e0a0002000001158009fff00001000100000000008001000100000100000000000000000000000000283c473db0",
		"00552e0a0002000001158009000f0200020002000200ffff090000000000000000000000000000000000002869f51dda",
		"00562e0a000200000115800afff0000100010000000000800100020000010000000000000000000000000028046b8feb",
		"00572e0a000200000115800a000f0200020002000200ffff090000000000000000000000000000000000002868b1495e",
		"00582e0a000200000115800bfff00001000100000000008001000300000100000000000000000000000000282b927566",
		"00592e0a000200000115800b000f0200020002000200ffff0900000000000000000000000000000000000028cd43de96",
		"005a2e0a000200000115800cfff0000100010000000000800100040000010000000000000000000000000028c49617dd",
		"005b2e0a000200000115800c000f0200020002000200ffff0900000000000000000000000000000000000028fbbb972a",
		"005c2e0a000200000115800dfff000010001000000000080010005000001000000000000000000000000002893d4c2b5",
		"005d2e0a000200000115800d000f0200020002000200ffff0900000000000000000000000000000000000028dc9d97ca",
		"005e2e0a000200000115800efff00001000100000000008001000600000100000000000000000000000000280e1ec245",
		"005f2e0a000200000115800e000f0200020002000200ffff0900000000000000000000000000000000000028be3d56f1",
		"00602e0a000200000115800ffff00001000100000000008001000700000100000000000000000000000000280c046099",
		"00612e0a000200000115800f000f0200020002000200ffff0900000000000000000000000000000000000028d770e4ea",
		"00622e0a0002000001158010fff00001000100000000008002000000000100000000000000000000000000281b449092",
		"00632e0a0002000001158010000f0200020002000200ffff09000000000000000000000000000000000000282b7a8604",
		"00642e0a0002000001158011fff0000100010000000000800200010000010000000000000000000000000028ad498068",
		"00652e0a0002000001158011000f0200020002000200ffff0900000000000000000000000000000000000028a114b304",
		"00662e0a0002000001158012fff0000100010000000000800200020000010000000000000000000000000028c091715d",
		"00672e0a0002000001158012000f0200020002000200ffff0900000000000000000000000000000000000028d4ab49e7",
		"00682e0a0002000001158013fff0000100010000000000800200030000010000000000000000000000000028e39dd5dd",
		"00692e0a0002000001158013000f0200020002000200ffff09000000000000000000000000000000000000288779ebf0",
		"006a2e0a0002000001158014fff0000100010000000000800200040000010000000000000000000000000028c47a741f",
		"006b2e0a0002000001158014000f0200020002000200ffff0900000000000000000000000000000000000028ce765fcd",
		"006c2e0a0002000001158015fff00001000100000000008002000500000100000000000000000000000000288f732591",
		"006d2e0a0002000001158015000f0200020002000200ffff0900000000000000000000000000000000000028920b6f5e",
		"006e2e0a0002000001158016fff0000100010000000000800200060000010000000000000000000000000028f072e1c3",
		"006f2e0a0002000001158016000f0200020002000200ffff0900000000000000000000000000000000000028b47ea00f",
		"00702e0a0002000001158017fff000010001000000000080020007000001000000000000000000000000002813461627",
		"00712e0a0002000001158017000f0200020002000200ffff090000000000000000000000000000000000002809013378",
		"00722e0a0002000001158018fff00001000100000000008003000000000100000000000000000000000000286168e200",
		"00732e0a0002000001158018000f0200020002000200ffff0900000000000000000000000000000000000028eccc81f7",
		"00742e0a0002000001158019fff000010001000000000080030001000001000000000000000000000000002855fe8072",
		"00752e0a0002000001158019000f0200020002000200ffff0900000000000000000000000000000000000028c159c496",
		"00762e0a000200000115801afff000010001000000000080030002000001000000000000000000000000002872652aca",
		"00772e0a000200000115801a000f0200020002000200ffff09000000000000000000000000000000000000283ba1c255",
		"00782e0a000200000115801bfff00001000100000000008003000300000100000000000000000000000000286b2ecb95",
		"00792e0a000200000115801b000f0200020002000200ffff0900000000000000000000000000000000000028441fbe05",
		"007a2e0a000200000115801cfff0000100010000000000800300040000010000000000000000000000000028f07ad5d8",
		"007b2e0a000200000115801c000f0200020002000200ffff0900000000000000000000000000000000000028237d6a28",
		"007c2e0a000200000115801dfff0000100010000000000800300050000010000000000000000000000000028e47dfdca",
		"007d2e0a000200000115801d000f0200020002000200ffff09000000000000000000000000000000000000280ca941be",
		"007e2e0a000200000115801efff00001000100000000008003000600000100000000000000000000000000283a1ef4d4",
		"007f2e0a000200000115801e000f0200020002000200ffff09000000000000000000000000000000000000289c905cd5",
		"00802e0a000200000115801ffff0000100010000000000800300070000010000000000000000000000000028384ae4c6",
		"00812e0a000200000115801f000f0200020002000200ffff0900000000000000000000000000000000000028be87eb55",
		"00822e0a0002000001158020fff0000100010000000000800400000000010000000000000000000000000028f0412282",
		"00832e0a0002000001158020000f0200020002000200ffff0900000000000000000000000000000000000028842ada0c",
		"00842e0a0002000001158021fff0000100010000000000800400010000010000000000000000000000000028a6eed1bc",
		"00852e0a0002000001158021000f0200020002000200ffff09000000000000000000000000000000000000280f3dd903",
		"00862e0a0002000001158022fff0000100010000000000800400020000010000000000000000000000000028474a0823",
		"00872e0a0002000001158022000f0200020002000200ffff0900000000000000000000000000000000000028e00456b3",
		"00882e0a0002000001158023fff000010001000000000080040003000001000000000000000000000000002851cbe1a6",
		"00892e0a0002000001158023000f0200020002000200ffff090000000000000000000000000000000000002869a99563",
		"008a2e0a0002000001158024fff000010001000000000080040004000001000000000000000000000000002867705534",
		"008b2e0a0002000001158024000f0200020002000200ffff09000000000000000000000000000000000000286f9570c0",
		"008c2e0a0002000001158025fff0000100010000000000800400050000010000000000000000000000000028450ef70e",
		"008d2e0a0002000001158025000f0200020002000200ffff090000000000000000000000000000000000002847588afa",
		"008e2e0a0002000001158026fff0000100010000000000800400060000010000000000000000000000000028c8218600",
		"008f2e0a0002000001158026000f0200020002000200ffff0900000000000000000000000000000000000028391a6ba7",
		"00902e0a0002000001158027fff0000100010000000000800400070000010000000000000000000000000028afc0878b",
		"00912e0a0002000001158027000f0200020002000200ffff090000000000000000000000000000000000002819130d66",
		"00922e0a0002000001158028fff00001000100000000008005000000000100000000000000000000000000289afa4cf7",
		"00932e0a0002000001158028000f0200020002000200ffff090000000000000000000000000000000000002873a4e20b",
		"00942e0a0002000001158029fff0000100010000000000800500010000010000000000000000000000000028633debd9",
		"00952e0a0002000001158029000f0200020002000200ffff09000000000000000000000000000000000000280397eb28",
		"00962e0a000200000115802afff00001000100000000008005000200000100000000000000000000000000280ed5ee7a",
		"00972e0a000200000115802a000f0200020002000200ffff0900000000000000000000000000000000000028f886ba59",
		"00982e0a000200000115802bfff000010001000000000080050003000001000000000000000000000000002888ff79b1",
		"00992e0a000200000115802b000f0200020002000200ffff090000000000000000000000000000000000002846baf278",
		"009a2e0a000200000115802cfff00001000100000000008005000400000100000000000000000000000000281fd1e68f",
		"009b2e0a000200000115802c000f0200020002000200ffff0900000000000000000000000000000000000028d99760f9",
		"009c2e0a000200000115802dfff0000100010000000000800500050000010000000000000000000000000028557aaf84",
		"009d2e0a000200000115802d000f0200020002000200ffff0900000000000000000000000000000000000028064210fd",
		"009e2e0a000200000115802efff00001000100000000008005000600000100000000000000000000000000285fd6c061",
		"009f2e0a000200000115802e000f0200020002000200ffff0900000000000000000000000000000000000028299efbb5",
		"00a02e0a000200000115802ffff000010001000000000080050007000001000000000000000000000000002834f127c4",
		"00a12e0a000200000115802f000f0200020002000200ffff0900000000000000000000000000000000000028edd30591",
		"00a22e0a0002000001158030fff0000100010000000000800600000000010000000000000000000000000028183183f2",
		"00a32e0a0002000001158030000f0200020002000200ffff0900000000000000000000000000000000000028a27e71f6",
		"00a42e0a0002000001158031fff0000100010000000000800600010000010000000000000000000000000028bd64dfc0",
		"00a52e0a0002000001158031000f0200020002000200ffff090000000000000000000000000000000000002839e2f37e",
		"00a62e0a0002000001158032fff00001000100000000008006000200000100000000000000000000000000283e72282e",
		"00a72e0a0002000001158032000f0200020002000200ffff0900000000000000000000000000000000000028cef19baa",
		"00a82e0a0002000001158033fff00001000100000000008006000300000100000000000000000000000000281c1caf44",
		"00a92e0a0002000001158033000f0200020002000200ffff090000000000000000000000000000000000002814712e27",
		"00aa2e0a0002000001158034fff0000100010000000000800600040000010000000000000000000000000028f02a30a4",
		"00ab2e0a0002000001158034000f0200020002000200ffff0900000000000000000000000000000000000028068fcbf5",
		"00ac2e0a0002000001158035fff0000100010000000000800600050000010000000000000000000000000028436bd783",
		"00ad2e0a0002000001158035000f0200020002000200ffff09000000000000000000000000000000000000288da3200f",
		"00ae2e0a0002000001158036fff0000100010000000000800600060000010000000000000000000000000028c26a02ca",
		"00af2e0a0002000001158036000f0200020002000200ffff0900000000000000000000000000000000000028147a41ee",
		"00b02e0a0002000001158037fff00001000100000000008006000700000100000000000000000000000000287c2bbec0",
		"00b12e0a0002000001158037000f0200020002000200ffff09000000000000000000000000000000000000284c86c11f",
		"00b22e0a0002000001158038fff000010001000000000080070000000001000000000000000000000000002895b94e06",
		"00b32e0a0002000001158038000f0200020002000200ffff0900000000000000000000000000000000000028a2b34012",
		"00b42e0a0002000001158039fff000010001000000000080070001000001000000000000000000000000002804b205a3",
		"00b52e0a0002000001158039000f0200020002000200ffff090000000000000000000000000000000000002886856d76",
		"00b62e0a000200000115803afff00001000100000000008007000200000100000000000000000000000000282a22752c",
		"00b72e0a000200000115803a000f0200020002000200ffff0900000000000000000000000000000000000028488e67db",
		"00b82e0a000200000115803bfff0000100010000000000800700030000010000000000000000000000000028a55f79ea",
		"00b92e0a000200000115803b000f0200020002000200ffff090000000000000000000000000000000000002842d77ba7",
		"00ba2e0a000200000115803cfff0000100010000000000800700040000010000000000000000000000000028da65268a",
		"00bb2e0a000200000115803c000f0200020002000200ffff0900000000000000000000000000000000000028c58443ec",
		"00bc2e0a000200000115803dfff0000100010000000000800700050000010000000000000000000000000028997aca59",
		"00bd2e0a000200000115803d000f0200020002000200ffff0900000000000000000000000000000000000028a2670b7d",
		"00be2e0a000200000115803efff000010001000000000080070006000001000000000000000000000000002813e904cb",
		"00bf2e0a000200000115803e000f0200020002000200ffff0900000000000000000000000000000000000028c387a9e5",
		"00c02e0a000200000115803ffff0000100010000000000800700070000010000000000000000000000000028d556a6b2",
		"00c12e0a000200000115803f000f0200020002000200ffff090000000000000000000000000000000000002868d9961a",
		"00c22e0a0002000001168000f000800000000200000000000000000000000000000000000000000000000028b69b53c1",
		"00c32e0a0002000001168001f000800000000200000000000000000000000000000000000000000000000028537705d4",
		"00c42e0a0002000001168002f000800000000200000000000000000000000000000000000000000000000028db171b7b",
		"00c52e0a0002000001168003f000800000000200000000000000000000000000000000000000000000000028f9b3fa54",
		"00c62e0a0002000001168004f000800000000200000000000000000000000000000000000000000000000028cdacda4e",
		"00c72e0a0002000001168005f00080000000020000000000000000000000000000000000000000000000002837133b6e",
		"00c82e0a0002000001168006f000800000000200000000000000000000000000000000000000000000000028d6447905",
		"00c92e0a0002000001168007f000800000000200000000000000000000000000000000000000000000000028021a3910",
		"00ca2e0a0002000001168008f00080010000020000000000000000000000000000000000000000000000002835d3cf43",
		"00cb2e0a0002000001168009f00080010000020000000000000000000000000000000000000000000000002887ad76fc",
		"00cc2e0a000200000116800af00080010000020000000000000000000000000000000000000000000000002895e3d838",
		"00cd2e0a000200000116800bf000800100000200000000000000000000000000000000000000000000000028a07489ac",
		"00ce2e0a000200000116800cf0008001000002000000000000000000000000000000000000000000000000285d08821d",
		"00cf2e0a000200000116800df000800100000200000000000000000000000000000000000000000000000028302249a4",
		"00d02e0a000200000116800ef0008001000002000000000000000000000000000000000000000000000000283966d3bc",
		"00d12e0a000200000116800ff0008001000002000000000000000000000000000000000000000000000000289519cdb5",
		"00d22e0a0002000001168010f0008002000002000000000000000000000000000000000000000000000000281bc99b7b",
		"00d32e0a0002000001168011f000800200000200000000000000000000000000000000000000000000000028e483b1a0",
		"00d42e0a0002000001168012f0008002000002000000000000000000000000000000000000000000000000286885d8bd",
		"00d52e0a0002000001168013f000800200000200000000000000000000000000000000000000000000000028cbe7afd8",
		"00d62e0a0002000001168014f00080020000020000000000000000000000000000000000000000000000002809009846",
		"00d72e0a0002000001168015f0008002000002000000000000000000000000000000000000000000000000285bee86c4",
		"00d82e0a0002000001168016f0008002000002000000000000000000000000000000000000000000000000281f25725c",
		"00d92e0a0002000001168017f00080020000020000000000000000000000000000000000000000000000002872e94fe1",
		"00da2e0a0002000001168018f000800300000200000000000000000000000000000000000000000000000028e39d572f",
		"00db2e0a0002000001168019f0008003000002000000000000000000000000000000000000000000000000281c9dcadd",
		"00dc2e0a000200000116801af0008003000002000000000000000000000000000000000000000000000000287c5b8405",
		"00dd2e0a000200000116801bf00080030000020000000000000000000000000000000000000000000000002826334420",
		"00de2e0a000200000116801cf00080030000020000000000000000000000000000000000000000000000002871ee1536",
		"00df2e0a000200000116801df0008003000002000000000000000000000000000000000000000000000000289dfeeeb9",
		"00e02e0a000200000116801ef000800300000200000000000000000000000000000000000000000000000028954d55b3",
		"00e12e0a000200000116801ff000800300000200000000000000000000000000000000000000000000000028930c564e",
		"00e22e0a0002000001168020f000800400000200000000000000000000000000000000000000000000000028b9cec3bf",
		"00e32e0a0002000001168021f0008004000002000000000000000000000000000000000000000000000000284263f268",
		"00e42e0a0002000001168022f000800400000200000000000000000000000000000000000000000000000028913e5219",
		"00e52e0a0002000001168023f000800400000200000000000000000000000000000000000000000000000028efe86fe1",
		"00e62e0a0002000001168024f000800400000200000000000000000000000000000000000000000000000028deb045df",
		"00e72e0a0002000001168025f000800400000200000000000000000000000000000000000000000000000028255bcd32",
		"00e82e0a0002000001168026f000800400000200000000000000000000000000000000000000000000000028355392ad",
		"00e92e0a0002000001168027f000800400000200000000000000000000000000000000000000000000000028404a6aca",
		"00ea2e0a0002000001168028f0008005000002000000000000000000000000000000000000000000000000281de78f94",
		"00eb2e0a0002000001168029f000800500000200000000000000000000000000000000000000000000000028501a3aae",
		"00ec2e0a000200000116802af0008005000002000000000000000000000000000000000000000000000000282947d976",
		"00ed2e0a000200000116802bf000800500000200000000000000000000000000000000000000000000000028095cfe0d",
		"00ee2e0a000200000116802cf000800500000200000000000000000000000000000000000000000000000028bbcfc27a",
		"00ef2e0a000200000116802df000800500000200000000000000000000000000000000000000000000000028dbb27396",
		"00f02e0a000200000116802ef000800500000200000000000000000000000000000000000000000000000028dbe9b225",
		"00f12e0a000200000116802ff000800500000200000000000000000000000000000000000000000000000028840c0b08",
		"00f22e0a0002000001168030f0008006000002000000000000000000000000000000000000000000000000287683e4f8",
		"00f32e0a0002000001168031f00080060000020000000000000000000000000000000000000000000000002844d131d1",
		"00f42e0a0002000001168032f0008006000002000000000000000000000000000000000000000000000000284d2c2c6d",
		"00f52e0a0002000001168033f000800600000200000000000000000000000000000000000000000000000028e89a166c",
		"00f62e0a0002000001168034f0008006000002000000000000000000000000000000000000000000000000280f47db8c",
		"00f72e0a0002000001168035f0008006000002000000000000000000000000000000000000000000000000283ede8b3e",
		"00f82e0a0002000001168036f000800600000200000000000000000000000000000000000000000000000028580547db",
		"00f92e0a0002000001168037f000800600000200000000000000000000000000000000000000000000000028d72a270e",
		"00fa2e0a0002000001168038f000800700000200000000000000000000000000000000000000000000000028c25ce712",
		"00fb2e0a0002000001168039f000800700000200000000000000000000000000000000000000000000000028b908637e",
		"00fc2e0a000200000116803af0008007000002000000000000000000000000000000000000000000000000285b66e6fa",
		"00fd2e0a000200000116803bf00080070000020000000000000000000000000000000000000000000000002855c10393",
		"00fe2e0a000200000116803cf0008007000002000000000000000000000000000000000000000000000000283e94c57d",
		"00ff2e0a000200000116803df0008007000002000000000000000000000000000000000000000000000000284347e7f0",
		"01002e0a000200000116803ef000800700000200000000000000000000000000000000000000000000000028be66429d",
		"01012e0a000200000116803ff0008007000002000000000000000000000000000000000000000000000000284f7db145",
		"01022e0a0002000001490401c000000000000000000000000000000000000000000000000000000000000028470aa043",
		"01032e0a00020000014904012000000000000000000000000000000000000000000000000000000000000028a6bc6e48",
		"01042e0a00020000014904011800ffffffff0000000000000000000000000000000000000000000000000028f747c739",
	}
	firstTid := uint16(3)

	for pkt_number, packetString := range mibUploadNextSequence {
		data, err := stringToPacket(packetString)
		assert.NoError(t, err)

		packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
		fmt.Printf("Packet: %v: %v", pkt_number, packet)
		assert.NotNil(t, packet)

		omciLayer := packet.Layer(LayerTypeOMCI)
		assert.NotNil(t, omciLayer)

		omciMsg, ok := omciLayer.(*OMCI)
		assert.True(t, ok)
		assert.Equal(t, omciMsg.TransactionID, firstTid)
		assert.Equal(t, omciMsg.MessageType, byte(me.MibUploadNext))
		assert.Equal(t, omciMsg.Length, uint16(40))

		msgLayer := packet.Layer(LayerTypeMibUploadNextResponse)
		assert.NotNil(t, msgLayer)

		uploadResponse, ok2 := msgLayer.(*MibUploadNextResponse)
		assert.True(t, ok2)
		assert.Equal(t, uploadResponse.EntityClass, me.OnuDataClassId)
		assert.Equal(t, uploadResponse.EntityInstance, uint16(0))

		// Test serialization back to former string
		var options gopacket.SerializeOptions
		options.FixLengths = true

		buffer := gopacket.NewSerializeBuffer()
		err = gopacket.SerializeLayers(buffer, options, omciMsg, uploadResponse)
		assert.NoError(t, err)

		outgoingPacket := buffer.Bytes()
		reconstituted := packetToString(outgoingPacket)
		assert.Equal(t,  strings.ToLower(packetString), reconstituted)

		// Advance TID
		firstTid += 1
	}
}
