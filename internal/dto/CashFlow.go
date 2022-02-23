package dto

import (
	"dcf-finance.com/v1/api"
	"dcf-finance.com/v1/internal/utils"
)

type CashflowAnnualReport struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	ReportedCurrency                                          string `json:"reportedCurrency"`
	OperatingCashflow                                         int    `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            int    `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           int    `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              int    `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   int    `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      int    `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       int    `json:"capitalExpenditures"`
	ChangeInReceivables                                       int    `json:"changeInReceivables"`
	ChangeInInventory                                         int    `json:"changeInInventory"`
	ProfitLoss                                                int    `json:"profitLoss"`
	CashflowFromInvestment                                    int    `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     int    `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     int    `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        int    `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             int    `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     int    `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            int    `json:"dividendPayout"`
	DividendPayoutCommonStock                                 int    `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              int    `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         int    `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet int    `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      int    `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            int    `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           int    `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            int    `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      int    `json:"changeInExchangeRate"`
	NetIncome                                                 int    `json:"netIncome"`
}

type CashflowQuarterlyReport struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	ReportedCurrency                                          string `json:"reportedCurrency"`
	OperatingCashflow                                         int    `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            int    `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           int    `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              int    `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   int    `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      int    `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       int    `json:"capitalExpenditures"`
	ChangeInReceivables                                       int    `json:"changeInReceivables"`
	ChangeInInventory                                         int    `json:"changeInInventory"`
	ProfitLoss                                                int    `json:"profitLoss"`
	CashflowFromInvestment                                    int    `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     int    `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     int    `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        int    `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             int    `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     int    `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            int    `json:"dividendPayout"`
	DividendPayoutCommonStock                                 int    `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              int    `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         int    `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet int    `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      int    `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            int    `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           int    `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            int    `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      int    `json:"changeInExchangeRate"`
	NetIncome                                                 int    `json:"netIncome"`
}

type CashflowReport struct {
	Symbol           string                    `json:"symbol"`
	AnnualReports    []CashflowAnnualReport    `json:"annualReports"`
	QuarterlyReports []CashflowQuarterlyReport `json:"quarterlyReports"`
}

func (report *CashflowAnnualReport) FreeCashFlow() int {
	return report.OperatingCashflow - report.CapitalExpenditures
}

func (report *CashflowQuarterlyReport) FreeCashFlow() int {
	return report.OperatingCashflow - report.CapitalExpenditures
}

func (report *CashflowReport) Init() error {
	query := api.FinancialQuery{Ticker: report.Symbol, Function: api.CASHFLOW}
	response, err := query.Fetch()
	if err != nil {
		return err
	}

	return utils.ToJSON(response, &report)
}
