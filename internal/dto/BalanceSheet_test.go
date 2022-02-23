package dto

import "testing"

func TestAnnualSharesOutstanding(t *testing.T) {
	report := AnnualBalanceSheet{}
	if report.SharesOutstanding() != 0 {
		t.Error("Expected FCF to be 0")
	}
}

func TestAnnualSharesOutstandingCalc(t *testing.T) {
	report := AnnualBalanceSheet{CommonStock: 11, CommonStockSharesOutstanding: 23, TreasuryStock: 17}
	if report.SharesOutstanding() != 17 {
		t.Error("Expected FCF to be 17")
	}
}

func TestQuarterSharesOutstanding(t *testing.T) {
	report := QuarterBalanceSheet{}
	if report.SharesOutstanding() != 0 {
		t.Error("Expected FCF to be 0")
	}
}

func TestQuarterSharesOutstandingCalc(t *testing.T) {
	report := QuarterBalanceSheet{CommonStock: 11, CommonStockSharesOutstanding: 23, TreasuryStock: 17}
	if report.SharesOutstanding() != 17 {
		t.Error("Expected FCF to be 17")
	}
}
