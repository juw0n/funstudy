/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/juw0n/funstudy/cmd"
	"github.com/juw0n/funstudy/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
