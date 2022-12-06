/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io/ioutil"
	"time"

	"github.com/Tatsu015/gotd4.git/internal/emulator"
	"github.com/Tatsu015/gotd4.git/internal/emulator/io"
	"github.com/Tatsu015/gotd4.git/internal/emulator/rom"
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
		clock, _ := cmd.Flags().GetInt("clock")
		debug, _ := cmd.Flags().GetBool("debug")

		rom := rom.NewROM(bytes)
		in := io.NewInput()
		out := io.NewOutput()
		e := emulator.NewEmulator(&rom, &in, &out, time.Duration(clock), debug)
		e.Run()
	},
}

func init() {
	rootCmd.AddCommand(emulateCmd)

	emulateCmd.Flags().StringP("file", "f", "", "Emurate TD4 CPU using specified 'file' program")
	emulateCmd.MarkFlagRequired("file")
	emulateCmd.Flags().IntP("clock", "c", 100, "Emurator clock ms")
	emulateCmd.Flags().BoolP("debug", "d", false, "Emurate TD4 CPU with debugging")
}
