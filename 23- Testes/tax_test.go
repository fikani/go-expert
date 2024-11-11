package tax

/*
Commands:
	go test . -v
	go test . -cover
	go test . -coverprofile=coverage.out
	go tool cover -html=coverage.out
	go test -bench=.
	go test -bench=. -run==^# -v
	go test -bench=. -run==^# -count=10
	go test -fuzz=.
*/
import (
	"math"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 1903.98
	expected := 0.0
	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTaxTest struct {
		income   float64
		expected float64
	}

	table := []calcTaxTest{
		{1903.98, 0.0},
		{2826.65, 69.198750},
		{3751.05, 207.857500},
		{4664.68, 413.423000},
		{5000.0, 505.640000},
	}

	for _, c := range table {
		result := CalculateTax(c.income)
		if !floatEqual(result, c.expected) {
			t.Errorf("Expected %.10f, got %.10f", c.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(5000.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(5000.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500, 1000, 1501}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Error("Result shoulbe 0 for <= 0")
		}
	})
}

const epsilon = 1e-9 // Define a small tolerance value

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
