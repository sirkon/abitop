// +build !amd64

package bitop

import "sync"

var globalLock sync.Mutex

func BitSet(p *uint64, bitno int)  {
	globalLock.Lock()
	*p |= (1 << bitno)
	globalLock.Unlock()
}

func BitUnset(p *uint64, bitno int) {
	globalLock.Lock()
	*p &= ^(1 << bitno)
	globalLock.Unlock()
}

