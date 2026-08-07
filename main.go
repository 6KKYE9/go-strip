package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 去掉每行首尾空白，可选去掉空行。
func stripLines(lines []string, dropBlank bool) []string {
	var out []string
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if dropBlank && l == "" {
			continue
		}
		out = append(out, l)
	}
	return out
}

func readLines() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	dropBlank := false
	for _, a := range os.Args[1:] {
		switch a {
		case "-b", "--blank":
			dropBlank = true
		case "-h", "--help":
			fmt.Println("go-strip 去掉行首尾空白")
			fmt.Println("用法: 管道 | go run . [-b]")
			fmt.Println("  -b  同时删掉空行")
			return
		}
	}
	for _, l := range stripLines(readLines(), dropBlank) {
		fmt.Println(l)
	}
}
