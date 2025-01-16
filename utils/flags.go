package utils

import "flag"

var (
	FlagSeparator string
	FlagIndent    int64
	FlagWebMode   bool
	// Strict    bool
)

func RegisterFlags() {

	flag.StringVar(&FlagSeparator, "separator", "=", "Sets the separator")
	flag.StringVar(&FlagSeparator, "s", "=", "Sets the separator (shorthand)")

	flag.Int64Var(&FlagIndent, "indent", 2, "Sets the indentation")
	flag.Int64Var(&FlagIndent, "i", 2, "Sets the indentation (shorthand)")

	flag.BoolVar(&FlagWebMode, "web", false, "Formats logs and outputs in HTML. Logs are updated in real-time")
	// flag.BoolVar(&Strict, "strict", false, "Strict JSON checking")

	flag.Parse()
}
