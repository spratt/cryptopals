package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spratt/cryptopals/common"
	"github.com/spratt/cryptopals/set/1/challenge"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	untrimmedHex, err := reader.ReadString('\n')
	common.PanicIfNotNil(err)

	hexStr := strings.Trim(untrimmedHex, "\n ")
	hexBytes, err := challenge.HexStringToBytes(hexStr)
	common.PanicIfNotNil(err)

	fmt.Println(challenge.BytesToBase64(hexBytes))
}
