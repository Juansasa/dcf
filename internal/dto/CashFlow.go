package dto

import (
	"dcf-finance.com/v1/api"
	"dcf-finance.com/v1/internal/utils"
)

type CashflowAnnualReport struct {
	FiscalDateEnding                                          string    `json:"fiscalDateEnding"`
	ReportedCurrency                                          string    `json:"reportedCurrency"`
	OperatingCashflow                                         StringInt `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            StringInt `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           StringInt `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              StringInt `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   StringInt `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      StringInt `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       StringInt `json:"capitalExpenditures"`
	ChangeInReceivables                                       StringInt `json:"changeInReceivables"`
	ChangeInInventory                                         StringInt `json:"changeInInventory"`
	ProfitLoss                                                StringInt `json:"profitLoss"`
	CashflowFromInvestment                                    StringInt `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     StringInt `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     StringInt `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        StringInt `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             StringInt `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     StringInt `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            StringInt `json:"dividendPayout"`
	DividendPayoutCommonStock                                 StringInt `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              StringInt `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         StringInt `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet StringInt `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      StringInt `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            StringInt `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           StringInt `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            StringInt `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      StringInt `json:"changeInExchangeRate"`
	NetIncome                                                 StringInt `json:"netIncome"`
}

type CashflowQuarterlyReport struct {
	FiscalDateEnding                                          string    `json:"fiscalDateEnding"`
	ReportedCurrency                                          string    `json:"reportedCurrency"`
	OperatingCashflow                                         StringInt `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            StringInt `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           StringInt `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              StringInt `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   StringInt `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      StringInt `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       StringInt `json:"capitalExpenditures"`
	ChangeInReceivables                                       StringInt `json:"changeInReceivables"`
	ChangeInInventory                                         StringInt `json:"changeInInventory"`
	ProfitLoss                                                StringInt `json:"profitLoss"`
	CashflowFromInvestment                                    StringInt `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     StringInt `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     StringInt `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        StringInt `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             StringInt `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     StringInt `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            StringInt `json:"dividendPayout"`
	DividendPayoutCommonStock                                 StringInt `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              StringInt `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         StringInt `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet StringInt `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      StringInt `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            StringInt `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           StringInt `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            StringInt `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      StringInt `json:"changeInExchangeRate"`
	NetIncome                                                 StringInt `json:"netIncome"`
}

type CashflowReport struct {
	Symbol           string                    `json:"symbol"`
	AnnualReports    []CashflowAnnualReport    `json:"annualReports"`
	QuarterlyReports []CashflowQuarterlyReport `json:"quarterlyReports"`
}

func (report *CashflowReport) FreeCashFlowTTM() int {
	if len(report.QuarterlyReports) < 4 {
		return 0
	}
	cashflowTTM := report.QuarterlyReports[0:4]
	total := 0

	for _, v := range cashflowTTM {
		total += (int(v.OperatingCashflow) - int(v.CapitalExpenditures))
	}

	return total
}

func (report *CashflowReport) Init() error {
	query := api.FinancialQuery{Ticker: report.Symbol, Function: api.CASHFLOW}
	response, err := query.Fetch()
	if err != nil {
		return err
	}

	return utils.ToJSON(response, &report)
}
