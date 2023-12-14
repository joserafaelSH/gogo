package services

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/joserafaelsh/syncrhonize-database/app/model"
)

/*
type EnergyRead struct {
	ID                    string  `json:"id"`
	DateRef               string  `json:"date_ref"`
	DatePayment           string  `json:"date_payment"`
	DateDue               string  `json:"date_due"`
	KwhCompensado         float64 `json:"kwh_compensado"`
	KwhBilled             float64 `json:"kwh_billed"`
	KwhConsumed           float64 `json:"kwh_consumed"`
	KwhAcumulatedCreditP  float64 `json:"kwh_acumulated_credit_p"`
	KwhAcumulatedCreditFp float64 `json:"kwh_acumulated_credit_fp"`
	DealershipBillCost    float64 `json:"dealership_bill_cost"`
	DealershipExtraFees   float64 `json:"dealership_extra_fees"`
	DealershipBillStatus  string  `json:"dealership_bill_status"`
	IsMicroMinGenerator   bool    `json:"is_micro_min_generator"`
	FileKey               string  `json:"file_key"`
	CreatedAt             string  `json:"created_at"`
	UpdatedAt             string  `json:"updated_at"`
	ClientID              string  `json:"client_id"`
	ScraperStatus         string  `json:"scraper_status"`
	ValorCompensado       float64 `json:"valor_compensado"`
	ValorEnergia          float64 `json:"valor_energia"`
	DealershipBillID      string  `json:"dealership_bill_id"`
	ConsumoMinimo         float64 `json:"consumo_minimo"`
}

type FaturaBeneficiaria struct {
	ID                   string  `json:"id"`
	DataVencimento       string  `json:"data_vencimento"`
	ValorCompensado      float64 `json:"valor_compensado"`
	KwhCompensado        float64 `json:"kwh_compensado"`
	BeneficiariaID       string  `json:"beneficiaria_id"`
	DataRef              string  `json:"data_ref"`
	FileKey              string  `json:"file_key"`
	KwhAcumuladoFp       float64 `json:"kwh_acumulado_fp"`
	KwhAcumuladoP        float64 `json:"kwh_acumulado_p"`
	KwhBilled            float64 `json:"kwh_billed"`
	KwhConsumido         float64 `json:"kwh_consumido"`
	MicroMiniGeracao     bool    `json:"micro_mini_geracao"`
	MicroMiniGeracaoKwh  float64 `json:"micro_mini_geracao_kwh"`
	ValorEnergia         float64 `json:"valor_energia"`
	UC                   string  `json:"uc"`
	DataPagamento        string  `json:"data_pagamento"`
	EnergiaInjetadaKwh   float64 `json:"energia_injetada_kwh"`
	EnergiaInjetadaReais float64 `json:"energia_injetada_reais"`
	SaldoKwh             float64 `json:"saldo_kwh"`
	SaldoReais           float64 `json:"saldo_reais"`
	ConsumoMinimo        float64 `json:"consumo_minimo"`
}

*/

func CreateEnergyRead(db *sql.DB, data []model.EnergyRead) {

	for _, energyRead := range data {

		_, err := db.Exec("INSERT INTO energy_read (id, date_ref, date_payment, date_due, kwh_compensado, kwh_billed, kwh_consumed, kwh_acumulated_credit_p, kwh_acumulated_credit_fp, dealership_bill_cost, dealership_extra_fees, dealership_bill_status, is_micro_min_generator, file_key, created_at, updated_at, client_id, scraper_status, valor_compensado, valor_energia, dealership_bill_id, consumo_minimo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10 , $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21,$22)",

			energyRead.ID,
			energyRead.DateRef,
			energyRead.DatePayment,
			energyRead.DateDue,
			energyRead.KwhCompensado,
			energyRead.KwhBilled,
			energyRead.KwhConsumed,
			energyRead.KwhAcumulatedCreditP,
			energyRead.KwhAcumulatedCreditFp,
			energyRead.DealershipBillCost,
			energyRead.DealershipExtraFees,
			energyRead.DealershipBillStatus,
			energyRead.IsMicroMinGenerator,
			energyRead.FileKey,
			energyRead.CreatedAt,
			energyRead.UpdatedAt,
			energyRead.ClientID,
			energyRead.ScraperStatus,
			energyRead.ValorCompensado,
			energyRead.ValorEnergia,
			energyRead.DealershipBillID,
			energyRead.ConsumoMinimo)
		if err != nil {

			panic(err.Error())
		}
	}
}

