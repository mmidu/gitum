package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetData gets input data
func GetData(data string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s: ", data)

	ret, _ := reader.ReadString('\n')

	return strings.Replace(ret, "\n", "", -1)
}
