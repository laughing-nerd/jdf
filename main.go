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

		start, jsonStr := getJSON(line)
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
