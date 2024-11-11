package tax

import "time"

func CalculateTax(income float64) float64 {
	if income <= 0 {
		return 0
	}

	if income <= 1903.98 {
		return 0
	} else if income <= 2826.65 {
		return (income * 0.075) - 142.80
	} else if income <= 3751.05 {
		return (income * 0.15) - 354.80
	} else if income <= 4664.68 {
		return (income * 0.225) - 636.13
	} else {
		return (income * 0.275) - 869.36
	}
}

func CalculateTax2(income float64) float64 {
	time.Sleep(time.Millisecond)
	if income <= 0 {
		return 0
	}

	if income <= 1903.98 {
		return 0
	} else if income <= 2826.65 {
		return (income * 0.075) - 142.80
	} else if income <= 3751.05 {
		return (income * 0.15) - 354.80
	} else if income <= 4664.68 {
		return (income * 0.225) - 636.13
	} else {
		return (income * 0.275) - 869.36
	}
}
