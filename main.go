package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var (
	Colors = map[string]string{
		"reset":  "\033[0m",
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[96m",
		"orange": "\033[38;5;222m",
	}
	count int

	// Flags
	separator *string = flag.String("s", "=", "Sets the separator")
)

type winsize struct {
	Row uint16
	Col uint16
}

func init() {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	count = int(ws.Col)
}

func main() {
	flag.Parse()

	var separatorStr string
	for range count {
		separatorStr += *separator
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line := scanner.Text()

		s := removeANSIColors(line)

		start, jsonStr := getJSON(s)
		if start == -1 {
			fmt.Println(line)
			continue
		}

		formattedJson, jsonErr := getFormattedJSON(jsonStr)
		if jsonErr != nil {
			panic(jsonErr.Error())
		}

		fmt.Printf("%s\n%s\n", separatorStr, formattedJson)

		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
	}
}
