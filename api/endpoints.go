package api

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const API_KEY1 = "NP26C2EFBOBNTUWL"
const API_KEY2 = "CFB31V1MI5HDKIA8"
const BASE_URL = "https://www.alphavantage.co/query?apikey=" + API_KEY2

const (
	CASHFLOW          = "CASH_FLOW"
	INCOME_STATEMENT  = "INCOME_STATEMENT"
	BALANCE_SHEET     = "BALANCE_SHEET"
	EARNINGS          = "EARNINGS"
	OVERVIEW          = "OVERVIEW"
	EARNINGS_CALENDAR = "EARNINGS_CALENDAR"
	GLOBAL_QUOTE      = "GLOBAL_QUOTE"
)

type FinancialQuery struct {
	Function string
	Ticker   string
}

func (query *FinancialQuery) Url() (string, error) {
	url, err := url.Parse(BASE_URL)
	if err != nil {
		return "", err
	}

	queries := url.Query()
	queries.Set("function", query.Function)
	queries.Set("symbol", query.Ticker)
	url.RawQuery = queries.Encode()

	return url.String(), nil
}

func (query *FinancialQuery) Fetch() ([]byte, error) {
	url, err := query.Url()
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}
