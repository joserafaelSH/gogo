package model

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
