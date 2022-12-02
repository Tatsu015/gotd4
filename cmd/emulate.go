/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io/ioutil"

	"github.com/Tatsu015/gotd4.git/internal/emulator"
	"github.com/spf13/cobra"
)

// emulateCmd represents the emulate command
var emulateCmd = &cobra.Command{
	Use:   "emulate",
	Short: "Emurates TD4 CPU",
	Long:  "Emurates TD4 CPU. A ROM file is required to run the emulator.",
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := cmd.Flags().GetString("file")
		bytes, err := ioutil.ReadFile(f)
		if err != nil {
			panic(err)
		}
		rom := emulator.NewROM(bytes)
		in := emulator.NewInput()
		out := emulator.NewOutput()
		e := emulator.NewEmulator(&rom, &in, &out)
		e.Run()
	},
}

func init() {
	rootCmd.AddCommand(emulateCmd)
	emulateCmd.Flags().StringP("file", "f", "", "Emurate TD4 CPU using specified 'file' program")
	emulateCmd.MarkFlagRequired("file")
}
