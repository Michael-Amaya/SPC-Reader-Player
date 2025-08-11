package models

type SPCProcessor struct {
	PC      uint16
	A       byte
	X       byte
	Y       byte
	PSW     byte
	SP      byte
	Ram64KB [65536]byte
}
