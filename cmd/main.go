package main

import (
	"fmt"

	"github.com/joserafaelsh/syncrhonize-database/app/model"
	"github.com/joserafaelsh/syncrhonize-database/app/services"
	"github.com/joserafaelsh/syncrhonize-database/infra/db"
)

func main() {

	db1 := db.ConnectDB1()
	db2 := db.ConnectDB2()

	db1.Ping()
	db2.Ping()

	createEnergyReads := model.CreateEnergyReadData()
	services.CreateEnergyRead(db1, createEnergyReads)
	energyReads := services.GetEnergyRead(db1)
	services.CreateFaturaBeneficiaria(db2, energyReads)

	defer db1.Close()
	defer db2.Close()

	fmt.Println("the end its over")

}
