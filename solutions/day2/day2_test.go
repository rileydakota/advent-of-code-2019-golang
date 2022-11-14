package day2

import (
	"log"
	"strconv"
	"testing"
)

func TestSolvepart1(t *testing.T) {
	inputdata := []string{"1,9,10,3,2,3,11,0,99,30,40,50"}

	want := 3500
	got, err := Solvepart1(inputdata)
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("%v", inputdata)

	if got[0] != want {
		t.Errorf("Got %s, wanted %s", strconv.Itoa(got[0]), strconv.Itoa(want))
	}

	inputdata = []string{"1,1,1,4,99,5,6,0,99"}
	want = 30
	got, err = Solvepart1(inputdata)
	if err != nil {
		t.Error(err)
		return
	}

	if got[0] != want {
		t.Errorf("Got %s, wanted %s", strconv.Itoa(got[0]), strconv.Itoa(want))
	}

	inputdata = []string{"2,4,4,5,99,0"}
	want = 2
	got, err = Solvepart1(inputdata)
	if err != nil {
		t.Error(err)
		return
	}

	if got[0] != want {
		t.Errorf("Got %s, wanted %s", strconv.Itoa(got[0]), strconv.Itoa(want))
	}

	if got[5] != 9801 {
		t.Errorf("Got %s, wanted %s", strconv.Itoa(got[0]), "9801")
	}

	inputdata = []string{"2,3,0,3,99"}
	want = 2
	got, err = Solvepart1(inputdata)
	if err != nil {
		t.Error(err)
		return
	}

	if got[0] != want {
		t.Errorf("Got %s, wanted %s", strconv.Itoa(got[0]), strconv.Itoa(want))
	}

	if got[3] != 6 {
		t.Errorf("Got %s, wanted %s", strconv.Itoa(got[0]), "3")
	}

}
