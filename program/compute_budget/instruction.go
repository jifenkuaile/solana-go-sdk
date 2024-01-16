package compute_budget

import (
	"fmt"
	"github.com/jifenkuaile/solana-go-sdk/common"
	"github.com/jifenkuaile/solana-go-sdk/pkg/bincode"
	"github.com/jifenkuaile/solana-go-sdk/types"
	"github.com/near/borsh-go"
)

type Instruction borsh.Enum

const (
	InstructionRequestUnits Instruction = iota
	InstructionRequestHeapFrame
	InstructionSetComputeUnitLimit
	InstructionSetComputeUnitPrice
)

type InstructionStruct struct {
	Instruction Instruction
}

type RequestUnitsParam struct {
	Units         uint32
	AdditionalFee uint32
}

type RequestUnitsStruct struct {
	Instruction   Instruction
	Units         uint32
	AdditionalFee uint32
}

// RequestUnits ...
func RequestUnits(param RequestUnitsParam) types.Instruction {
	data, err := borsh.Serialize(RequestUnitsStruct{
		Instruction:   InstructionRequestUnits,
		Units:         param.Units,
		AdditionalFee: param.AdditionalFee,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.ComputeBudgetProgramID,
		Accounts:  []types.AccountMeta{},
		Data:      data,
	}
}

type RequestHeapFrameParam struct {
	Bytes uint32
}

type RequestHeapFrameStruct struct {
	Instruction Instruction
	Bytes       uint32
}

// RequestHeapFrame ...
func RequestHeapFrame(param RequestHeapFrameParam) types.Instruction {
	data, err := borsh.Serialize(RequestHeapFrameStruct{
		Instruction: InstructionRequestHeapFrame,
		Bytes:       param.Bytes,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.ComputeBudgetProgramID,
		Accounts:  []types.AccountMeta{},
		Data:      data,
	}
}

type SetComputeUnitLimitParam struct {
	Units uint32
}

type SetComputeUnitLimitStruct struct {
	Instruction Instruction
	Units       uint32
}

// SetComputeUnitLimit set a specific compute unit limit that the transaction is allowed to consume.
func SetComputeUnitLimit(param SetComputeUnitLimitParam) types.Instruction {
	data, err := borsh.Serialize(SetComputeUnitLimitStruct{
		Instruction: InstructionSetComputeUnitLimit,
		Units:       param.Units,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.ComputeBudgetProgramID,
		Accounts:  []types.AccountMeta{},
		Data:      data,
	}
}

type SetComputeUnitPriceParam struct {
	MicroLamports uint64
}

type SetComputeUnitPriceStruct struct {
	Instruction   Instruction
	MicroLamports uint64
}

// SetComputeUnitPrice set a compute unit price in "micro-lamports" to pay a higher transaction
// fee for higher transaction prioritization.
func SetComputeUnitPrice(param SetComputeUnitPriceParam) types.Instruction {
	data, err := borsh.Serialize(SetComputeUnitPriceStruct{
		Instruction:   InstructionSetComputeUnitPrice,
		MicroLamports: param.MicroLamports,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.ComputeBudgetProgramID,
		Accounts:  []types.AccountMeta{},
		Data:      data,
	}
}

func GetInstructionType(data []byte) (Instruction, error) {
	instructionType := &InstructionStruct{}

	err := bincode.DeserializeData(data, instructionType)
	if err != nil {
		return 0, fmt.Errorf("unknown instructionType")
	}

	return instructionType.Instruction, nil
}

func DeSerializeInstruction(data []byte, instruction interface{}) error {
	err := bincode.DeserializeData(data, instruction)
	if err != nil {
		return err
	}

	return nil
}
