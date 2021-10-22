package math

import "testing"

// special test function, must start with Test<rest-of-function-name>
func TestAverage1(t *testing.T) {
	testcase := []float64{1, 2}
	expected := 1.5
	res := Average(testcase)
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage2(t *testing.T) {
	for _, pair := range tests {
		res := Average(pair.values)
		if res != pair.average {
			t.Error(
				"For:", pair.values,
				"expected:", pair.average,
				"got:", res,
			)
		}
	}
}
