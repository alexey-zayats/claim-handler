package application

import (
	"bytes"
	"strconv"
	"strings"
)

// ValidationErrors ...
type ValidationErrors map[string][]string

func (ve ValidationErrors) Error() string {

	buff := bytes.NewBufferString("")

	for key, s := range ve {
		buff.WriteString(key)
		buff.WriteString(": ")
		buff.WriteString(strings.Join(s, ", "))
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

func parseInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		n = 0
	}
	return n
}

// Validate ...
