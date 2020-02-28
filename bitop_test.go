package bitop

import (
	"reflect"
	"sync"
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
	res := BitUnset(&data[bitNo>>6], bitNo&0x3f)
	if !res {
		t.Fatal("bit no 65 expected to be set before the call")
	}
	sample16 := []uint64{0, 0, 1}
	if !reflect.DeepEqual(data, sample16) {
		t.Fatalf("%v expected, got %v", sample16, data)
	}

	// check if bit no 65 is off setting it again
	res = BitSet(&data[bitNo>>6], bitNo&0x3f)
	if res {
		t.Fatal("bit no 65 expected to be set off before the call")
	}
	if !reflect.DeepEqual(data, sample916) {
		t.Fatalf("%v expected, got %v", sample916, data)
	}
}

const iterations = 1048576 // 2^20

func BenchmarkAMD64(b *testing.B) {
	data := make([]uint64, 4096)
	worker := func(step int, wg *sync.WaitGroup) {
		start := 0
		for i := 0; i < iterations; i++ {
			if i&1 > 0 {
				BitSet(&data[start>>6], start&0x3f)
			} else {
				BitUnset(&data[start>>6], start&0x3f)
			}
			start += step
			start %= 4096 * 64
		}
		wg.Done()
	}
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(4)
		go worker(17, &wg)
		go worker(371, &wg)
		go worker(512, &wg)
		go worker(1023, &wg)
		wg.Wait()
	}
}

// ---- копипаста из bitop_native.go с небольшим изменением ----
var pureGoGlobalLock sync.Mutex

func pureGoBitSet(p *uint64, bitno int) bool {
	pureGoGlobalLock.Lock()
	var prev = *p & (1 << bitno)
	*p |= (1 << bitno)
	pureGoGlobalLock.Unlock()
	return prev != 0
}

func pureGoBitUnset(p *uint64, bitno int) bool {
	pureGoGlobalLock.Lock()
	var prev = *p & (1 << bitno)
	*p &= ^(1 << bitno)
	pureGoGlobalLock.Unlock()
	return prev != 0
}

func BenchmarkPureGo(b *testing.B) {
	data := make([]uint64, 4096)
	worker := func(step int, wg *sync.WaitGroup) {
		start := 0
		for i := 0; i < iterations; i++ {
			if i&1 > 0 {
				pureGoBitSet(&data[start>>6], start&0x3f)
			} else {
				pureGoBitUnset(&data[start>>6], start&0x3f)
			}
			start += step
			start %= 4096 * 64
		}
		wg.Done()
	}
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(4)
		go worker(17, &wg)
		go worker(371, &wg)
		go worker(512, &wg)
		go worker(1023, &wg)
		wg.Wait()
	}
}

func calc1(i int, j int) int {
	return i + i&j
}

func calc2(i, j int) int {
	return i + i | j
}

func BenchmarkOverheadEstimation(b *testing.B) {
	var sum int
	simpleJob := func(step int, wg *sync.WaitGroup) {
		var start int
		for i := 0; i < iterations; i++ {
			if i&1 > 0 {
				sum += calc1(start >> 6, start & 0x3f)
			} else {
				sum += calc2(start >> 6, start & 0x3f)
			}
			start += step
		}
		wg.Done()
	}
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(4)
		go simpleJob(17, &wg)
		go simpleJob(371, &wg)
		go simpleJob(512, &wg)
		go simpleJob(1023, &wg)
		wg.Wait()
	}
}

/*
Результаты
BenchmarkAMD64
BenchmarkAMD64-4                	      90	  13340163 ns/op
BenchmarkPureGo
BenchmarkPureGo-4               	       3	 438659858 ns/op
BenchmarkOverheadEstimation
BenchmarkOverheadEstimation-4   	     475	   2426861 ns/op

Разница: (438659858 - 2426861) / (12367631 - 2426861) ≈ 33.47
*/
