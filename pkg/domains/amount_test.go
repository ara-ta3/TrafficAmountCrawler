package domains

import (
	"testing"

	"github.com/golang-module/carbon/v2"
)

func TestCalculateAmountRestDays(t *testing.T) {
	a := CalculateAmount(carbon.Parse("2024-05-04"), 1000)
	actual := a.RestDays()

	if actual != 30 {
		t.Errorf("%+v is not 30", actual)
	}
}

func TestCalculateAmountAverageUsedAmountShouldBe0IfTodayEqualsToStartDay(t *testing.T) {
	a := CalculateAmount(carbon.Parse("2024-05-04"), 1000)
	actual := a.AverageUsedAmount()

	if actual != 0 {
		t.Errorf("%+v is not 0", actual)
	}
}
