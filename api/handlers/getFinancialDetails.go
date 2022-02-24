package handlers

import (
	"net/http"

	"dcf-finance.com/v1/api"
	"dcf-finance.com/v1/internal/dcf"
	"dcf-finance.com/v1/internal/dto"
	"dcf-finance.com/v1/internal/utils"
	"github.com/gin-gonic/gin"
)

func getDtoType(apiFunction string) interface{} {
	switch apiFunction {
	case api.CASHFLOW:
		return dto.CashflowReport{}
	case api.INCOME_STATEMENT:
		return dto.IncomeStatement{}
	case api.BALANCE_SHEET:
		return dto.BalanceSheet{}
	case api.OVERVIEW:
		return dto.CompanyOverview{}
	case api.EARNINGS:
		return dto.Earning{}
	}

	return nil
}

func GetFinancial(c *gin.Context) {
	ticker := c.Param("ticker")
	apiFunction := c.Param("function")
	query := api.FinancialQuery{Function: apiFunction, Ticker: ticker}
	response, err := query.Fetch()
	if err != nil {
		c.Error(err)
	}
	cashflow := getDtoType(apiFunction)
	err = utils.ToJSON(response, &cashflow)
	if err != nil {
		c.Error(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"response": cashflow})
}

func GetDCF(c *gin.Context) {
	ticker := c.Param("ticker")
	dcfParams := dcf.DCFParameters{
		Ticker: ticker,
		/*
			CAGR:                  0.2,
			Probability:           0.5,
			DiscountRate:          0.08,
			PerpetualGrowthRate:   0.02,
			ExitCashFlowMultiples: 15,
			ProjectedYears:        5,
		*/
	}

	err := c.BindJSON(&dcfParams)
	if err != nil {
		c.Error(err)
	}

	summary, err := dcfParams.Summary()
	if err != nil {
		c.Error(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"Summary":               summary,
		"ProjectedCashFlow":     summary.ProjectedCash(),
		"TerminalPerpetualCash": summary.TerminalPerpetualCash(),
		"TerminalExitCash":      summary.TerminalExitMultipleCash(),
		"TerminalCashAverage":   summary.TerminalCash(),
		"Fair Price":            summary.FairPrice(),
		"Current Price":         summary.CurrentPrice,
		"Upside":                summary.Upside(),
	})
}
