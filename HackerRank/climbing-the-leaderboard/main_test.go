package main

import (
	"testing"
)

func arrEqual(arr1, arr2 []int32) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func TestRmDup(t *testing.T) {
	ar1 := rmDup([]int32{3, 3, 3, 2, 1, 1})
	ar1Correct := []int32{3, 2, 1}
	if !arrEqual(ar1, ar1Correct) {
		t.Fatal()
	}

	ar2 := rmDup([]int32{3, 3, 3, 3, 3, 3, 3, 3, 3})
	ar2Correct := []int32{3}
	if !arrEqual(ar2, ar2Correct) {
		t.Fatal()
	}

	ar3 := rmDup([]int32{3})
	ar3Correct := []int32{3}
	if !arrEqual(ar3, ar3Correct) {
		t.Fatal()
	}
}

func TestBinaryInsert(t *testing.T) {
	arr := []int32{100, 50, 40, 20, 10}
	if binaryInsert(arr, 5) != 5 {
		t.Fatal()
	}
	if binaryInsert(arr, 25) != 3 {
		t.Fatal()
	}
	if binaryInsert(arr, 50) != 1 {
		t.Fatal()
	}
	if binaryInsert(arr, 120) != 0 {
		t.Fatal()
	}
}
