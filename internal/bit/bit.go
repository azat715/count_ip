package bit

const size uint64 = 4294967296

type bitmap [size]bool

func (b *bitmap) Set(idx uint32) {
	b[idx] = true
}

func (b *bitmap) Count() int {
	count := 0
	for i := 0; i < len(b); i++ {
		if b[i] == true {
			count++
		}
	}
	return count
}

func New() *bitmap {
	return &bitmap{}
}
