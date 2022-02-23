package dto

import (
	"testing"
)

func TestFreeCashFlow(t *testing.T) {
	report := CashflowAnnualReport{OperatingCashflow: 10, CapitalExpenditures: 1}
	if report.FreeCashFlow() != 9 {
		t.Error("Expected FCF to be 9")
	}
}
