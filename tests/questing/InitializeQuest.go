// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package questing

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeQuest is the `initializeQuest` instruction.
type InitializeQuest struct {
	QuestIndex *uint64

	// [0] = [WRITE, SIGNER] initializer
	//
	// [1] = [WRITE] questAccount
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeQuestInstructionBuilder creates a new `InitializeQuest` instruction builder.
func NewInitializeQuestInstructionBuilder() *InitializeQuest {
	nd := &InitializeQuest{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetQuestIndex sets the "questIndex" parameter.
func (inst *InitializeQuest) SetQuestIndex(questIndex uint64) *InitializeQuest {
	inst.QuestIndex = &questIndex
	return inst
}

// SetInitializerAccount sets the "initializer" account.
func (inst *InitializeQuest) SetInitializerAccount(initializer ag_solanago.PublicKey) *InitializeQuest {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(initializer).WRITE().SIGNER()
	return inst
}

// GetInitializerAccount gets the "initializer" account.
func (inst *InitializeQuest) GetInitializerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetQuestAccountAccount sets the "questAccount" account.
func (inst *InitializeQuest) SetQuestAccountAccount(questAccount ag_solanago.PublicKey) *InitializeQuest {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(questAccount).WRITE()
	return inst
}

// GetQuestAccountAccount gets the "questAccount" account.
func (inst *InitializeQuest) GetQuestAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeQuest) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeQuest {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeQuest) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst InitializeQuest) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeQuest,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeQuest) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeQuest) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.QuestIndex == nil {
			return errors.New("QuestIndex parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Initializer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.QuestAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *InitializeQuest) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeQuest")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("QuestIndex", *inst.QuestIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  initializer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        quest", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj InitializeQuest) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `QuestIndex` param:
	err = encoder.Encode(obj.QuestIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeQuest) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `QuestIndex`:
	err = decoder.Decode(&obj.QuestIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeQuestInstruction declares a new InitializeQuest instruction with the provided parameters and accounts.
func NewInitializeQuestInstruction(
	// Parameters:
	questIndex uint64,
	// Accounts:
	initializer ag_solanago.PublicKey,
	questAccount ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitializeQuest {
	return NewInitializeQuestInstructionBuilder().
		SetQuestIndex(questIndex).
		SetInitializerAccount(initializer).
		SetQuestAccountAccount(questAccount).
		SetSystemProgramAccount(systemProgram)
}
