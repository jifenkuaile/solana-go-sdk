package bincode

import (
	"encoding/base64"
	"testing"
)

type Instruction uint8

const (
	InstructionInitializeMint Instruction = iota
	InstructionInitializeAccount
	InstructionInitializeMultisig
	InstructionTransfer
	InstructionApprove
	InstructionRevoke
	InstructionSetAuthority
	InstructionMintTo
	InstructionBurn
	InstructionCloseAccount
	InstructionFreezeAccount
	InstructionThawAccount
	InstructionTransferChecked
	InstructionApproveChecked
	InstructionMintToChecked
	InstructionBurnChecked
	InstructionInitializeAccount2
	InstructionSyncNative
	InstructionInitializeAccount3
	InstructionInitializeMultisig2
	InstructionInitializeMint2
)

func TestSerialize(t *testing.T) {
	type testStruct struct {
		Instruction Instruction
		Amount      uint64
		Decimals    uint8
	}
	type testType struct {
		Instruction Instruction
	}

	s := testStruct{
		Instruction: InstructionTransferChecked,
		Amount:      100,
		Decimals:    6,
	}

	data, err := SerializeData(s)
	if err != nil {
		t.Fatalf("error when serialized data, err: %v", err.Error())
	}

	t.Logf("result: %v", base64.StdEncoding.EncodeToString(data))

	info := &testStruct{}
	err = DeserializeData(data, info)
	if err != nil {
		t.Fatalf("error when deserialized data, err: %v", err.Error())
	}

	t.Logf("info: %#v", info)

	instructionType := &testType{}
	err = DeserializeData(data, instructionType)
	if err != nil {
		t.Fatalf("error when get instructionDataType, err: %v", err.Error())
	}

	t.Logf("instruction type: %v", instructionType)
}
