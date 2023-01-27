/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"awsonbudget/cli/cmd"
	_ "awsonbudget/cli/cmd/pod"
)

func main() {
	cmd.Execute()
}
