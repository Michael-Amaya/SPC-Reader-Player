package models

import (
	"encoding/binary"
	"errors"
	"log"
	"mike-a/spc-reader/models/enums"
	"mike-a/spc-reader/utils"
)

type ExtendedID666Value struct {
	DataType enums.DataType
	Data     []byte
}

type SPCFile struct {
	Header                       [33]byte    // 0x00000
	TwentySixTwentySix           [2]byte     // 0x00021
	ID666Info                    byte        // 0x00023
	VersionMinor                 byte        // 0x00024
	PC                           [2]byte     // 0x00025
	A                            byte        // 0x00027
	X                            byte        // 0x00028
	Y                            byte        // 0x00029
	PSW                          byte        // 0x0002A
	SP                           byte        // 0x0002B
	Reserved                     [2]byte     // 0x0002C
	TextSongTitle                [32]byte    // 0x0002E
	TextGameTitle                [32]byte    // 0x0004E
	TextDumper                   [16]byte    // 0x0006E
	TextComments                 [32]byte    // 0x0007E
	TextDateDumped               [11]byte    // 0x0009E
	TextNumSecondsBeforeFading   [3]byte     // 0x000A9
	TextLengthFadeMs             [5]byte     // 0x000AC
	TextArtist                   [32]byte    // 0x000B1
	TextDefaultChannelDisable    byte        // 0x000D1
	TextEmulatorUsed             byte        // 0x000D2
	TextReserved                 [45]byte    // 0x000D3
	BinarySongTitle              [32]byte    // 0x0002E
	BinaryGameTitle              [32]byte    // 0x0004E
	BinaryDumper                 [16]byte    // 0x0006E
	BinaryComments               [32]byte    // 0x0007E
	BinaryDateDumped             [4]byte     // 0x0009E
	BinaryUnused                 [7]byte     // 0x000A2
	BinaryNumSecondsBeforeFading [3]byte     // 0x000A9
	BinaryLengthFadeMs           [3]byte     // 0x000AC
	BinaryArtist                 [32]byte    // 0x000B0
	BinaryDefaultChannelDisable  byte        // 0x000D0
	BinaryEmulatorUsed           byte        // 0x000D1
	BinaryReserved               [46]byte    // 0x000D2
	Ram64KB                      [65536]byte // 0x00100
	DSPRegisters                 [128]byte   // 0x10100
	Unused2                      [64]byte    // 0x10180
	ExtraRam                     [64]byte    // 0x101C0

	// Internal
	ContainsID666Info bool
	UsingText         bool
	ExtendedID666Info map[enums.ExtendedID666Type]ExtendedID666Value
}

