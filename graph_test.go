package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	bytes, _ := ioutil.ReadFile("testdata/graph.txt")
	content := string(bytes)
	topNodes := Parse(content)
	dBytes, err := json.Marshal(topNodes)

	fmt.Println(string(dBytes))
	fmt.Println(err)
}
