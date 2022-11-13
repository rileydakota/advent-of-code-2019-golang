package day1

import (
	"strconv"
	"testing"
)

func TestCalcFuel(t *testing.T) {
	test1 := calcFuel(14)

	if test1 != 2 {
		str := strconv.Itoa(test1)
		t.Errorf("Got %s, wanted 14", str)
	}

	test2 := calcFuel(2)

	if test2 != 0 {
		str := strconv.Itoa(test2)
		t.Errorf("Got %s, wanted 0", str)
	}

	test3 := calcFuel(1969)

	if test3 != 654 {
		str := strconv.Itoa(test3)
		t.Errorf("Got %s, wanted 0", str)
	}

	test4 := calcFuel(216)

	if test4 != 70 {
		str := strconv.Itoa(test4)
		t.Errorf("Got %s, wanted 70", str)
	}

}

func TestSolvepart1(t *testing.T) {
	testdata := []string{
		"12",
		"14",
		"1969",
		"100756",
	}

	got, err := Solvepart1(testdata)
	if err != nil {
		t.Errorf("%q", err)
		return
	}
	want := 34241

	if got != want {
		t.Errorf("Got %s, wanted %s", strconv.Itoa(got), strconv.Itoa(want))
	}
}

func TestSolvepart2(t *testing.T) {
	testdata := []string{
		"14",
		"1969",
		"100756",
	}

	got, err := Solvepart2(testdata)
	if err != nil {
		t.Errorf("%q", err)
		return
	}
	want := 51314

	if got != want {
		t.Errorf("Got %s, wanted %s", strconv.Itoa(got), strconv.Itoa(want))
	}
}
