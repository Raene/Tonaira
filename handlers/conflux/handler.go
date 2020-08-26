package conflux

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/models"
)

func (e *Env) getAddr(ctx *fiber.Ctx) {
	//db := e.Config.Db
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

func generateConfluxAddress() (types.Address, error) {
	am := sdk.NewAccountManager("./keystore")
	url := "http://mainnet-jsonrpc.conflux-chain.org:12537"
	client, err := sdk.NewClient(url)
	if err != nil {
		panic(err)
	}
	client.SetAccountManager(am)
	addr, err := am.Create("hello")
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
	fmt.Println(addr)
	x := 1
	if x == 1 {
		s.Stop()
	}
}
