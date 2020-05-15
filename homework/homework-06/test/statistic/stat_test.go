package statistic

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var testsAverage = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}
var testsSum = []testpair{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range testsAverage {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestSumm(t *testing.T) {

	for _, pair := range testsSum {
		v := Summ(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
