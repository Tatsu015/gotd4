/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io/ioutil"
	"strings"

	"github.com/Tatsu015/gotd4.git/internal/compiler"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build TD4 assembler file to binary",
	Long:  "Build TD4 assembler file to binary. assembler file extension is 'td4asm'.",
	Run: func(cmd *cobra.Command, args []string) {
		srcPath, _ := cmd.Flags().GetString("file")
		newPath, _ := cmd.Flags().GetString("out")
		if newPath == "" {
			newPath = srcPath[:strings.LastIndex(srcPath, ".")]
		}

		c := compiler.NewCompiler(srcPath)
		b := c.Compile()

		err := ioutil.WriteFile(newPath, b, 0755)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringP("file", "f", "", "Build TD4 assembler file.")
	buildCmd.MarkFlagRequired("file")
	buildCmd.Flags().StringP("out", "o", "", "Build TD4 binary file output path.")
}
