package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func validate(input string) (bool, error) {
	if input == "" {
		return false, errors.New("cannot be empty")
	}
	m := regexp.MustCompile("[a-zA-Z]")
	if m.MatchString(input) == false {
		return false, errors.New("please input strings")
	}
	return true, nil
}

type text struct {
	input string
}

func (t text) deleteNumber(s string) string {
	m := regexp.MustCompile("[0-9]")
	res := m.ReplaceAllString(s, "")
	return res
}
func (t text) reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	result = strings.ToLower(result)
	result = strings.Title(result)
	return result
}

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("Input some string :  ")
	input, _ := consoleReader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	var console = text{input}
	if valid, err := validate(input); valid {
		fmt.Println(console.reverse(console.deleteNumber(console.input)))
	} else {
		fmt.Println(err.Error())
	}
}
