package models

type Instruction struct {
	Mnemonic             string // MOV, JMP, ADC
	AddressingMode       AddressingMode
	Length               int
	Cycles               int
	HasConditionalCycles bool
}

type AddressingMode int

const (
	AddressingModeDirectPage AddressingMode = iota
	AddressingModeXIndexedDirectPage
	AddressingModeYIndexedDirectPage
	AddressingModeIndirect
	AddressingModeIndirectAutoIncrement
	AddressingModeDirectPageToDirectPage
	AddressingModeIndirectPageToIndirectPage
	AddressingModeImmediateDataToDirectPage
	AddressingModeDirectPageBit
	AddressingModeDirectPageBitRelative
	AddressingModeAbsoluteBooleanBit
	AddressingModeAbsolute
	AddressingModeAbsoluteXIndexedIndirect
	AddressingModeXIndexedAbsolute
	AddressingModeYIndexedAbsolute
	AddressingModeXIndexedIndirect
	AddressingModeIndirectYIndexed
	AddressingModeRelative
	AddressingModeImmediate
	AddressingModeAccumulator
	AddressingModeImplied
)

func BuildInstructionArray() [256]Instruction {
	arr := [256]Instruction{}
	// First column of matrix
	arr[0x00] = Instruction{Mnemonic: "NOP", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0x01] = Instruction{Mnemonic: "BPL", AddressingMode: AddressingModeRelative, Length: 2, Cycles: 2, HasConditionalCycles: true}
	arr[0x02] = Instruction{Mnemonic: "CLRP", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0x03] = Instruction{Mnemonic: "BMI", AddressingMode: AddressingModeRelative, Length: 2, Cycles: 2, HasConditionalCycles: true}
	arr[0x04] = Instruction{Mnemonic: "SETP", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0x05] = Instruction{Mnemonic: "BVC", AddressingMode: AddressingModeRelative, Length: 2, Cycles: 2, HasConditionalCycles: true}
	arr[0x06] = Instruction{Mnemonic: "CLRC", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0x07] = Instruction{Mnemonic: "BVS", AddressingMode: AddressingModeRelative, Length: 2, Cycles: 2, HasConditionalCycles: true}
	arr[0x08] = Instruction{Mnemonic: "SETC", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0x09] = Instruction{Mnemonic: "BCC", AddressingMode: AddressingModeRelative, Length: 2, Cycles: 2, HasConditionalCycles: true}
	arr[0x0A] = Instruction{Mnemonic: "EI", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0x0B] = Instruction{Mnemonic: "BCS", AddressingMode: AddressingModeRelative, Length: 2, Cycles: 2, HasConditionalCycles: true}
	arr[0x0C] = Instruction{Mnemonic: "DI", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0x0D] = Instruction{Mnemonic: "BNE", AddressingMode: AddressingModeRelative, Length: 2, Cycles: 2, HasConditionalCycles: true}
	arr[0x0E] = Instruction{Mnemonic: "CLRV", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0x0F] = Instruction{Mnemonic: "BEQ", AddressingMode: AddressingModeRelative, Length: 2, Cycles: 2, HasConditionalCycles: true}

	// Second column of matrix, all TCALL
	arr[0x10] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x11] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x12] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x13] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x14] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x15] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x16] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x17] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x18] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x19] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x1A] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x1B] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x1C] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x1D] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x1E] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0x1F] = Instruction{Mnemonic: "TCALL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}

	// Third column of matrix,
	arr[0x20] = Instruction{Mnemonic: "SET1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x21] = Instruction{Mnemonic: "CLR1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x22] = Instruction{Mnemonic: "SET1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x23] = Instruction{Mnemonic: "CLR1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x24] = Instruction{Mnemonic: "SET1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x25] = Instruction{Mnemonic: "CLR1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x26] = Instruction{Mnemonic: "SET1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x27] = Instruction{Mnemonic: "CLR1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x28] = Instruction{Mnemonic: "SET1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x29] = Instruction{Mnemonic: "CLR1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x2A] = Instruction{Mnemonic: "SET1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x2B] = Instruction{Mnemonic: "CLR1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x2C] = Instruction{Mnemonic: "SET1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x2D] = Instruction{Mnemonic: "CLR1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x2E] = Instruction{Mnemonic: "SET1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x2F] = Instruction{Mnemonic: "CLR1", AddressingMode: AddressingModeDirectPageBit, Length: 2, Cycles: 4, HasConditionalCycles: false}

	// Fourth Column
	arr[0x30] = Instruction{Mnemonic: "BBS", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x31] = Instruction{Mnemonic: "BBC", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x32] = Instruction{Mnemonic: "BBS", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x33] = Instruction{Mnemonic: "BBC", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x34] = Instruction{Mnemonic: "BBS", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x35] = Instruction{Mnemonic: "BBC", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x36] = Instruction{Mnemonic: "BBS", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x37] = Instruction{Mnemonic: "BBC", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x38] = Instruction{Mnemonic: "BBS", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x39] = Instruction{Mnemonic: "BBC", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x3A] = Instruction{Mnemonic: "BBS", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x3B] = Instruction{Mnemonic: "BBC", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x3C] = Instruction{Mnemonic: "BBS", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x3D] = Instruction{Mnemonic: "BBC", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x3E] = Instruction{Mnemonic: "BBS", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}
	arr[0x3F] = Instruction{Mnemonic: "BBC", AddressingMode: AddressingModeDirectPageBitRelative, Length: 2, Cycles: 5, HasConditionalCycles: true}

	// Fifth Column
	arr[0x40] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 3, HasConditionalCycles: false}
	arr[0x41] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x42] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 3, HasConditionalCycles: false}
	arr[0x43] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x44] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 3, HasConditionalCycles: false}
	arr[0x45] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x46] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0x47] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x48] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 3, HasConditionalCycles: false}
	arr[0x49] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x4A] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 3, HasConditionalCycles: false}
	arr[0x4B] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x4C] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x4D] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0x4E] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 3, HasConditionalCycles: false}
	arr[0x4F] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}

	// Sixth Column
	arr[0x50] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0x51] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeXIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x52] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0x53] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeXIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x54] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0x55] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeXIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x56] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0x57] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeXIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x58] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0x59] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeXIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x5A] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0x5B] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeXIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x5C] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x5D] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeXIndexedAbsolute, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0x5E] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0x5F] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeXIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}

	// Seventh Column
	arr[0x60] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeIndirect, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0x61] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeYIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x62] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeIndirect, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0x63] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeYIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x64] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeIndirect, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0x65] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeYIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x66] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeIndirect, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0x67] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeYIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x68] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeIndirect, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0x69] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeYIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x6A] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeIndirect, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0x6B] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeYIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x6C] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeIndirect, Length: 1, Cycles: 4, HasConditionalCycles: false}
	arr[0x6D] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeYIndexedAbsolute, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0x6E] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeIndirect, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0x6F] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeYIndexedAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}

	// Eighth Column
	arr[0x70] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeXIndexedIndirect, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x71] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeIndirectYIndexed, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x72] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeXIndexedIndirect, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x73] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeIndirectYIndexed, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x74] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeXIndexedIndirect, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x75] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeIndirectYIndexed, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x76] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeXIndexedIndirect, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x77] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeIndirectYIndexed, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x78] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeXIndexedIndirect, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x79] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeIndirectYIndexed, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x7A] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeXIndexedIndirect, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x7B] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeIndirectYIndexed, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x7C] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeXIndexedIndirect, Length: 2, Cycles: 7, HasConditionalCycles: false}
	arr[0x7D] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeIndirectYIndexed, Length: 2, Cycles: 7, HasConditionalCycles: false}
	arr[0x7E] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeXIndexedIndirect, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0x7F] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeIndirectYIndexed, Length: 2, Cycles: 6, HasConditionalCycles: false}

	// Ninth Column
	arr[0x80] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeImmediate, Length: 2, Cycles: 2, HasConditionalCycles: false}
	arr[0x81] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeImmediateDataToDirectPage, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x82] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeImmediate, Length: 2, Cycles: 2, HasConditionalCycles: false}
	arr[0x83] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeImmediateDataToDirectPage, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x84] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeImmediate, Length: 2, Cycles: 2, HasConditionalCycles: false}
	arr[0x85] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeImmediateDataToDirectPage, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x86] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeImmediate, Length: 2, Cycles: 2, HasConditionalCycles: false}
	arr[0x87] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeImmediateDataToDirectPage, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x88] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeImmediate, Length: 2, Cycles: 2, HasConditionalCycles: false}
	arr[0x89] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeImmediateDataToDirectPage, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x8A] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeImmediate, Length: 2, Cycles: 2, HasConditionalCycles: false}
	arr[0x8B] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeImmediateDataToDirectPage, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x8C] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeImmediate, Length: 2, Cycles: 2, HasConditionalCycles: false}
	arr[0x8D] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0x8E] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeImmediate, Length: 2, Cycles: 2, HasConditionalCycles: false}
	arr[0x8F] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 3, HasConditionalCycles: false}

	// Tenth Column
	arr[0x90] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeDirectPageToDirectPage, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0x91] = Instruction{Mnemonic: "OR", AddressingMode: AddressingModeIndirectPageToIndirectPage, Length: 1, Cycles: 5, HasConditionalCycles: false}
	arr[0x92] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeDirectPageToDirectPage, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0x93] = Instruction{Mnemonic: "AND", AddressingMode: AddressingModeIndirectPageToIndirectPage, Length: 1, Cycles: 5, HasConditionalCycles: false}
	arr[0x94] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeDirectPageToDirectPage, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0x95] = Instruction{Mnemonic: "EOR", AddressingMode: AddressingModeIndirectPageToIndirectPage, Length: 1, Cycles: 5, HasConditionalCycles: false}
	arr[0x96] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeDirectPageToDirectPage, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0x97] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeIndirectPageToIndirectPage, Length: 1, Cycles: 5, HasConditionalCycles: false}
	arr[0x98] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeDirectPageToDirectPage, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0x99] = Instruction{Mnemonic: "ADC", AddressingMode: AddressingModeIndirectPageToIndirectPage, Length: 1, Cycles: 5, HasConditionalCycles: false}
	arr[0x9A] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeDirectPageToDirectPage, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0x9B] = Instruction{Mnemonic: "SBC", AddressingMode: AddressingModeIndirectPageToIndirectPage, Length: 1, Cycles: 5, HasConditionalCycles: false}
	arr[0x9C] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0x9D] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeYIndexedDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0x9E] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0x9F] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeYIndexedDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}

	// Eleventh Column
	arr[0xA0] = Instruction{Mnemonic: "OR1", AddressingMode: AddressingModeAbsoluteBooleanBit, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xA1] = Instruction{Mnemonic: "DECW", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0xA2] = Instruction{Mnemonic: "OR1", AddressingMode: AddressingModeAbsoluteBooleanBit, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xA3] = Instruction{Mnemonic: "INCW", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0xA4] = Instruction{Mnemonic: "AND1", AddressingMode: AddressingModeAbsoluteBooleanBit, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0xA5] = Instruction{Mnemonic: "CMPW", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0xA6] = Instruction{Mnemonic: "AND1", AddressingMode: AddressingModeAbsoluteBooleanBit, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0xA7] = Instruction{Mnemonic: "ADDW", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0xA8] = Instruction{Mnemonic: "EOR1", AddressingMode: AddressingModeAbsoluteBooleanBit, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xA9] = Instruction{Mnemonic: "SUBW", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0xAA] = Instruction{Mnemonic: "MOV1", AddressingMode: AddressingModeAbsoluteBooleanBit, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0xAB] = Instruction{Mnemonic: "MOVW", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0xAC] = Instruction{Mnemonic: "MOV1", AddressingMode: AddressingModeAbsoluteBooleanBit, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0xAD] = Instruction{Mnemonic: "MOVW", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0xAE] = Instruction{Mnemonic: "NOT1", AddressingMode: AddressingModeAbsoluteBooleanBit, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xAF] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeDirectPageToDirectPage, Length: 3, Cycles: 5, HasConditionalCycles: false}

	// Twelfth Column
	arr[0xB0] = Instruction{Mnemonic: "ASL", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0xB1] = Instruction{Mnemonic: "ASL", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0xB2] = Instruction{Mnemonic: "ROL", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0xB3] = Instruction{Mnemonic: "ROL", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0xB4] = Instruction{Mnemonic: "LSR", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0xB5] = Instruction{Mnemonic: "LSR", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0xB6] = Instruction{Mnemonic: "ROR", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0xB7] = Instruction{Mnemonic: "ROR", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0xB8] = Instruction{Mnemonic: "DEC", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0xB9] = Instruction{Mnemonic: "DEC", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0xBA] = Instruction{Mnemonic: "INC", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0xBB] = Instruction{Mnemonic: "INC", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0xBC] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}
	arr[0xBD] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 5, HasConditionalCycles: false}
	arr[0xBE] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 3, HasConditionalCycles: false}
	arr[0xBF] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeXIndexedDirectPage, Length: 2, Cycles: 4, HasConditionalCycles: false}

	// Thirteenth Column
	arr[0xC0] = Instruction{Mnemonic: "ASL", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xC1] = Instruction{Mnemonic: "ASL", AddressingMode: AddressingModeAccumulator, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xC2] = Instruction{Mnemonic: "ROL", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xC3] = Instruction{Mnemonic: "ROL", AddressingMode: AddressingModeAccumulator, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xC4] = Instruction{Mnemonic: "LSR", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xC5] = Instruction{Mnemonic: "LSR", AddressingMode: AddressingModeAccumulator, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xC6] = Instruction{Mnemonic: "ROR", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xC7] = Instruction{Mnemonic: "ROR", AddressingMode: AddressingModeAccumulator, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xC8] = Instruction{Mnemonic: "DEC", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xC9] = Instruction{Mnemonic: "DEC", AddressingMode: AddressingModeAccumulator, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xCA] = Instruction{Mnemonic: "INC", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xCB] = Instruction{Mnemonic: "INC", AddressingMode: AddressingModeAccumulator, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xCC] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xCD] = Instruction{Mnemonic: "DEC", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xCE] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0xCF] = Instruction{Mnemonic: "INC", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}

	// Fourteenth Column
	arr[0xD0] = Instruction{Mnemonic: "PUSH", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 4, HasConditionalCycles: false}
	arr[0xD1] = Instruction{Mnemonic: "DEC", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xD2] = Instruction{Mnemonic: "PUSH", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 4, HasConditionalCycles: false}
	arr[0xD3] = Instruction{Mnemonic: "INC", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xD4] = Instruction{Mnemonic: "PUSH", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 4, HasConditionalCycles: false}
	arr[0xD5] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xD6] = Instruction{Mnemonic: "PUSH", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 4, HasConditionalCycles: false}
	arr[0xD7] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xD8] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeImmediate, Length: 2, Cycles: 2, HasConditionalCycles: false}
	arr[0xD9] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xDA] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeImmediate, Length: 2, Cycles: 2, HasConditionalCycles: false}
	arr[0xDB] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xDC] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeImmediate, Length: 2, Cycles: 2, HasConditionalCycles: false}
	arr[0xDD] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}
	arr[0xDE] = Instruction{Mnemonic: "NOTC", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0xDF] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 2, HasConditionalCycles: false}

	// Fifteenth Column
	arr[0xE0] = Instruction{Mnemonic: "TSET1", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0xE1] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0xE2] = Instruction{Mnemonic: "CBNE", AddressingMode: AddressingModeRelative, Length: 3, Cycles: 5, HasConditionalCycles: true}
	arr[0xE3] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 3, HasConditionalCycles: false}
	arr[0xE4] = Instruction{Mnemonic: "TCLR1", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0xE5] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 4, HasConditionalCycles: false}
	arr[0xE6] = Instruction{Mnemonic: "DBNZ", AddressingMode: AddressingModeRelative, Length: 3, Cycles: 5, HasConditionalCycles: true}
	arr[0xE7] = Instruction{Mnemonic: "CMP", AddressingMode: AddressingModeDirectPage, Length: 2, Cycles: 3, HasConditionalCycles: false}
	arr[0xE8] = Instruction{Mnemonic: "POP", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 4, HasConditionalCycles: false}
	arr[0xE9] = Instruction{Mnemonic: "DIV", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 12, HasConditionalCycles: false}
	arr[0xEA] = Instruction{Mnemonic: "POP", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 4, HasConditionalCycles: false}
	arr[0xEB] = Instruction{Mnemonic: "DAS", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0xEC] = Instruction{Mnemonic: "POP", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 4, HasConditionalCycles: false}
	arr[0xED] = Instruction{Mnemonic: "CBNE", AddressingMode: AddressingModeXIndexedDirectPage, Length: 3, Cycles: 6, HasConditionalCycles: true}
	arr[0xEE] = Instruction{Mnemonic: "POP", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 4, HasConditionalCycles: false}
	arr[0xEF] = Instruction{Mnemonic: "DBNZ", AddressingMode: AddressingModeRelative, Length: 3, Cycles: 5, HasConditionalCycles: true}

	// Sixteenth Column
	arr[0xF0] = Instruction{Mnemonic: "BRK", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 8, HasConditionalCycles: false}
	arr[0xF1] = Instruction{Mnemonic: "JMP", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 6, HasConditionalCycles: false}
	arr[0xF2] = Instruction{Mnemonic: "BRA", AddressingMode: AddressingModeRelative, Length: 2, Cycles: 4, HasConditionalCycles: true}
	arr[0xF3] = Instruction{Mnemonic: "CALL", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 8, HasConditionalCycles: false}
	arr[0xF4] = Instruction{Mnemonic: "PCALL", AddressingMode: AddressingModeImplied, Length: 2, Cycles: 6, HasConditionalCycles: false}
	arr[0xF5] = Instruction{Mnemonic: "JMP", AddressingMode: AddressingModeAbsolute, Length: 3, Cycles: 3, HasConditionalCycles: false}
	arr[0xF6] = Instruction{Mnemonic: "RET", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 5, HasConditionalCycles: false}
	arr[0xF7] = Instruction{Mnemonic: "RETI", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 6, HasConditionalCycles: false}
	arr[0xF8] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeImmediateDataToDirectPage, Length: 3, Cycles: 5, HasConditionalCycles: false}
	arr[0xF9] = Instruction{Mnemonic: "XCN", AddressingMode: AddressingModeAccumulator, Length: 1, Cycles: 5, HasConditionalCycles: false}
	arr[0xFA] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeIndirectAutoIncrement, Length: 1, Cycles: 4, HasConditionalCycles: false}
	arr[0xFB] = Instruction{Mnemonic: "MOV", AddressingMode: AddressingModeIndirectAutoIncrement, Length: 1, Cycles: 4, HasConditionalCycles: false}
	arr[0xFC] = Instruction{Mnemonic: "MUL", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 9, HasConditionalCycles: false}
	arr[0xFD] = Instruction{Mnemonic: "DAA", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0xFE] = Instruction{Mnemonic: "SLEEP", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 3, HasConditionalCycles: false}
	arr[0xFF] = Instruction{Mnemonic: "STOP", AddressingMode: AddressingModeImplied, Length: 1, Cycles: 3, HasConditionalCycles: false}

	return arr
}
