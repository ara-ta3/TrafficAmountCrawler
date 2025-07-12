package domains

import (
	"testing"

	"github.com/dromara/carbon/v2"
)

func TestCalculateAmountRestDays(t *testing.T) {
	now := carbon.Parse("2024-05-04", carbon.UTC)
	t.Logf("Input date: %s", now.ToDateString())
	a := CalculateAmount(now, 1000)
	actual := a.RestDays()

	t.Logf("Period: Begin=%s, End=%s", a.Period.Begin.ToDateString(), a.Period.End.ToDateString())
	t.Logf("CurrentDate: %s", a.CurrentDate.ToDateString())
	t.Logf("RestDays: %d", actual)

	if actual != 30 {
		t.Errorf("%+v is not 30", actual)
	}
}

func TestCalculateAmountAverageUsedAmountShouldBe0IfTodayEqualsToStartDay(t *testing.T) {
	now := carbon.Parse("2024-05-04", carbon.UTC)
	a := CalculateAmount(now, 1000)
	actual := a.AverageUsedAmount()

	t.Logf("Period: Begin=%s, End=%s", a.Period.Begin.ToDateString(), a.Period.End.ToDateString())
	t.Logf("CurrentDate: %s", a.CurrentDate.ToDateString())
	t.Logf("UsedDays: %d", a.UsedDays())
	t.Logf("AverageUsedAmount: %f", actual)

	if actual != 0 {
		t.Errorf("%+v is not 0", actual)
	}
}
