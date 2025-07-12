package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/laughing-nerd/jdf/src"
	"github.com/laughing-nerd/jdf/utils"
	"github.com/laughing-nerd/jdf/utils/term"
)

var (
	separatorStr string
	version      string = "v0.0.0"
	resultch            = make(chan string)
)

type winsize struct {
	Row uint16
	Col uint16
}

func init() {
	utils.RegisterFlags()
	separatorStr = term.InitSeparator(utils.FlagSeparator)
}

func main() {

	// greet the user if no data is piped
	if isStdinEmpty() {
		os.Stdout.Write([]byte("jdf - JSON Detect and Format 💪\nVersion: "))
		os.Stdout.Write([]byte(version))
		os.Stdout.Write([]byte("\nFor more info, visit: https://github.com/laughing-nerd/jdf\n"))
		return
	}

	if utils.FlagWebMode {
		go src.StartServer(resultch)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		s := utils.RemoveANSIColor(line)

		isJson, start, end := src.DetectJSON(s)
		if !isJson {
			if utils.FlagWebMode {
				resultch <- line
			} else {
				os.Stdout.Write([]byte(line))
				os.Stdout.Write([]byte("\n"))
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
			os.Stdout.Write([]byte(separatorStr))
			os.Stdout.Write([]byte(formatted.String()))
			os.Stdout.Write([]byte("\n"))
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
