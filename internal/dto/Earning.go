package dto

type AnnualEarning struct {
	FiscalDateEnding string `json:"fiscalDateEnding"`
	ReportedEPS      string `json:"reportedEPS"`
}

type QuarterlyEarning struct {
	FiscalDateEnding   string `json:"fiscalDateEnding"`
	ReportedDate       string `json:"reportedDate"`
	ReportedEPS        string `json:"reportedEPS"`
	EstimatedEPS       string `json:"estimatedEPS"`
	Surprise           string `json:"surprise"`
	SurprisePercentage string `json:"surprisePercentage"`
}

type Earning struct {
	Symbol            string             `json:"symbol"`
	AnnualEarnings    []AnnualEarning    `json:"annualEarnings"`
	QuarterlyEarnings []QuarterlyEarning `json:"quarterlyEarnings"`
}
