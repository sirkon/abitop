package bitop

// BitSet включает бит с номером bitno в числе по адресу p
func BitSet(p *uint64, bitno int) bool {
	return bitSet(p, bitno) > 0
}

// BitUnset выключает бит с номером bitno в числе по адресу p
func BitUnset(p *uint64, bitno int) bool {
	return bitUnset(p, bitno) > 0
}

func bitSet(p *uint64, bitno int) uint64
func bitUnset(p *uint64, bitno int) uint64