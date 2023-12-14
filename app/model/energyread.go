package model

import (
	"time"

	"github.com/google/uuid"
)

type EnergyRead struct {
	ID                    string    `json:"id"`
	DateRef               time.Time `json:"date_ref"`
	DatePayment           time.Time `json:"date_payment"`
	DateDue               time.Time `json:"date_due"`
	KwhCompensado         float64   `json:"kwh_compensado"`
	KwhBilled             float64   `json:"kwh_billed"`
	KwhConsumed           float64   `json:"kwh_consumed"`
	KwhAcumulatedCreditP  float64   `json:"kwh_acumulated_credit_p"`
	KwhAcumulatedCreditFp float64   `json:"kwh_acumulated_credit_fp"`
	DealershipBillCost    float64   `json:"dealership_bill_cost"`
	DealershipExtraFees   float64   `json:"dealership_extra_fees"`
	DealershipBillStatus  string    `json:"dealership_bill_status"`
	IsMicroMinGenerator   bool      `json:"is_micro_min_generator"`
	FileKey               string    `json:"file_key"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	ClientID              string    `json:"client_id"`
	ScraperStatus         string    `json:"scraper_status"`
	ValorCompensado       float64   `json:"valor_compensado"`
	ValorEnergia          float64   `json:"valor_energia"`
	DealershipBillID      string    `json:"dealership_bill_id"`
	ConsumoMinimo         float64   `json:"consumo_minimo"`
}

func CreateEnergyReadData() []EnergyRead {
	newEnergyRead := []EnergyRead{}

	for i := 0; i < 10000; i++ {
		newEnergyRead = append(newEnergyRead, EnergyRead{
			ID:                    uuid.New().String(),
			DateRef:               time.Now(),
			DatePayment:           time.Now(),
			DateDue:               time.Now(),
			KwhCompensado:         1,
			KwhBilled:             1,
			KwhConsumed:           1,
			KwhAcumulatedCreditP:  1,
			KwhAcumulatedCreditFp: 1,
			DealershipBillCost:    1,
			DealershipExtraFees:   1,
			DealershipBillStatus:  "1",
			IsMicroMinGenerator:   true,
			FileKey:               "1",
			CreatedAt:             time.Now(),
			UpdatedAt:             time.Now(),
			ClientID:              uuid.New().String(),
			ScraperStatus:         "1",
			ValorCompensado:       1,
			ValorEnergia:          1,
			DealershipBillID:      uuid.New().String(),
			ConsumoMinimo:         1,
		})
	}
	return newEnergyRead
}
