/*
Copyright © 2025 Sanjay Budhia <snpbudhia@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/sanxbud/index/internal/indexer"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the symlinked archive from --source dir to specified --dest dir",
	Long: `Builds the symlinked archive from --source dir to specified --dest dir. 
	For example: index build --source ~/Pictures --dest ~/PicturesIndex
`,
	Run: func(cmd *cobra.Command, args []string) {
		sourceDir, err := cmd.Flags().GetString("source")
		if (err != nil){
			fmt.Println("Error reading source directory. Please include --source sourcedir for the target directory.")
			return
		}
		destDir, err := cmd.Flags().GetString("dest")
		if (err != nil){
			fmt.Println("Error reading destination directory. Please include --dest destination_dir for the intended index location.")
			return
		}
		fmt.Println("Traversing: ",sourceDir)
		archive, err:= indexer.Run(sourceDir)
		if err != nil{
			fmt.Println("Error traversing: ", err)
			return
		}
		// validating functionality	
		archive.PrintSummary()
		fmt.Println("Building archive in: ",destDir)
		/*if err:= builder.Run(destDir); err!= nil{
			fmt.Println("Error building :",err)
		}*/
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().String("source", "", "Source directory")
	buildCmd.Flags().String("dest", "", "Destination directory to build index")
	buildCmd.MarkFlagsRequiredTogether("source","dest")
  
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
