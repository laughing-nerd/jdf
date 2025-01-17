package utils

import (
	"fmt"
	"regexp"
)

var Colors = map[string][]string{
	"reset":  {"\033[0m", ""},
	"red":    {"\033[31m", "<span&nbspcustom;style='color:&nbspcustom;red;'>%v</span>"},
	"green":  {"\033[32m", "<span&nbspcustom;style='color:&nbspcustom;#6e9565;'>%v</span>"},
	"yellow": {"\033[33m", "<span&nbspcustom;style='color:&nbspcustom;yelow;'>%v</span>"},
	"blue":   {"\033[96m", "<span&nbspcustom;style='color:&nbspcustom;#6ba99e;'>%v</span>"},
	"orange": {"\033[38;5;222m", "<span&nbspcustom;style='color:&nbspcustom;#ffd57a;'>%v</span>"},
	"purple":   {"\033[38;5;13m", "<span&nbspcustom;style='color:&nbspcustom;#D7A2FF;'>%v</span>"},
}

func Colorize[T any](x T, color string) string {
	if FlagWebMode {
		return fmt.Sprintf(Colors[color][1], x)
	}
	return fmt.Sprintf("%s%v%s", Colors[color][0], x, Colors["reset"][0])
}

func RemoveANSIColor(s string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(s, "")
}
