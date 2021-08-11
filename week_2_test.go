package main

import (
	"reflect"
	"testing"
)

func TestSub(t *testing.T){
	newSubs := []float64{2,2}
	result, subs := float64(0), Subtraction(newSubs)
	if result != subs {
		t.Errorf("incorrect, should be %v", result)
	}
}

func TestAdds(t *testing.T){
	newAdd := []float64{2,2}
	result, adds := float64(4), Addition(newAdd)
	if result != adds {
		t.Errorf("incorrect, it should be %v", result)
	}
}

func TestDivs(t *testing.T){
	newDivs := []float64{2,2}
	result, divs := float64(1), Divs(newDivs)
	if result != divs {
		t.Errorf("incorrect, it should be %v", result)
	}
}
func TestMultiply(t *testing.T){
	newMult := []float64{2,2}
	result, multi := float64(4), Multiply(newMult)
	if result != multi {
		t.Errorf("incorrect, it should be %v", result)
	}
}
func TestVar(t *testing.T){
	result, calculate := []float64{4,1,4,0}, varNumber("2*2","2/2","2+2","2-2")
	if !reflect.DeepEqual(result,calculate) {
		t.Errorf("incorrect, should be %v", result)
	}
}
