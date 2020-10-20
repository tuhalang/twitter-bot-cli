package rules

import (
	"github.com/spf13/cobra"
	"github.com/tuhalang/twitter-bot/util"
)

var getRuleCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all rules query",
	Long:  "Get all rules query, print {ID VALUE TAG}",
	Run: func(cmd *cobra.Command, args []string) {
		util.GetAllRules()
	},
}
