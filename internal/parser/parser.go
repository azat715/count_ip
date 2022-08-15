// example ip'145.67.23.4'
package parser

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"unsafe"

	bitmap "example.com/count_ip/internal/bit"
	readfile "example.com/count_ip/internal/read_file"
)

var bitArr = bitmap.New()

func b2s(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

const sep = byte(46)

var pos int = 0

var indx [3]int

var buffer [4]byte

var strArr [4]string

func parseByte(s []byte) error {
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			indx[pos] = i
			pos++
		}
	}

	strArr[0] = b2s(s[:indx[0]])
	strArr[1] = b2s(s[indx[0]+1 : indx[1]])
	strArr[2] = b2s(s[indx[1]+1 : indx[2]])
	strArr[3] = b2s(s[indx[2]+1:])

	for i := 0; i < 4; i++ {
		num, err := strconv.ParseUint(strArr[i], 0, 8)
		if err != nil {
			return fmt.Errorf("convert ip: %w", err)
		}
		buffer[i] = uint8(num)
	}

	bitArr.Set(binary.BigEndian.Uint32(buffer[:]))
	pos = 0
	return nil
}

func Parser(file string) (int, error) {
	f, closer, err := readfile.New(file)
	if err != nil {
		return 0, fmt.Errorf("open file: %w", err)
	}
	defer closer()

	for f.Scan() {
		s := f.Bytes()
		err := parseByte(s)
		if err != nil {
			return 0, err
		}
	}
	return bitArr.Count(), nil

}
