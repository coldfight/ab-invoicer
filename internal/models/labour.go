package models

type Labour struct {
	Date        Date
	Description string
	Amount      float64
}
type LabourList []Labour

func (l Labour) TotalCost() float64 {
	return l.Amount
}

func (ll LabourList) LabourSubtotal() float64 {
	sum := 0.0
	for _, l := range ll {
		sum += l.TotalCost()
	}
	return sum
}
