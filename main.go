package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/laughing-nerd/jdf/src"
	"github.com/laughing-nerd/jdf/utils"
)

var (
	count    int
	resultch        = make(chan string)
	Version  string = "v1.2.0"
)

type winsize struct {
	Row uint16
	Col uint16
}

func init() {
	ws := &winsize{}
	retCode, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic("Don't worry! This error is from our side. Apologies 😅\n" + err.Error())
	}
	count = int(ws.Col)
}

func main() {
	utils.RegisterFlags()

	// greet the user if no data is piped
	if isStdinEmpty() {
		fmt.Println("jdf - JSON Detect and Format 💪\nVersion: " + Version + "\nFor more info, visit: https://github.com/laughing-nerd/jdf")
		return
	}

	if utils.FlagWebMode {
		go src.StartServer(resultch)
	}

	// Generate separator string
	var separatorStr string
	for range count {
		separatorStr += utils.FlagSeparator
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		// s := utils.RemoveANSIColor(line) // Remove ansi color code if present

		isJson, start, end := src.DetectJSON(line)
		if !isJson {
			if utils.FlagWebMode {
				resultch <- line
			} else {
				fmt.Println(line)
			}
			continue
		}

		formatted := line[:start] + "\n" + src.FormatJSON(line[start:end+1]) + "\n" + line[end+1:]

		if utils.FlagWebMode {
			resultch <- formatted
		} else {
			fmt.Printf("%s\n%s\n", separatorStr, formatted) // Display the formatted json
		}

		if err := scanner.Err(); err != nil {
			panic("Don't worry! This error is from our side. Apologies 😅\n" + err.Error())
		}
	}

	// Block indefinitely if web mode is enabled
	if utils.FlagWebMode {
		select {}
	}
}

func isStdinEmpty() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error checking stdin:", err)
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) != 0
}
