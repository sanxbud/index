/*
Copyright © 2025 Sanjay Budhia

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "index",
	Short: "A symlink-based virtual archiver for media files",
	Long: `
 Index - A symlink-based virtual archiver for media files.

 Pass in a source directory and intended destination and, after a scan, an index will be created.
 Each file in the index is a symlink, meaning you can instantly have an organized copy of your files 
 with minimal additional disk usage, and without disturbing your original folder setup`,

}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.CompletionOptions.DisableDefaultCmd = true

}


