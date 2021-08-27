package main

import (
	"fmt"
	"github.com/ismdeep/args"
	"github.com/ismdeep/gauth"
)

const helpMsg = `Usage: gauth --gen-key | --gen-code -key <KEY> | --help`

func main() {
	if args.Exists("--help") {
		fmt.Println(helpMsg)
		return
	}

	if args.Exists("--gen-key") {
		fmt.Println(gauth.CreateSecret(16))
		return
	}

	if args.Exists("--gen-code") && args.Exists("-key") {
		fmt.Println(gauth.GetCode(args.GetValue("-key")))
		return
	}

	fmt.Println(helpMsg)
	return
}
