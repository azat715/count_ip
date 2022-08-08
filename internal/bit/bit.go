/*
2^32 = 4294967296
64 битная машина
4294967296 / 64 = 67108864
*/

package bit

const step uint32 = 64 // bit в одной ячейке bitmap

type bitmap [67108864]uint64

func (b *bitmap) Set(idx uint32) {
	// word индекс массива bitmap
	// bit позиция в ячейке массива bitmap
	word, bit := idx/step, idx%step

	// b[word] = b[word] | 1 << bit
	b[word] |= 1 << bit
}

func (b *bitmap) Test(idx uint32) bool {
	word, bit := idx/step, idx%step
	return (b[word])&(1<<bit) != 0
}

func (b *bitmap) Count() int {
	count := 0
	for _, word := range b {
		for i := 0; i < int(step); i++ {
			if word&(1<<i) != 0 {
				count++
			}
		}
	}
	return count
}

func New() *bitmap {
	return &bitmap{}
}
