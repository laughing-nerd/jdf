package utils

import (
	"fmt"
	"regexp"
)

var Colors = map[string][]string{
	"reset":  {"\033[0m", "</span>"},
	"red":    {"\033[31m", "<span&nbspcustom;style='color:&nbspcustom;red;'>%v"},
	"green":  {"\033[32m", "<span&nbspcustom;style='color:&nbspcustom;#6e9565;'>%v"},
	"yellow": {"\033[33m", "<span&nbspcustom;style='color:&nbspcustom;yelow;'>%v"},
	"blue":   {"\033[96m", "<span&nbspcustom;style='color:&nbspcustom;#6ba99e;'>%v"},
	"orange": {"\033[38;5;222m", "<span&nbspcustom;style='color:&nbspcustom;#ffd57a;'>%v"},
	"purple": {"\033[38;5;13m", "<span&nbspcustom;style='color:&nbspcustom;#D7A2FF;'>%v"},
}

func Colorize[T any](x T, color string) string {
	if FlagWebMode {
		return Colors[color][1] + fmt.Sprintf("%v", x) + Colors["reset"][1]
	}

	return Colors[color][0]

}

func RemoveANSIColor(s string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(s, "")
}
