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
	result, calculate := []float64{4,1,4,0}, varNumber("2*2","2/0","2+2","2-2")
	if !reflect.DeepEqual(result,calculate) {
		t.Errorf("incorrect, should be %v", result)
	}
}


//package main
//import (
//	"testing"
//)
//func TestAddition(t *testing.T) {
//	var additionStructs = [] struct {
//		input [] float64
//		expected float64
//	}{
//		{[]float64{1,2,3,8}, 14},
//		{[]float64{3,5,6},14},
//		{[]float64{2,3,4,3,3.5,7},22.5},
//		{[]float64{2},2},
//
//	}
//	for _, test:= range additionStructs {
//		if output:=Addition(test.input); output != test.expected {
//			t.Error("Test failed", test.input, test.expected, output)
//		}
//	}
//}
//
//func TestSubtract(t *testing.T) {
//	var subtractStruct = [] struct{
//		input [] float64
//		expected  float64
//	} {
//		{[] float64{1,2,3,8}, -12},
//		{[] float64{6,2,-4,-1}, 9},
//		{[] float64{9,0,-7,-90}, 106},
//		{[]float64{2},2},
//	}
//	for _,test:=range subtractStruct{
//		if output:=Subtraction(test.input); output != test.expected {
//			t.Error("Test failed", test.input, test.expected, output)
//		}
//	}
//}
//func TestMultiply(t *testing.T) {
//	var multiplyStruct = [] struct{
//		input [] float64
//		expected  float64
//	} {
//		{[] float64{1,2,3,8}, 48},
//		{[] float64{6,2,-4,-10}, 480},
//		{[] float64{9,0,7,-90}, 0},
//		{[]float64{2},2},
//	}
//	for _,test:=range multiplyStruct{
//		if output:=Multiply(test.input); output != test.expected {
//			t.Error("Test failed", test.input, test.expected, output)
//		}
//	}
//}
//func TestDivide(t *testing.T) {
//	var divideStruct = [] struct{
//		input [] float64
//		expected  float64
//	} {
//		{[] float64{4,2,2,8}, 0.125},
//		{[] float64{6,2,-4,-10}, 0.075},
//		{[] float64{90,10,3}, 3},
//		{[]float64{2},2},
//	}
//	for _,test:=range divideStruct{
//		if output:=Divs(test.input); output != test.expected {
//			t.Error("Test failed", test.input, test.expected, output)
//		}
//	}
//}