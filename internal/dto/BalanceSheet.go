package dto

import (
	"dcf-finance.com/v1/api"
	"dcf-finance.com/v1/internal/utils"
)

type AnnualBalanceSheet struct {
	FiscalDateEnding                       string    `json:"fiscalDateEnding"`
	ReportedCurrency                       string    `json:"reportedCurrency"`
	TotalAssets                            StringInt `json:"totalAssets"`
	TotalCurrentAssets                     StringInt `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  StringInt `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            StringInt `json:"cashAndShortTermInvestments"`
	Inventory                              StringInt `json:"inventory"`
	CurrentNetReceivables                  StringInt `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  StringInt `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 StringInt `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE StringInt `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       StringInt `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      StringInt `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               StringInt `json:"goodwill"`
	Investments                            StringInt `json:"investments"`
	LongTermInvestments                    StringInt `json:"longTermInvestments"`
	ShortTermInvestments                   StringInt `json:"shortTermInvestments"`
	OtherCurrentAssets                     StringInt `json:"otherCurrentAssets"`
	OtherNonCurrrentAssets                 StringInt `json:"otherNonCurrrentAssets"`
	TotalLiabilities                       StringInt `json:"totalLiabilities"`
	TotalCurrentLiabilities                StringInt `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 StringInt `json:"currentAccountsPayable"`
	DeferredRevenue                        StringInt `json:"deferredRevenue"`
	CurrentDebt                            StringInt `json:"currentDebt"`
	ShortTermDebt                          StringInt `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             StringInt `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                StringInt `json:"capitalLeaseObligations"`
	LongTermDebt                           StringInt `json:"longTermDebt"`
	CurrentLongTermDebt                    StringInt `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 StringInt `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 StringInt `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                StringInt `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             StringInt `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 StringInt `json:"totalShareholderEquity"`
	TreasuryStock                          StringInt `json:"treasuryStock"`
	RetainedEarnings                       StringInt `json:"retainedEarnings"`
	CommonStock                            StringInt `json:"commonStock"`
	CommonStockSharesOutstanding           StringInt `json:"commonStockSharesOutstanding"`
}

func (report *AnnualBalanceSheet) SharesOutstanding() int {
	return 123
}

type QuarterBalanceSheet struct {
	FiscalDateEnding                       string    `json:"fiscalDateEnding"`
	ReportedCurrency                       string    `json:"reportedCurrency"`
	TotalAssets                            StringInt `json:"totalAssets"`
	TotalCurrentAssets                     StringInt `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  StringInt `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            StringInt `json:"cashAndShortTermInvestments"`
	Inventory                              StringInt `json:"inventory"`
	CurrentNetReceivables                  StringInt `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  StringInt `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 StringInt `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE StringInt `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       StringInt `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      StringInt `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               StringInt `json:"goodwill"`
	Investments                            StringInt `json:"investments"`
	LongTermInvestments                    StringInt `json:"longTermInvestments"`
	ShortTermInvestments                   StringInt `json:"shortTermInvestments"`
	OtherCurrentAssets                     StringInt `json:"otherCurrentAssets"`
	OtherNonCurrrentAssets                 StringInt `json:"otherNonCurrrentAssets"`
	TotalLiabilities                       StringInt `json:"totalLiabilities"`
	TotalCurrentLiabilities                StringInt `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 StringInt `json:"currentAccountsPayable"`
	DeferredRevenue                        StringInt `json:"deferredRevenue"`
	CurrentDebt                            StringInt `json:"currentDebt"`
	ShortTermDebt                          StringInt `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             StringInt `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                StringInt `json:"capitalLeaseObligations"`
	LongTermDebt                           StringInt `json:"longTermDebt"`
	CurrentLongTermDebt                    StringInt `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 StringInt `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 StringInt `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                StringInt `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             StringInt `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 StringInt `json:"totalShareholderEquity"`
	TreasuryStock                          StringInt `json:"treasuryStock"`
	RetainedEarnings                       StringInt `json:"retainedEarnings"`
	CommonStock                            StringInt `json:"commonStock"`
	CommonStockSharesOutstanding           StringInt `json:"commonStockSharesOutstanding"`
}

type BalanceSheet struct {
	Symbol           string                `json:"symbol"`
	AnnualReports    []AnnualBalanceSheet  `json:"annualReports"`
	QuarterlyReports []QuarterBalanceSheet `json:"quarterlyReports"`
}

func (report *BalanceSheet) SharesOutstanding() int {
	return int(report.QuarterlyReports[0].CommonStockSharesOutstanding)
}

func (balanceSheet *BalanceSheet) Init() error {
	query := api.FinancialQuery{Ticker: balanceSheet.Symbol, Function: api.BALANCE_SHEET}
	data, err := query.Fetch()
	if err != nil {
		return err
	}

	return utils.ToJSON(data, &balanceSheet)
}
