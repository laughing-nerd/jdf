package utils

import "flag"

var (
	Separator string
	Indent    int64
	// Strict    bool
)

func RegisterFlags() {

	flag.StringVar(&Separator, "separator", "=", "Sets the separator")
	flag.StringVar(&Separator, "s", "=", "Sets the separator (shorthand)")

	flag.Int64Var(&Indent, "indent", 2, "Sets the indentation")
	flag.Int64Var(&Indent, "i", 2, "Sets the indentation (shorthand)")

	// flag.BoolVar(&Strict, "strict", false, "Strict JSON checking")

	flag.Parse()
}
