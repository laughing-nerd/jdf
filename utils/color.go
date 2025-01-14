package utils

import (
	"fmt"
	"regexp"
)

var Colors = map[string]string{
	"reset":  "\033[0m",
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[96m",
	"orange": "\033[38;5;222m",
}

func Colorize[T any](x T, color string) string {
	return fmt.Sprintf("%s%v%s", Colors[color], x, Colors["reset"])
}

func RemoveANSIColor(s string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(s, "")
}
