package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitmap(t *testing.T) {
	var b = New()
	b.Set(3232235521)
	b.Set(4294967295)
	b.Set(4294967295)
	b.Set(0)
	b.Set(136447255)
	assert := assert.New(t)
	assert.Equal(b.Count(), 4)
}

func BenchmarkBitmap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b = &bitmap{}
		b.Set(3232235521)
		b.Set(4294967295)
		b.Set(4294967295)
		b.Set(0)
		b.Set(136447255)
		b.Count()
	}

}

var bitArr = New()

func BenchmarkBitset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitArr.Set(3232235521)
		bitArr.Count()
	}

}
