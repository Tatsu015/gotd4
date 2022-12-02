/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io/ioutil"

	"github.com/Tatsu015/gotd4.git/internal/app"
	"github.com/spf13/cobra"
)

// emulateCmd represents the emulate command
var emulateCmd = &cobra.Command{
	Use:   "emulate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := cmd.Flags().GetString("file")
		if f != "" {
			bytes, err := ioutil.ReadFile(f)
			if err != nil {
				panic(err)
			}
			rom := app.NewROM(bytes)
			in := app.NewIO()
			out := app.NewIO()
			cpu := app.NewCPU(rom, in, out)
			cpu.Run()
		}
	},
}

func init() {
	rootCmd.AddCommand(emulateCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// emulateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// emulateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	emulateCmd.Flags().StringP("file", "f", "", "emurate TD4 CPU using specified 'file' program")
}
