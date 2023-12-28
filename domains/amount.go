package domains

import "github.com/golang-module/carbon/v2"

const START_DAY_OF_MONTH = 4

const MAX_AMOUNT = 1000

func CalculateAmount(now carbon.Carbon, currentAmount int) Amount {
	begin := now.SetDay(START_DAY_OF_MONTH)
	if now.DayOfMonth() < START_DAY_OF_MONTH {
		begin = begin.SubMonth()
	}
	end := begin.AddMonth()
	return Amount{
		Period: Period{
			Begin: begin,
			End:   end,
		},
		CurrentAmount: currentAmount,
		CurrentDate:   now,
	}
}

type Period struct {
	Begin carbon.Carbon
	End   carbon.Carbon
}

type Amount struct {
	Period        Period
	CurrentAmount int
	CurrentDate   carbon.Carbon
}

func (a Amount) UsedDays() int64 {
	return a.Period.Begin.DiffInDays(a.CurrentDate)
}

func (a Amount) AverageUsedAmount() float64 {
	return float64(a.CurrentAmount) / float64(a.UsedDays())
}

func (a Amount) RestAmount() int64 {
	return int64(MAX_AMOUNT - a.CurrentAmount)
}

func (a Amount) RestDays() int64 {
	return a.CurrentDate.DiffInDays(a.Period.End)
}

func (a Amount) AverageRestAmount() float64 {
	return float64(a.RestAmount()) / float64(a.RestDays())
}

func (a Amount) ExpectedRestDays() int64 {
	return a.RestAmount() / int64(a.AverageUsedAmount())
}

func (a Amount) ExpectedEndDate() carbon.Carbon {
	return a.CurrentDate.AddDays(int(a.ExpectedRestDays()))
}
