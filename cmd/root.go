/*
Copyright © 2024 Hashim Colombowala
*/
package cmd

import (
	"fmt"
	"os"

	"path/filepath"

	"github.com/hash167/go-head/internal"
	"github.com/spf13/cobra"
)

var filename string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-head",
	Short: "Display the first few lines of a file in Go",
	Long:  `go-head is a CLI tool that displays the first few lines of a file in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		n, _ := cmd.Flags().GetInt("n")
		isByte, _ := cmd.Flags().GetBool("byte")
		for _, arg := range args {
			filenameWithExt := filepath.Base(arg)
			fmt.Println("==>", filenameWithExt, "<==")
			fmt.Println(internal.ReadFile(arg, n, isByte))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-head.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&filename, "file", "f", "", "Input file name")
	rootCmd.Flags().IntP("n", "n", 0, "Number of lines or bytes to read")
	rootCmd.Flags().BoolP("byte", "b", false, "Whether to read lines or bytes")
}
