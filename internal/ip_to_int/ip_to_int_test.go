package iptoint

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertIP(t *testing.T) {
	tests := []struct {
		input string
		want  uint32
	}{
		{"192.168.0.1", 3232235521},
		{"0.0.0.0", 0},
		{"255.255.255.255", 4294967295},
		{"145.67.23.4", 2437093124},
		{"8.34.5.23", 136447255},
		{"89.54.3.124", 1496712060},
		{"3.45.71.5", 53298949},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test=%d", i),
			func(t *testing.T) {
				assert := assert.New(t)
				i, err := Convert(tc.input)
				if assert.NoError(err) {
					assert.Equal(i, tc.want)
				}
			})
	}
	t.Run("Fail value out of range", func(t *testing.T) {
		assert := assert.New(t)
		i, err := Convert("256.255.255.255")
		assert.Equal(i, uint32(0))
		assert.EqualError(err, "strconv.ParseUint: parsing \"256\": value out of range")
	})
	t.Run("Fail invalid syntax", func(t *testing.T) {
		assert := assert.New(t)
		i, err := Convert("192.168.0.-1")
		assert.Equal(i, uint32(0))
		assert.EqualError(err, "strconv.ParseUint: parsing \"-1\": invalid syntax")
	})
}