func GetEnergyRead(db *sql.DB) []model.EnergyRead {

	rows, err := db.Query("SELECT * FROM energy_read")
	if err != nil {
		panic(err.Error())
	}
	energyReadArray := []model.EnergyRead{}

	var dateRef, datePayment, dateDue, createdAt, updatedAt interface{}

	for rows.Next() {

		var energyRead model.EnergyRead

		err = rows.Scan(
			&energyRead.ID,
			&dateRef,
			&datePayment,
			&dateDue,
			&energyRead.KwhCompensado,
			&energyRead.KwhBilled,
			&energyRead.KwhConsumed,
			&energyRead.KwhAcumulatedCreditP,
			&energyRead.KwhAcumulatedCreditFp,
			&energyRead.DealershipBillCost,
			&energyRead.DealershipExtraFees,
			&energyRead.DealershipBillStatus,
			&energyRead.IsMicroMinGenerator,
			&energyRead.FileKey,
			&createdAt,
			&updatedAt,
			&energyRead.ClientID,
			&energyRead.ScraperStatus,
			&energyRead.ValorCompensado,
			&energyRead.ValorEnergia,
			&energyRead.DealershipBillID,
			&energyRead.ConsumoMinimo)

		if err != nil {
			panic(err.Error())
		}

		energyRead.DateRef, _ = dateRef.(time.Time)
		energyRead.DatePayment, _ = datePayment.(time.Time)
		energyRead.DateDue, _ = dateDue.(time.Time)
		energyReadArray = append(energyReadArray, energyRead)
	}

	return energyReadArray
}

func CreateFaturaBeneficiaria(db *sql.DB, data []model.EnergyRead) {

	for _, energyRead := range data {

		_, err := db.Exec(`INSERT INTO fatura_beneficiaria 
			(id, 
			data_vencimento, 
			valor_compensado, 
			kwh_compensado, 
			beneficiaria_id, 
			data_ref, 
			file_key, 
			kwh_acumulado_fp, 
			kwh_acumulado_p, 
			kwh_billed, 
			kwh_consumido, 
			micro_mini_geracao, 
			micro_mini_geracao_kwh, 
			valor_energia, 
			uc, 
			data_pagamento, 
			energia_injetada_kwh, 
			energia_injetada_reais, 
			saldo_kwh, 
			saldo_reais, 
			consumo_minimo) 
			
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9 , $10, $11, $12, $13, $14, $15, $16, $17, $18, $19,$20, $21)`,

			uuid.New().String(),
			energyRead.DateDue,
			energyRead.ValorCompensado,
			energyRead.KwhCompensado,
			energyRead.ClientID,
			energyRead.DateRef,
			energyRead.FileKey,
			energyRead.KwhAcumulatedCreditFp,
			energyRead.KwhAcumulatedCreditP,
			energyRead.KwhBilled,
			energyRead.KwhConsumed,
			energyRead.IsMicroMinGenerator,
			0,
			energyRead.ValorEnergia,
			"UC123",
			energyRead.DatePayment,
			0,
			energyRead.ValorCompensado,
			energyRead.ValorCompensado,
			energyRead.ValorCompensado,
			energyRead.ConsumoMinimo)
		if err != nil {

			panic(err.Error())
		}
	}
}
