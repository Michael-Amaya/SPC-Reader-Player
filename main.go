package main

import (
	"encoding/binary"
	"log"
	"mike-a/spc-reader/models"
	"os"

	"github.com/k0kubun/pp"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <FILENAME>", os.Args[0])
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	encodedSpc, err := models.ReadSPC(file)
	if err != nil {
		log.Fatal(err)
	}
	// pp.Printf("%+v\n", encodedSpc)
	// printEncodedSPCHeader(encodedSpc)

	// Emulate SPC700 CPU
	// Populate CPU with registers and memory
	spcProcessor := models.SPCProcessor{
		PC:      binary.LittleEndian.Uint16(encodedSpc.PC[:]),
		A:       encodedSpc.A,
		X:       encodedSpc.X,
		Y:       encodedSpc.Y,
		PSW:     encodedSpc.PSW,
		SP:      encodedSpc.SP,
		Ram64KB: encodedSpc.Ram64KB,
	}

	pp.Printf("%+v\n", spcProcessor)
	// fmt.Printf("% 0x \n", spcProcessor.Ram64KB[:])

	// instructionSet := models.BuildInstructionArray()
	// pp.Printf("%+v\n", instructionSet)
}

func printEncodedSPCHeader(encodedSpc *models.SPCFile) {
	if encodedSpc == nil {
		return
	}
	pp.Printf("Header:\t\t%+v\n", string(encodedSpc.Header[:]))
	pp.Printf("Song Title:\t%+v\n", string(encodedSpc.TextSongTitle[:]))
	pp.Printf("Game Title:\t%+v\n", string(encodedSpc.TextGameTitle[:]))
	pp.Printf("Dumper:\t\t%+v\n", string(encodedSpc.TextDumper[:]))
	pp.Printf("Comments:\t%+v\n", string(encodedSpc.TextComments[:]))
	pp.Printf("Date Dumped:\t%+v\n", string(encodedSpc.TextDateDumped[:]))
	pp.Printf("Fading:\t\t%+v\n", string(encodedSpc.TextNumSecondsBeforeFading[:]))
	pp.Printf("LenFade:\t%+v\n", string(encodedSpc.TextLengthFadeMs[:]))
	pp.Printf("Artist:\t\t%+v\n", string(encodedSpc.TextArtist[:]))
	pp.Printf("Emulator:\t%+v\n", int(encodedSpc.TextEmulatorUsed))
}
