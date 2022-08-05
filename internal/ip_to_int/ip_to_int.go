package iptoint

import (
	"encoding/binary"
	"strconv"
	"strings"
)

// IPv4 4 byte
var buf = make([]byte, 4)

// Конвентирование IPv4 адреса в число
//
// example arg s = '145.67.23.4'
func Convert(s string) (uint32, error) {
	// не валидируется что строка должна состоять из 4 полей
	for i, s := range strings.SplitN(s, ".", 4) {
		num, err := strconv.ParseUint(s, 0, 8) // самое большое значение 255
		if err != nil {
			return uint32(0), err
		}
		buf[i] = uint8(num)
	}
	return binary.BigEndian.Uint32(buf), nil // Порядок от старшего к младшему
}
