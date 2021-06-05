package router

import "github.com/mattn/go-shellwords"

type Parser func(string) ([]string, error)

func NewDefaultParser() Parser {
	return shellwords.Parse
}
