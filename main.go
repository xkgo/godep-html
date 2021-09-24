package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func PrintUsage() {
	fmt.Printf("\nUsages:\n\n")
	switch runtime.GOOS {
	case "darwin":
		fmt.Printf("\tgo mod graph | godep-html > graph.html")
	case "linux":
		fmt.Printf("\tgo mod graph | godep-html > graph.html")
	case "windows":
		fmt.Printf("\tgo mod graph | godep-html > graph.html")
	}
	fmt.Printf("\n\n")
}

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("os.Stdin.Stat:", err)
		PrintUsage()
		os.Exit(1)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("command err: command is intended to work with pipes.")
		PrintUsage()
		os.Exit(1)
	}

	// 读取
	bufReader := bufio.NewReader(os.Stdin)
	bytes, err := ioutil.ReadAll(bufReader)
	if nil != err {
		fmt.Println("read err: read go mod graph output err by pipes, err:", err)
		os.Exit(1)
	}

	content := string(bytes)
	topNodes := Parse(content)
	dBytes, err := json.Marshal(topNodes)

	graphJsonData := string(dBytes)
	graphHtml := strings.Replace(GraphHtmlTemplate, "${graphJsonData}", graphJsonData, 1)
	fmt.Println(graphHtml)
}
