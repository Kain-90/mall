package cmd

import (
	"github.com/spf13/cobra"
	"mall/global"
)

var rootCmd = &cobra.Command{}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		global.GVA_LOG.Error("rootCmd.Execute err")
		panic(err)
	}
}
