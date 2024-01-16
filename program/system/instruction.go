package system

import (
	"fmt"
	"github.com/jifenkuaile/solana-go-sdk/common"
	"github.com/jifenkuaile/solana-go-sdk/pkg/bincode"
	"github.com/jifenkuaile/solana-go-sdk/types"
)

type Instruction uint32

const (
	InstructionCreateAccount Instruction = iota
	InstructionAssign
	InstructionTransfer
	InstructionCreateAccountWithSeed
	InstructionAdvanceNonceAccount
	InstructionWithdrawNonceAccount
	InstructionInitializeNonceAccount
	InstructionAuthorizeNonceAccount
	InstructionAllocate
	InstructionAllocateWithSeed
	InstructionAssignWithSeed
	InstructionTransferWithSeed
	InstructionUpgradeNonceAccount
)

type InstructionStruct struct {
	Instruction Instruction
}

type CreateAccountParam struct {
	From     common.PublicKey
	New      common.PublicKey
	Owner    common.PublicKey
	Lamports uint64
	Space    uint64
}

type CreateAccountStruct struct {
	Instruction Instruction
	Lamports    uint64
	Space       uint64
	Owner       common.PublicKey
}

func CreateAccount(param CreateAccountParam) types.Instruction {
	data, err := bincode.SerializeData(CreateAccountStruct{
		Instruction: InstructionCreateAccount,
		Lamports:    param.Lamports,
		Space:       param.Space,
		Owner:       param.Owner,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.From, IsSigner: true, IsWritable: true},
			{PubKey: param.New, IsSigner: true, IsWritable: true},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type AssignParam struct {
	From  common.PublicKey
	Owner common.PublicKey
}

type AssignStruct struct {
	Instruction       Instruction
	AssignToProgramID common.PublicKey
}

func Assign(param AssignParam) types.Instruction {
	data, err := bincode.SerializeData(AssignStruct{
		Instruction:       InstructionAssign,
		AssignToProgramID: param.Owner,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.From, IsSigner: true, IsWritable: true},
		},
		Data: data,
	}
}

type TransferParam struct {
	From   common.PublicKey
	To     common.PublicKey
	Amount uint64
}

type TransferStruct struct {
	Instruction Instruction
	Lamports    uint64
}

func Transfer(param TransferParam) types.Instruction {
	data, err := bincode.SerializeData(TransferStruct{
		Instruction: InstructionTransfer,
		Lamports:    param.Amount,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.From, IsSigner: true, IsWritable: true},
			{PubKey: param.To, IsSigner: false, IsWritable: true},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type CreateAccountWithSeedParam struct {
	From     common.PublicKey
	New      common.PublicKey
	Base     common.PublicKey
	Owner    common.PublicKey
	Seed     string
	Lamports uint64
	Space    uint64
}

type CreateAccountWithSeedStruct struct {
	Instruction Instruction
	Base        common.PublicKey
	Seed        string
	Lamports    uint64
	Space       uint64
	ProgramID   common.PublicKey
}

func CreateAccountWithSeed(param CreateAccountWithSeedParam) types.Instruction {
	data, err := bincode.SerializeData(CreateAccountWithSeedStruct{
		Instruction: InstructionCreateAccountWithSeed,
		Base:        param.Base,
		Seed:        param.Seed,
		Lamports:    param.Lamports,
		Space:       param.Space,
		ProgramID:   param.Owner,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3)
	accounts = append(accounts,
		types.AccountMeta{PubKey: param.From, IsSigner: true, IsWritable: true},
		types.AccountMeta{PubKey: param.New, IsSigner: false, IsWritable: true},
	)
	if param.Base != param.From {
		accounts = append(accounts, types.AccountMeta{PubKey: param.Base, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

type AdvanceNonceAccountParam struct {
	Nonce common.PublicKey
	Auth  common.PublicKey
}

type AdvanceNonceAccountStruct struct {
	Instruction Instruction
}

func AdvanceNonceAccount(param AdvanceNonceAccountParam) types.Instruction {
	data, err := bincode.SerializeData(AdvanceNonceAccountStruct{
		Instruction: InstructionAdvanceNonceAccount,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.Nonce, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
			{PubKey: param.Auth, IsSigner: true, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type WithdrawNonceAccountParam struct {
	Nonce  common.PublicKey
	Auth   common.PublicKey
	To     common.PublicKey
	Amount uint64
}

type WithdrawNonceAccountStruct struct {
	Instruction Instruction
	Lamports    uint64
}

func WithdrawNonceAccount(param WithdrawNonceAccountParam) types.Instruction {
	data, err := bincode.SerializeData(WithdrawNonceAccountStruct{
		Instruction: InstructionWithdrawNonceAccount,
		Lamports:    param.Amount,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.Nonce, IsSigner: false, IsWritable: true},
			{PubKey: param.To, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
			{PubKey: param.Auth, IsSigner: true, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type InitializeNonceAccountParam struct {
	Nonce common.PublicKey
	Auth  common.PublicKey
}

type InitializeNonceAccountStruct struct {
	Instruction Instruction
	Auth        common.PublicKey
}

func InitializeNonceAccount(param InitializeNonceAccountParam) types.Instruction {
	data, err := bincode.SerializeData(InitializeNonceAccountStruct{
		Instruction: InstructionInitializeNonceAccount,
		Auth:        param.Auth,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.Nonce, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type AuthorizeNonceAccountParam struct {
	Nonce   common.PublicKey
	Auth    common.PublicKey
	NewAuth common.PublicKey
}

type AuthorizeNonceAccountStruct struct {
	Instruction Instruction
	Auth        common.PublicKey
}

func AuthorizeNonceAccount(param AuthorizeNonceAccountParam) types.Instruction {
	data, err := bincode.SerializeData(AuthorizeNonceAccountStruct{
		Instruction: InstructionAuthorizeNonceAccount,
		Auth:        param.NewAuth,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.Nonce, IsSigner: false, IsWritable: true},
			{PubKey: param.Auth, IsSigner: true, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type AllocateParam struct {
	Account common.PublicKey
	Space   uint64
}

type AllocateStruct struct {
	Instruction Instruction
	Space       uint64
}

func Allocate(param AllocateParam) types.Instruction {
	data, err := bincode.SerializeData(AllocateStruct{
		Instruction: InstructionAllocate,
		Space:       param.Space,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Account, IsSigner: true, IsWritable: true},
		},
		Data: data,
	}
}

type AllocateWithSeedParam struct {
	Account common.PublicKey
	Base    common.PublicKey
	Owner   common.PublicKey
	Seed    string
	Space   uint64
}

type AllocateWithSeedStruct struct {
	Instruction Instruction
	Base        common.PublicKey
	Seed        string
	Space       uint64
	ProgramID   common.PublicKey
}

func AllocateWithSeed(param AllocateWithSeedParam) types.Instruction {
	data, err := bincode.SerializeData(AllocateWithSeedStruct{
		Instruction: InstructionAllocateWithSeed,
		Base:        param.Base,
		Seed:        param.Seed,
		Space:       param.Space,
		ProgramID:   param.Owner,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Account, IsSigner: false, IsWritable: true},
			{PubKey: param.Base, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}
}

type AssignWithSeedParam struct {
	Account common.PublicKey
	Owner   common.PublicKey
	Base    common.PublicKey
	Seed    string
}

type AssignWithSeedStruct struct {
	Instruction       Instruction
	Base              common.PublicKey
	Seed              string
	AssignToProgramID common.PublicKey
}

func AssignWithSeed(param AssignWithSeedParam) types.Instruction {
	data, err := bincode.SerializeData(AssignWithSeedStruct{
		Instruction:       InstructionAssignWithSeed,
		Base:              param.Base,
		Seed:              param.Seed,
		AssignToProgramID: param.Owner,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Account, IsSigner: false, IsWritable: true},
			{PubKey: param.Base, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}
}

type TransferWithSeedParam struct {
	From   common.PublicKey
	To     common.PublicKey
	Base   common.PublicKey
	Owner  common.PublicKey
	Seed   string
	Amount uint64
}

type TransferWithSeedStruct struct {
	Instruction Instruction
	Lamports    uint64
	Seed        string
	ProgramID   common.PublicKey
}

func TransferWithSeed(param TransferWithSeedParam) types.Instruction {
	data, err := bincode.SerializeData(TransferWithSeedStruct{
		Instruction: InstructionTransferWithSeed,
		Lamports:    param.Amount,
		Seed:        param.Seed,
		ProgramID:   param.Owner,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.From, IsSigner: false, IsWritable: true},
			{PubKey: param.Base, IsSigner: true, IsWritable: false},
			{PubKey: param.To, IsSigner: false, IsWritable: true},
		},
		Data: data,
	}
}

type UpgradeNonceAccountParam struct {
	NonceAccountPubkey common.PublicKey
}

type UpgradeNonceAccountStruct struct {
	Instruction Instruction
}

func UpgradeNonceAccount(param UpgradeNonceAccountParam) types.Instruction {
	data, err := bincode.SerializeData(UpgradeNonceAccountStruct{
		Instruction: InstructionUpgradeNonceAccount,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.NonceAccountPubkey, IsSigner: false, IsWritable: true},
		},
		Data: data,
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
