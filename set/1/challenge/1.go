package challenge

import (
	"bytes"
	"encoding/base64"
	"strconv"
)

const (
	HexBase = 16
	ByteBitSize = 8
)

func HexStringToBytes(hex string) ([]byte, error) {
	// If the length of the string is odd, add a leading zero
	if len(hex) % 2 != 0 {
		hex = "0" + hex
	}
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < len(hex) / 2; i++ {
		byteStr := hex[2*i:(2*i)+2]
		curByte, err := strconv.ParseInt(byteStr, HexBase, ByteBitSize)
		if err != nil {
			return nil, err
		}
		buf.Write([]byte{byte(curByte)})
	}
	return buf.Bytes(), nil
}

func BytesToBase64(byts []byte) string {
	return base64.StdEncoding.EncodeToString(byts)
}
