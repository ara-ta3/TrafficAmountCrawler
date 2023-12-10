package domains

import "github.com/golang-module/carbon/v2"

const START_DAY_OF_MONTH = 4

func CalculateAmount(now carbon.Carbon, currentAmount int) Amount {
	start := now.SetDay(START_DAY_OF_MONTH)
	if now.DayOfMonth() < START_DAY_OF_MONTH {
		start.SubMonth()
	}
	n := start.DiffInDays(now)
	a := float64(currentAmount) / float64(n)
	rest := int(float64(1000-currentAmount) / a)
	end := now.AddDays(rest).ToDateString()
	return Amount{
		CurrentAmount:   currentAmount,
		CurrentDays:     n,
		Average:         a,
		RestDays:        rest,
		ExpectedEndDate: end,
	}
}

type Amount struct {
	CurrentAmount   int
	CurrentDays     int64
	Average         float64
	RestDays        int
	ExpectedEndDate string
}
