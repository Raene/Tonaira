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
	s1.Every(3).Seconds().Do(task, s1, addr)

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

func task(s *gocron.Scheduler, addr types.Address) {
	fmt.Println(addr)
	x := 2
	if x == 1 {
		s.Stop()
	}
}
