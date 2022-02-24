package dcf

import (
	"fmt"
	"math"

	"dcf-finance.com/v1/internal/dto"
)

type CAGR struct {
	Rate        float64 `json:"rate"`
	Probability float64 `json:"probability"`
}

type ProjectedCAGR struct {
	*CAGR
	Value float64
}

type DCFParameters struct {
	Ticker                string  `json:"ticker"`
	CAGR                  float64 `json:"cagr"`
	Probability           float64 `json:"probability"`
	DiscountRate          float64 `json:"dicountRate"`
	PerpetualGrowthRate   float64 `json:"perpetualGrowthRate"`
	ExitCashFlowMultiples int     `json:"exitCashFlowMultiple"`
	ProjectedYears        int     `json:"projectedYears"`
}

func (dcfParams *DCFParameters) Summary() (DCF, error) {
	symbol := dcfParams.Ticker
	cashflow := dto.CashflowReport{Symbol: symbol}
	balanceSheet := dto.BalanceSheet{Symbol: symbol}

	err := cashflow.Init()
	if err != nil {
		fmt.Println(err)
		return DCF{}, err
	}

	err = balanceSheet.Init()
	if err != nil {
		fmt.Println(err)
		return DCF{}, err
	}

	return DCF{
		DCFParameters:     dcfParams,
		Symbol:            symbol,
		FreeCashFlowTTM:   cashflow.FreeCashFlowTTM(),
		SharesOutstanding: balanceSheet.SharesOutstanding(),
	}, nil
}

type DCF struct {
	*DCFParameters
	Symbol            string `json:"symbol"`
	FreeCashFlowTTM   int    `json:"freeCashFlowTTM"`
	SharesOutstanding int    `json:"sharesOutstanding"`
}

func (summary *DCF) ProjectedCash() []int {
	cash := make([]int, summary.ProjectedYears+1)
	for i := 0; i <= summary.ProjectedYears; i++ {
		compoundGrowth := math.Pow(1+summary.CAGR, float64(i)) * float64(summary.FreeCashFlowTTM)
		cash[i] = int(compoundGrowth)
	}

	return cash
}

func (summary *DCF) TerminalPerpetualCash() int {
	endCash := summary.ProjectedCash()[summary.ProjectedYears]
	return int(float64(endCash) * (1 + summary.PerpetualGrowthRate) / (summary.DiscountRate - summary.PerpetualGrowthRate))
}

func (summary *DCF) TerminalExitMultipleCash() int {
	endCash := summary.ProjectedCash()[summary.ProjectedYears]
	return endCash * summary.ExitCashFlowMultiples
}

func (summary *DCF) TerminalCash() int {
	return (summary.TerminalExitMultipleCash() + summary.TerminalPerpetualCash()) / 2
}

func (summary *DCF) FairPrice() int {
	return summary.TerminalCash() / summary.SharesOutstanding
}
