package models

import (
	"fmt"
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-co-op/gocron"
	"github.com/jinzhu/gorm"
)

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
func GenerateConfluxAddress() (types.Address, error) {
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
func VerifyTransaction(s *gocron.Scheduler, tx Transaction, db *gorm.DB) {
	client, err := initClient()
	if err != nil {
		panic(err)
	}
	am := initAccountManager()
	client.SetAccountManager(am)

	bal, err := client.GetBalance(tx.Address)
	if err != nil {
		fmt.Println(err)
		s.Stop()
	}
	b := big.NewInt(0)
	bigAmount := hexutil.Big(*bal)
	fmt.Println(bal.Cmp(b))
	if bal.Cmp(b) != 0 {
		TransToMainAccount(client, am, tx.Address, types.Address("0x17a77c881ff8861507c047db7ecb49b5745274fa"), &bigAmount)
		//set account status in db to true
		tx.Status = true
		err := tx.Update(db)
		if err != nil {
			fmt.Println(err)
		}
		s.Stop()
	}

}

//new logic, fetch all transactions from database and start a veriyTransaction for each of them, this should be done in the main function as a subroutine, update and move above logic to a file called Cron

func TransToMainAccount(client *sdk.Client, am *sdk.AccountManager, from types.Address, to types.Address, amount *hexutil.Big) {
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

func SpawnConfluxCron(db *gorm.DB) {
	t := Transaction{Status: false}
	txs, errs := t.GetWhere(db)
	if len(errs) != 0 {
		fmt.Println(errs)
	}
	//fetch accounts where paid is false
	s1 := gocron.NewScheduler(time.UTC)
	for _, tx := range txs {
		s1.Every(60).Seconds().Do(VerifyTransaction, s1, tx, db)
		s1.StartAsync()
	}
}
