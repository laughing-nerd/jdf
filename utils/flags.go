package utils

import "flag"

var (
	FlagSeparator string
	FlagIndent    int64
	FlagWebMode   bool
	FlagPort      int64
	FlagNoEscape  bool
	// Strict    bool
)

func RegisterFlags() {

	flag.StringVar(&FlagSeparator, "separator", "=", "Sets the separator")
	flag.StringVar(&FlagSeparator, "s", "=", "Sets the separator (shorthand)")

	flag.Int64Var(&FlagIndent, "indent", 2, "Sets the indentation")
	flag.Int64Var(&FlagIndent, "i", 2, "Sets the indentation (shorthand)")

	flag.Int64Var(&FlagPort, "port", 6969, "Sets the port")
	flag.Int64Var(&FlagPort, "p", 6969, "Sets the port (shorthand)")

	flag.BoolVar(&FlagWebMode, "web", false, "Formats logs and outputs in HTML. Logs are updated in real-time")
	// flag.BoolVar(&Strict, "strict", false, "Strict JSON checking")

	flag.BoolVar(&FlagNoEscape, "no-escape", false, "Do not escape special characters")
	flag.BoolVar(&FlagNoEscape, "ne", false, "Do not escape special characters (shorthand)")

	flag.Parse()
}
