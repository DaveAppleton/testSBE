package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/DaveAppleton/ether_go/ethKeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
)

func keyTx(key *ethKeys.AccountKey) *bind.TransactOpts {
	return bind.NewKeyedTransactor(key.GetKey())
}

func role(job string) *bind.TransactOpts {
	return keyTx(roleKey(job))
}

func roleKey(job string) *ethKeys.AccountKey {
	transactorAcc := ethKeys.NewKey("adminKeys/" + job)
	if err := transactorAcc.RestoreOrCreate(); err != nil {
		log.Println(err)
	}
	return transactorAcc
}

var baseClient *backends.SimulatedBackend

func getClient() (client *backends.SimulatedBackend, err error) {
	if baseClient != nil {
		return baseClient, nil
	}
	funds, _ := new(big.Int).SetString("1000000000000000000000000", 10)
	bankerKey := roleKey("banker")
	baseClient = backends.NewSimulatedBackend(core.GenesisAlloc{
		bankerKey.PublicKey(): {Balance: funds},
	})
	return baseClient, nil
}

var test *Test

func getX(ch chan bool) {
	x, err := test.X(nil)
	fmt.Println("X ", x, err)
	ch <- true
}
func getY(ch chan bool) {
	y, err := test.Y(nil)
	fmt.Println("Y ", y, err)
	ch <- true
}

func main() {
	var err error
	client, err := getClient()
	if err != nil {
		log.Fatal(err)
	}
	ownerTx := role("banker")
	_, _, test, err = DeployTest(ownerTx, client)
	client.Commit()
	doneX := make(chan bool)
	doneY := make(chan bool)
	go getX(doneX)
	go getY(doneY)

	<-doneX
	<-doneY

}
