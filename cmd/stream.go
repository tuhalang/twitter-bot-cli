package cmd

import (
	"github.com/tuhalang/twitter-bot/util"

	"github.com/spf13/cobra"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Get stream data",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		util.GetStreamData()
	},
}
