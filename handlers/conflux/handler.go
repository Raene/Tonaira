package conflux

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-co-op/gocron"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/models"
)

func (e *Env) getAddr(ctx *fiber.Ctx) {
	db := e.Config.Db
	cfxTransaction := models.Transaction{}

	err := ctx.BodyParser(&cfxTransaction)
	if err != nil {
		ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	addr, err := generateConfluxAddress()
	if err != nil {
		ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	cfxTransaction.Address = addr
	//insert into database here
	err = cfxTransaction.Create(db)
	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	fmt.Println(cfxTransaction)

	var data = map[string]interface{}{
		"address": addr,
	}

	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(60).Seconds().Do(VerifyTransaction, s1, addr)

	// scheduler starts running jobs and current thread continues to execute
	s1.StartAsync()
	ctx.Status(200).JSON(fiber.Map{
		"data":    data,
		"message": "success",
	})
}

func initAccountManager() *sdk.AccountManager {
	keydir := "./keystore"
	am := sdk.NewAccountManager(keydir)
	return am
}

func initClient() (*sdk.Client, error) {
	url := "http://mainnet-jsonrpc.conflux-chain.org:12537"
	client, err := sdk.NewClient(url)
	if err != nil {
		return client, err
	}
	return client, err
}
func generateConfluxAddress() (types.Address, error) {
	am := initAccountManager()
	var addr types.Address
	client, err := initClient()
	if err != nil {
		return addr, err
	}
	client.SetAccountManager(am)
	addr, err = am.Create("hello")
	if err != nil {
		fmt.Println("create account error", err)
		return addr, err
	}
	return addr, nil
}

/*VerifyTransaction fetches a list of conflux generated addresses from the database
where verified is false
loops through them and checks if they have a balance.
If the balance matches the specified crypto amount in the transaction db
It is transferred to our main account and the user generated account is deleted
*/
func VerifyTransaction(s *gocron.Scheduler, addr types.Address) {
	client, err := initClient()
	if err != nil {
		panic(err)
	}
	am := initAccountManager()
	client.SetAccountManager(am)
	bal, err := client.GetBalance(addr)
	if err != nil {
		panic(err)
	}
	b := big.NewInt(0)
	bigAmount := hexutil.Big(*bal)
	fmt.Println(bal.Cmp(b))
	if bal.Cmp(b) != 0 {

		transToMainAccount(client, am, addr, types.Address("0x17a77c881ff8861507c047db7ecb49b5745274fa"), &bigAmount)
		s.Stop()
	}
}

func transToMainAccount(client *sdk.Client, am *sdk.AccountManager, from types.Address, to types.Address, amount *hexutil.Big) {
	unSignedTx := types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase{
			From:  &from,
			Value: amount,
		},
		To: &to,
	}
	err := client.ApplyUnsignedTransactionDefault(&unSignedTx)
	if err != nil {
		panic(err)
	}
	signedTx, err := am.SignAndEcodeTransactionWithPassphrase(unSignedTx, "hello")
	if err != nil {
		panic(err)
	}
	fmt.Printf("signed tx %+v result:\n0x%x\n\n", unSignedTx, signedTx)

	txhash, err := client.SendRawTransaction(signedTx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("send transaction hash: %v\n\n", txhash)
	_, err = client.GetTransactionByHash(txhash)
	if err != nil {
		panic(err)
	}

}
