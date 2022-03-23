package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"creaturez.nft/questing/v2/questing"
	sendAndConfirmTransaction "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/text"

	"github.com/gagliardetto/solana-go/rpc/ws"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func init() {
	questing.SetProgramID(solana.MustPublicKeyFromBase58("BwM8ipL2fHyBwbMfJSniNvqNKnbyxEXbySKFPRs1AVni"))
}

func main() {
	test_depo()
}

func test_depo() {
	user, err := solana.PrivateKeyFromSolanaKeygenFile("./test.key")
	if err != nil {
		panic(err)
	}

	ballzMint := solana.MustPublicKeyFromBase58("5JSNsqFush4aeZpBcLHhq1KJXSzsTThFNtn9Dubg2V4A")
	ballzAta := solana.MustPublicKeyFromBase58("AbVwyfuZA4R58syxLyRCfHaGP3pvthKY2j4S7qCGHe5p")
	pixelBallzMint := solana.MustPublicKeyFromBase58("QYv41r4wMkMxazSr8UxAh6D6pxWZvTwjXwyos5H8ZUB")
	pixelBallzAta := solana.MustPublicKeyFromBase58("DUjpaABqq2y5ZQAsyaWMNkefq3eVTfbwcg2gW4KpApy3")

	questIndex := 1000234562341
	depositTokenAccount, _ := GetDepositTokenAccount(user.PublicKey(), questIndex)
	questAccount, _ := GetQuestAccount(user.PublicKey(), questIndex)

	qx := questing.NewInitializeQuestInstructionBuilder().
		SetInitializerAccount(user.PublicKey()).
		SetQuestAccountAccount(questAccount).
		SetQuestIndex(uint64(questIndex)).
		SetSystemProgramAccount(solana.SystemProgramID)

	bx := questing.NewBurnBallsInstructionBuilder().
		SetBallzMintAccount(ballzMint).
		SetBallzTokenAccountAccount(ballzAta).
		SetInitializerAccount(user.PublicKey()).
		SetQuestAccountAccount(questAccount).
		SetQuestIndex(uint64(questIndex)).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetTokenProgramAccount(solana.TokenProgramID)

	pbx := questing.NewTransferPixelballzInstructionBuilder().
		SetDepositTokenAccountAccount(depositTokenAccount).
		SetInitializerAccount(user.PublicKey()).
		SetPixelballzMintAccount(pixelBallzMint).
		SetPixelballzTokenAccountAccount(pixelBallzAta).
		SetQuestAccountAccount(questAccount).
		SetQuestIndex(uint64(questIndex)).
		SetRentAccount(solana.SysVarRentPubkey).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetTokenProgramAccount(solana.TokenProgramID)

	sx := questing.NewStartQuestInstructionBuilder().
		SetQuestAccountAccount(questAccount).
		SetInitializerAccount(user.PublicKey()).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetTokenProgramAccount(solana.TokenProgramID).
		SetQuestIndex(uint64(questIndex))

	err = bx.Validate()
	if err != nil {
		panic(err)
	}
	err = pbx.Validate()
	if err != nil {
		panic(err)
	}
	err = qx.Validate()
	if err != nil {
		panic(err)
	}
	err = sx.Validate()
	if err != nil {
		panic(err)
	}

	sendTx(
		"init",
		append(
			make([]solana.Instruction, 0),
			qx.Build(),
			bx.Build(),
			pbx.Build(),
			sx.Build(),
		),
		append(
			make([]solana.PrivateKey, 0),
			user,
		),
		user.PublicKey(),
	)

}

func sendTx(
	doc string,
	instructions []solana.Instruction,
	signers []solana.PrivateKey,
	feePayer solana.PublicKey,
) {
	rpcClient := rpc.New("https://psytrbhymqlkfrhudd.dev.genesysgo.net:8899/")
	wsClient, err := ws.Connect(context.TODO(), "wss://psytrbhymqlkfrhudd.dev.genesysgo.net:8900/")
	if err != nil {
		log.Println("PANIC!!!", fmt.Errorf("unable to open WebSocket Client - %w", err))
	}

	recent, err := rpcClient.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		log.Println("PANIC!!!", fmt.Errorf("unable to fetch recent blockhash - %w", err))
		return
	}

	tx, err := solana.NewTransaction(
		instructions,
		recent.Value.Blockhash,
		solana.TransactionPayer(feePayer),
	)
	if err != nil {
		log.Println("PANIC!!!", fmt.Errorf("unable to create transaction"))
		return
	}

	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		for _, candidate := range signers {
			if candidate.PublicKey().Equals(key) {
				return &candidate
			}
		}
		return nil
	})
	if err != nil {
		log.Println("PANIC!!!", fmt.Errorf("unable to sign transaction: %w", err))
		return
	}

	tx.EncodeTree(text.NewTreeEncoder(os.Stdout, doc))

	sig, err := sendAndConfirmTransaction.SendAndConfirmTransaction(
		context.TODO(),
		rpcClient,
		wsClient,
		tx,
	)
	if err != nil {
		log.Println("PANIC!!!", fmt.Errorf("unable to send transaction - %w", err))
		return
	}
	wsClient.Close()
	log.Println(sig)
}

func GetDepositTokenAccount(
	initializer solana.PublicKey,
	questIndex int,
) (solana.PublicKey, uint8) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(questIndex))
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("questing"),
			[]byte("quest"),
			initializer.Bytes(),
			buf,
		},
		questing.ProgramID,
	)
	fmt.Println("Depo PDA - questing - quest", initializer, questIndex)
	return addr, bump
}

func GetQuestAccount(
	initializer solana.PublicKey,
	questIndex int,
) (solana.PublicKey, uint8) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(questIndex))
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("questing"),
			initializer.Bytes(),
			buf,
		},
		questing.ProgramID,
	)
	fmt.Println("Quest PDA - questing", initializer, questIndex)
	return addr, bump
}
