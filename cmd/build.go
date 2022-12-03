/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
		f, _ := cmd.Flags().GetString("file")
		c := compiler.NewCompiler(f)
		b := c.Compile()

		newpath := f[:strings.LastIndex(f, ".")]
		fmt.Println(b, "-", newpath)
		err := ioutil.WriteFile(newpath, b, 0755)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringP("file", "f", "", "Build TD4 assembler file.")
	buildCmd.MarkFlagRequired("file")
}
