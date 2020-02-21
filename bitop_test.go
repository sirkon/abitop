package bitop

import (
	"reflect"
	"testing"
)

func TestBitOp(t *testing.T) {
	data := make([]uint64, 3)

	bitNo := 3

	BitSet(&data[bitNo>>6], bitNo&0x3f)
	sample3 := []uint64{8, 0, 0}
	if !reflect.DeepEqual(data, sample3) {
		t.Fatalf("%v expected, got %v", sample3, data)
	}

	bitNo = 65
	BitSet(&data[bitNo>>6], bitNo&0x3f)
	sample39 := []uint64{8, 2, 0}
	if !reflect.DeepEqual(data, sample39) {
		t.Fatalf("%v expected, got %v", sample39, data)
	}

	bitNo = 128
	BitSet(&data[bitNo>>6], bitNo&0x3f)
	sample3916 := []uint64{8, 2, 1}
	if !reflect.DeepEqual(data, sample3916) {
		t.Fatalf("%v expected, got %v", sample3916, data)
	}

	bitNo = 3
	BitUnset(&data[bitNo>>6], bitNo&0x3f)
	sample916 := []uint64{0, 2, 1}
	if !reflect.DeepEqual(data, sample916) {
		t.Fatalf("%v expected, got %v", sample916, data)
	}

	bitNo = 65
	BitUnset(&data[bitNo>>6], bitNo&0x3f)
	sample16 := []uint64{0, 0, 1}
	if !reflect.DeepEqual(data, sample16) {
		t.Fatalf("%v expected, got %v", sample16, data)
	}
}
