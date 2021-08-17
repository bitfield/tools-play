package play

import (
	"strings"
)

func Publish(args []string) string {
	return strings.Join(args, " ")
}