// example ip'145.67.23.4'
package parser

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"unsafe"

	bitmap "example.com/count_ip/internal/bit"
	readfile "example.com/count_ip/internal/read_file"
)

const sep = byte(46)

var (
	bitArr = bitmap.New()
	buffer [4]byte
)

func b2s(b []byte) string {
	// на основе strings.Builder.String
	return *(*string)(unsafe.Pointer(&b))
}

type iterator struct {
	Buffer []byte
	Token  []byte
}

func (iter *iterator) Next() bool {
	for i := 0; i < len(iter.Buffer); i++ {
		if iter.Buffer[i] == sep {
			iter.Token = iter.Buffer[:i]
			iter.Buffer = iter.Buffer[i+1:]
			return true
		}
	}
	if iter.Buffer != nil {
		iter.Token = iter.Buffer
		iter.Buffer = nil
		return true
	}
	return false
}

var iter = iterator{}

func parseByte(s []byte) error {
	iter.Buffer = s
	for i := 0; i < 4; i++ {
		if iter.Next() == true {
			num, err := strconv.ParseUint(b2s(iter.Token), 0, 8)
			if err != nil {
				return fmt.Errorf("convert ip: %w", err)
			}
			buffer[i] = uint8(num)
		} else {
			return errors.New("bad count fields")
		}
	}

	bitArr.Set(binary.BigEndian.Uint32(buffer[:]))
	return nil
}

func Parser(file string) (int, error) {
	f, closer, err := readfile.New(file)
	if err != nil {
		return 0, fmt.Errorf("open file: %w", err)
	}
	defer closer()

	for f.Scan() {
		err := parseByte(f.Bytes())
		if err != nil {
			return 0, err
		}
	}
	return bitArr.Count(), nil

}
