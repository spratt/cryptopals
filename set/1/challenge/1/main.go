package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	HexBase = 16
	ByteBitSize = 8
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func HexStringToBytes(hex string) []byte {
	// If the length of the string is odd, add a leading zero
	if len(hex) % 2 != 0 {
		hex = "0" + hex
	}
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < len(hex) / 2; i++ {
		byteStr := hex[2*i:(2*i)+2]
		curByte, err := strconv.ParseInt(byteStr, HexBase, ByteBitSize)
		check(err)
		buf.Write([]byte{byte(curByte)})
	}
	return buf.Bytes()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
    untrimmedHex, err := reader.ReadString('\n')
	check(err)
	hexStr := strings.Trim(untrimmedHex, "\n ")
	fmt.Println(hexStr)

	hexBytes := HexStringToBytes(hexStr)
	fmt.Println(base64.StdEncoding.EncodeToString(hexBytes))
}
