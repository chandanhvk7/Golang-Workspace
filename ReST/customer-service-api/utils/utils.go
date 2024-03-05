package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CheckForError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Input(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	val, err := reader.ReadString('\n')
	CheckForError(err)
	return strings.TrimSpace(val)
}
