/*
Output of below code: 

Decoded struct: {1234, "keepdecoding", 56, "dontstop", 789, "congratulations", , 123456789}

*/

package main

import (
    "encoding/binary"
    "fmt"
)

type DecodedStruct struct {
    Short1     uint16
    String1    string
    Byte1      uint8
    String2    string
    Short2     uint16
    String3    string
    Long1      uint32
}

func (d DecodedStruct) String() string {
    return fmt.Sprintf("{%d, \"%s\", %d, \"%s\", %d, \"%s\", , %d}",
        d.Short1, d.String1, d.Byte1, d.String2, d.Short2, d.String3, d.Long1)
}

func decodePacket(packet []byte) (DecodedStruct, error) {
    if len(packet) != 44 {
        return DecodedStruct{}, fmt.Errorf("Packet size must be 44 bytes")
    }

    decoded := DecodedStruct{}

    decoded.Short1 = binary.BigEndian.Uint16(packet[0:2])
    decoded.String1 = string(packet[2:14])
    decoded.Byte1 = packet[14]
    decoded.String2 = string(packet[15:23])
    decoded.Short2 = binary.BigEndian.Uint16(packet[23:25])
    decoded.String3 = string(packet[25:40])
    decoded.Long1 = binary.BigEndian.Uint32(packet[40:44])

    return decoded, nil
}

func main() {
    packet := []byte{
        0x04, 0xD2, 0x6B, 0x65, 0x65, 0x70, 0x64, 0x65,
        0x63, 0x6F, 0x64, 0x69, 0x6E, 0x67, 0x38, 0x64,
        0x6F, 0x6E, 0x74, 0x73, 0x74, 0x6F, 0x70, 0x03,
        0x15, 0x63, 0x6F, 0x6E, 0x67, 0x72, 0x61, 0x74,
        0x75, 0x6C, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x73,
        0x07, 0x5B, 0xCD, 0x15,
    }

    decoded, err := decodePacket(packet)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Decoded struct:", decoded)
}