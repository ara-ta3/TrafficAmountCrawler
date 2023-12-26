package domains

import "github.com/golang-module/carbon/v2"

const START_DAY_OF_MONTH = 4

const MAX_AMOUNT = 1000

func CalculateAmount(now carbon.Carbon, currentAmount int) Amount {
	start := now.SetDay(START_DAY_OF_MONTH)
	if now.DayOfMonth() < START_DAY_OF_MONTH {
		start = start.SubMonth()
	}
	end := start.AddMonth()
	restUntilEnd := now.DiffInDays(end)

	n := start.DiffInDays(now)
	a := float64(currentAmount) / float64(n)
	rest := int(float64(MAX_AMOUNT-currentAmount) / a)
	traficEnd := now.AddDays(rest)
	return Amount{
		CurrentAmount:    currentAmount,
		CurrentDays:      n,
		Average:          a,
		RestDaysUntilEnd: restUntilEnd,
		End:              end,
		RestDays:         rest,
		ExpectedEnd:      traficEnd,
	}
}

type Amount struct {
	CurrentAmount    int
	CurrentDays      int64
	Average          float64
	RestDaysUntilEnd int64
	End              carbon.Carbon
	RestDays         int
	ExpectedEnd      carbon.Carbon
}

func (a Amount) ExpireByTheLastDay() bool {
	return a.ExpectedEnd.Lt(a.End)
}
