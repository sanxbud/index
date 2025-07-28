/*
Copyright © 2025 Sanjay Budhia <snpbudhia@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/sanxbud/index/internal/cleaner"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Prunes the targeted directory and subdirectories of broken symlinks",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		targDir, err := cmd.Flags().GetString("target")
		if (err != nil){
			fmt.Println("Error reading target directory. Please include --source sourcedir for the target directory.")
			return
		}
		count, err:= cleaner.Run(targDir)
		if err != nil{
			return 	
		}
	  fmt.Printf("Clean complete. %d broken symlinks removed from %s\n",count,targDir)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().String("target", "", "Target directory")
	cleanCmd.MarkFlagsOneRequired("target")
  
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
