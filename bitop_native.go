// +build !amd64

package bitop

import "sync"

var globalLock sync.Mutex

func bitSet(p *uint64, bitno int) uint64 {
	globalLock.Lock()
	var prev = *p & (1 << bitno)
	*p |= (1 << bitno)
	globalLock.Unlock()
	return prev
}

func bitUnset(p *uint64, bitno int) uint64 {
	globalLock.Lock()
	var prev = *p & (1 << bitno)
	*p &= ^(1 << bitno)
	globalLock.Unlock()
	return prev
}

