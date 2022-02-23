package dcf

import (
	"math"
	"testing"
)

const PRECISION = 0.00001

func TestGetCompoundedCAGRAtYear(t *testing.T) {
	input := DCFParameters{
		CAGR: 0.1,
	}

	if diff := math.Abs(input.GetCompoundedCAGRAtYear(1) - 1.1); diff > PRECISION {
		t.Errorf("Expected CAGR for year: %d to be: %.5f but got %.5f", 1, 1.1, input.GetCompoundedCAGRAtYear(1))
	}

	if diff := math.Abs(input.GetCompoundedCAGRAtYear(2) - 1.211); diff > PRECISION {
		t.Errorf("Expected CAGR for year: %d to be: %.5f but got %.5f", 1, 1.21, input.GetCompoundedCAGRAtYear(2))
	}
}

func TestProjectCashFlowAtYear(t *testing.T) {
	input := DCFParameters{
		CAGR: 0.1,
	}
}
