package bmicalculator_test

import (
	"fmt"
	"testing"

	"github.com/robertoesteves13/bmi-calculator"
)

type CalculatorTest struct {
	Weight      float64
	Height      float64
	ExpectedBMI string
}

func TestCalculatBMI(t *testing.T) {
	cases := []CalculatorTest{
		{Weight: 66, Height: 1.75, ExpectedBMI: "21.55"},
		{Weight: 72, Height: 1.55, ExpectedBMI: "29.97"},
		{Weight: 80, Height: 1.9, ExpectedBMI: "22.16"},
		{Weight: 80, Height: 0, ExpectedBMI: "+Inf"},
	}

	for i := 0; i < len(cases); i++ {
		bmi := bmicalculator.CalculateBMI(cases[i].Weight, cases[i].Height)

		// INFO(roberto): We are truncationg with the precision of 2 decimal places,
		// so testing can be simpler because of floating points being so precise
		// (also who writes unit tests for one-line functions lmao).
		if fmt.Sprintf("%.2f", bmi) != cases[i].ExpectedBMI {
			t.Logf("Expected %s, but got %.2f", cases[i].ExpectedBMI, bmi)
			t.FailNow()
		}
	}
}
