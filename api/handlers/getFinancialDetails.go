package handlers

import (
	"net/http"

	"dcf-finance.com/v1/api"
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
