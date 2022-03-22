/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A note is anywritten document you'd like to jote down, study or review and action to be done later",
	Long:  `A note is anywritten document you'd like to jote down, study or review and action to be done later`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("note called")
	// },
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
