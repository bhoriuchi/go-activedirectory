package common

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
)

// Hex2sid converts hex to an sid
func Hex2sid(h string) (sid string, err error) {
	var revision int64
	var dashes int64
	var authority int64
	var subAuthority int
	sid = "S"

	pos := 0
	if revision, err = Hexdec(Substr(h, pos, 2)); err != nil {
		return
	}

	sid = fmt.Sprintf("%s-%d", sid, revision)
	pos += 2

	if dashes, err = Hexdec(Substr(h, pos, 2)); err != nil {
		return
	}

	pos += 2

	if authority, err = Hexdec(Substr(h, pos, 12)); err != nil {
		return
	}

	sid = fmt.Sprintf("%s-%d", sid, authority)
	pos += 12

	if subAuthority, err = HexULong32Le2int(Substr(h, pos, 8)); err != nil {
		return
	}

	sid = fmt.Sprintf("%s-%d", sid, subAuthority)
	for i := 0; i < int(dashes)-1; i++ {
		var part int
		pos += 8
		if part, err = HexULong32Le2int(Substr(h, pos, 8)); err != nil {
			return
		}
		sid = fmt.Sprintf("%s-%d", sid, part)
	}

	return
}

// Substr gets substring
func Substr(str string, start, length int) string {
	return str[start : start+length]
}

// Hexdec hex to decimal
func Hexdec(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 0)
}

// HexULong32Le2int hex to long 32 little-endian int
func HexULong32Le2int(h string) (int, error) {
	var i uint32
	decoded, err := hex.DecodeString(h)
	if err != nil {
		return 0, err
	}

	if err := binary.Read(bytes.NewReader(decoded), binary.LittleEndian, &i); err != nil {
		return 0, err
	}

	return int(i), nil
}

// HexUShort16Le2int hex to short 16 little-endian int
func HexUShort16Le2int(h string) (int, error) {
	var i uint16
	decoded, err := hex.DecodeString(h)
	if err != nil {
		return 0, err
	}

	if err := binary.Read(bytes.NewReader(decoded), binary.LittleEndian, &i); err != nil {
		return 0, err
	}

	return int(i), nil
}
