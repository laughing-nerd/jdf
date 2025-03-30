package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/laughing-nerd/jdf/src"
	"github.com/laughing-nerd/jdf/utils"
)

const BUFFERSIZE = 128 * 1024 // 128 KB

var (
	separatorStr string
	version  string = "v1.3.1"
	resultch        = make(chan string)
)

type winsize struct {
	Row uint16
	Col uint16
}

func init() {
	utils.RegisterFlags()

	ws := &winsize{}
	retCode, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic("Don't worry! This error is from our side. Apologies 😅\n" + err.Error())
	}

	// Generate separatorString
	separatorStr = strings.Repeat(utils.FlagSeparator, int(ws.Col))
}

func main() {
	// Creating a buffered reader and writer of 128KB
	writer := bufio.NewWriterSize(os.Stdout, BUFFERSIZE)
	reader := bufio.NewReaderSize(os.Stdin, BUFFERSIZE)
	defer writer.Flush()

	// greet the user if no data is piped
	if isStdinEmpty() {
		writer.Write([]byte("jdf - JSON Detect and Format 💪\nVersion: "))
		writer.Write([]byte(version))
		writer.Write([]byte("\nFor more info, visit: https://github.com/laughing-nerd/jdf\n"))
		writer.Flush()
		return
	}

	if utils.FlagWebMode {
		go src.StartServer(resultch)
	}

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			// break out of the loop in case of eof with no data
			if err == io.EOF && len(line) == 0 {
				break
			}
		}

		s := utils.RemoveANSIColor(line)

		isJson, start, end := src.DetectJSON(s)
		if !isJson {
			if utils.FlagWebMode {
				resultch <- line
			} else {
				os.Stdout.Write([]byte(line))
			}
			continue
		}

		formattedPrefix := s[:start]
		if len(formattedPrefix) > 0 {
			formattedPrefix += "\n"
		}
		formattedSuffix := s[end+1:]
		if len(formattedSuffix) > 0 {
			formattedSuffix = "\n" + formattedSuffix
		}

		var formatted strings.Builder
		formatted.Grow(len(formattedPrefix) + len(formattedSuffix) + (end - start + 1))
		formatted.WriteString(formattedPrefix)
		formatted.WriteString(src.FormatJSON(s[start : end+1]))
		formatted.WriteString(formattedSuffix)

		if utils.FlagWebMode {
			resultch <- formatted.String()
		} else {
			writer.Write([]byte(separatorStr))
			writer.Write([]byte(formatted.String()))
		}
	}

	// Block indefinitely if web mode is enabled
	if utils.FlagWebMode {
		select {}
	}
}

// helper func ...
func isStdinEmpty() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		os.Stdout.Write([]byte("Error checking stdin: "))
		os.Stdout.Write([]byte(err.Error()))
		os.Stdout.Write([]byte("\n"))
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) != 0
}
