package dcf

import (
	"math"
)

type CAGR struct {
	Rate        float64 `json:"rate"`
	Probability float64 `json:"probability"`
}

type DCFParameters struct {
	Ticker                string  `json:"ticker"`
	CAGR                  []CAGR  `json:"cagr"`
	DiscountRate          float64 `json:"dicountRate"`
	PerpetualGrowthRate   float64 `json:"perpetualGrowthRate"`
	ExitCashFlowMultiples int     `json:"exitCashFlowMultiple"`
}

type ProjectedCAGR struct {
	*CAGR
	Value float64
}

func (input *DCFParameters) ProjectCAGR(nrOfYears int) []ProjectedCAGR {
	rates := []ProjectedCAGR{}
	for i := 1; i <= nrOfYears; i++ {
		for j := 0; j <= len(input.CAGR); j++ {
			cagr := input.CAGR[j]
			compoundedRate := math.Pow(1+cagr.Rate, float64(i))
			rates[i] = ProjectedCAGR{
				CAGR:  &cagr,
				Value: compoundedRate,
			}
		}
	}

	return rates
}
