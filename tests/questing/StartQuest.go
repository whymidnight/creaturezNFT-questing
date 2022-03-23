// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package questing

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// StartQuest is the `startQuest` instruction.
type StartQuest struct {
	QuestIndex *uint64

	// [0] = [WRITE, SIGNER] initializer
	//
	// [1] = [WRITE] questAccount
	//
	// [2] = [] systemProgram
	//
	// [3] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewStartQuestInstructionBuilder creates a new `StartQuest` instruction builder.
func NewStartQuestInstructionBuilder() *StartQuest {
	nd := &StartQuest{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetQuestIndex sets the "questIndex" parameter.
func (inst *StartQuest) SetQuestIndex(questIndex uint64) *StartQuest {
	inst.QuestIndex = &questIndex
	return inst
}

// SetInitializerAccount sets the "initializer" account.
func (inst *StartQuest) SetInitializerAccount(initializer ag_solanago.PublicKey) *StartQuest {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(initializer).WRITE().SIGNER()
	return inst
}

// GetInitializerAccount gets the "initializer" account.
func (inst *StartQuest) GetInitializerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetQuestAccountAccount sets the "questAccount" account.
func (inst *StartQuest) SetQuestAccountAccount(questAccount ag_solanago.PublicKey) *StartQuest {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(questAccount).WRITE()
	return inst
}

// GetQuestAccountAccount gets the "questAccount" account.
func (inst *StartQuest) GetQuestAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *StartQuest) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *StartQuest {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *StartQuest) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *StartQuest) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *StartQuest {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *StartQuest) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst StartQuest) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_StartQuest,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst StartQuest) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *StartQuest) Validate() error {
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
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *StartQuest) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("StartQuest")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("QuestIndex", *inst.QuestIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  initializer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        quest", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta(" tokenProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj StartQuest) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `QuestIndex` param:
	err = encoder.Encode(obj.QuestIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *StartQuest) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `QuestIndex`:
	err = decoder.Decode(&obj.QuestIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewStartQuestInstruction declares a new StartQuest instruction with the provided parameters and accounts.
func NewStartQuestInstruction(
	// Parameters:
	questIndex uint64,
	// Accounts:
	initializer ag_solanago.PublicKey,
	questAccount ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *StartQuest {
	return NewStartQuestInstructionBuilder().
		SetQuestIndex(questIndex).
		SetInitializerAccount(initializer).
		SetQuestAccountAccount(questAccount).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram)
}
