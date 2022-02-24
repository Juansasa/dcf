package dto

import (
	"dcf-finance.com/v1/api"
	"dcf-finance.com/v1/internal/utils"
)

type GlobalQuote struct {
	Symbol        string      `json:"01. symbol"`
	Open          StringFloat `json:"02. open"`
	High          StringFloat `json:"03. high"`
	Low           StringFloat `json:"04. low"`
	Price         StringFloat `json:"05. price"`
	Volume        StringInt   `json:"06. volume"`
	PreviousClose StringFloat `json:"08. previous close"`
	Change        StringFloat `json:"09. change"`
}

type Quote struct {
	GlobalQuote GlobalQuote `json:"Global Quote"`
}

func (quote *Quote) Init() error {
	query := api.FinancialQuery{Ticker: quote.GlobalQuote.Symbol, Function: api.GLOBAL_QUOTE}
	response, err := query.Fetch()
	if err != nil {
		return err
	}

	return utils.ToJSON(response, &quote)
}