// ReadSPC reads a SPC file, returns a pointer to the SPCFile struct
func ReadSPC(data []byte) (*SPCFile, error) {
	// This func uses a lot of magic numbers for offsets, but it is gotten from:
	// https://wiki.superfamicom.org/spc-and-rsn-file-format
	MINIMUM_SIZE := 46 // Sanity check to make sure header exists
	if len(data) < MINIMUM_SIZE {
		return nil, errors.New("invalid file")
	}

	encodedSPCFile := SPCFile{}

	// Copy Header and check if it's a valid SPC
	copy(encodedSPCFile.Header[:], data[0x0:0x0+33])
	copy(encodedSPCFile.TwentySixTwentySix[:], data[0x21:0x21+2])
	if string(encodedSPCFile.Header[0:8]) != "SNES-SPC" || (encodedSPCFile.TwentySixTwentySix[0] != utils.VERIFICATION && encodedSPCFile.TwentySixTwentySix[1] != utils.VERIFICATION) {
		return nil, errors.New("invalid file, not SPC")
	}

	encodedSPCFile.ID666Info = data[0x23]
	encodedSPCFile.VersionMinor = data[0x24]
	copy(encodedSPCFile.PC[:], data[0x25:0x25+2])
	encodedSPCFile.A = data[0x27]
	encodedSPCFile.X = data[0x28]
	encodedSPCFile.Y = data[0x29]
	encodedSPCFile.PSW = data[0x2A]
	encodedSPCFile.SP = data[0x2B]
	copy(encodedSPCFile.Reserved[:], data[0x2C:0x2C+2])

	if encodedSPCFile.ID666Info == utils.VERIFICATION { // 0x1A is the verification for ID666
		encodedSPCFile.ContainsID666Info = true
		// Text based data
		copy(encodedSPCFile.TextSongTitle[:], data[0x2E:0x2E+32])
		copy(encodedSPCFile.TextGameTitle[:], data[0x4E:0x4E+32])
		copy(encodedSPCFile.TextDumper[:], data[0x6E:0x6E+16])
		copy(encodedSPCFile.TextComments[:], data[0x7E:0x7E+32])
		copy(encodedSPCFile.TextDateDumped[:], data[0x9E:0x9E+11])
		copy(encodedSPCFile.TextNumSecondsBeforeFading[:], data[0xA9:0xA9+3])
		copy(encodedSPCFile.TextLengthFadeMs[:], data[0xAC:0xAC+5])
		copy(encodedSPCFile.TextArtist[:], data[0xB1:0xB1+32])
		encodedSPCFile.TextDefaultChannelDisable = data[0xD1]
		encodedSPCFile.TextEmulatorUsed = data[0xD2]
		copy(encodedSPCFile.TextReserved[:], data[0xD3:0xD3+45])
	}

	// Check if contains I666 and if it is text based
	if encodedSPCFile.ContainsID666Info && checkID666InfoIsText(&encodedSPCFile) {
		encodedSPCFile.UsingText = true
	} else if encodedSPCFile.ContainsID666Info && !checkID666InfoIsText(&encodedSPCFile) {
		// Binary Based Data
		copy(encodedSPCFile.BinarySongTitle[:], data[0x2E:0x2E+32])
		copy(encodedSPCFile.BinaryGameTitle[:], data[0x4E:0x4E+32])
		copy(encodedSPCFile.BinaryDumper[:], data[0x6E:0x6E+16])
		copy(encodedSPCFile.BinaryComments[:], data[0x7E:0x7E+32])
		copy(encodedSPCFile.BinaryDateDumped[:], data[0x9E:0x9E+4])
		copy(encodedSPCFile.BinaryUnused[:], data[0xA2:0xA2+7])
		copy(encodedSPCFile.BinaryNumSecondsBeforeFading[:], data[0xA9:0xA9+3])
		copy(encodedSPCFile.BinaryLengthFadeMs[:], data[0xAC:0xAC+3])
		copy(encodedSPCFile.BinaryArtist[:], data[0xB0:0xB0+32])
		encodedSPCFile.BinaryDefaultChannelDisable = data[0xD0]
		encodedSPCFile.BinaryEmulatorUsed = data[0xD1]
		copy(encodedSPCFile.BinaryReserved[:], data[0xD2:0xD2+46])
	}

	// Populate 64Kb ram, DSP register, and extra RAM
	copy(encodedSPCFile.Ram64KB[:], data[0x100:0x100+65536])
	copy(encodedSPCFile.DSPRegisters[:], data[0x10100:0x10100+128])
	copy(encodedSPCFile.Unused2[:], data[0x10180:0x10180+64])
	copy(encodedSPCFile.ExtraRam[:], data[0x101C0:0x101C0+64])
	encodedSPCFile.ExtendedID666Info = map[enums.ExtendedID666Type]ExtendedID666Value{}

	// Process Extended ID666, starts at 0x10200, header is 4 bytes
	if len(data) > 0x10200+4 && string(data[0x10200:0x10200+4]) == "xid6" {
		// Contains Extended ID666
		chunkSize := binary.LittleEndian.Uint32(data[0x10200+4 : 0x10200+8])
		log.Printf("Chunk Size: %d\n", chunkSize)
		remainingSize := chunkSize - 8 // Already processed 8 bytes of header
		currentOffset := 0x10208       // Start of data
		for remainingSize > 0 {
			idByte := data[currentOffset]
			typeByte := data[currentOffset+1]
			typeType := enums.DataType(typeByte)
			dataBytes := data[currentOffset+2 : currentOffset+4] // 2 bytes
			idType := enums.ExtendedID666Type(idByte)
			remainingSize -= 4
			currentOffset += 4
			if typeByte == 0x0 {
				// Data Bytes contains the data
				encodedSPCFile.ExtendedID666Info[idType] = ExtendedID666Value{DataType: typeType, Data: dataBytes}
			} else {

				// dataBytes contains the length of the data, which follows this header, so starts at 0x10208 + 4
				dataLength := binary.LittleEndian.Uint16(dataBytes)
				startLocation := currentOffset
				endLocation := currentOffset + int(dataLength)

				if endLocation > len(data) {
					log.Printf("Went over? %d > %d\n", endLocation, len(data))
					break
				}

				encodedSPCFile.ExtendedID666Info[idType] = ExtendedID666Value{DataType: typeType, Data: data[startLocation:endLocation]} //data[startLocation:endLocation]
				currentOffset += int(dataLength)
				remainingSize -= uint32(dataLength)
			}
		}
	}

	return &encodedSPCFile, nil
}

func checkID666InfoIsText(spc *SPCFile) bool {
	toCheck := [][]byte{
		spc.TextSongTitle[:],
		spc.TextGameTitle[:],
		spc.TextDumper[:],
		spc.TextComments[:],
		spc.TextDateDumped[:],
		spc.TextNumSecondsBeforeFading[:],
		spc.TextLengthFadeMs[:],
		spc.TextArtist[:],
	}
	toPass := len(toCheck) / 2
	numFailed := 0
	for _, v := range toCheck {
		for _, c := range v {
			if c != 0x0 && (c < utils.MINIMUM_PRINTABLE_CHAR || c > utils.MAXIMUM_PRINTABLE_CHAR) {
				numFailed++
				break // Breaks out of inner loop
			}
		}
	}

	return numFailed < toPass
}
