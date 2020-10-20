package rules

import (
	"github.com/spf13/cobra"
	"github.com/tuhalang/twitter-bot/util"
)

var deleteRuleCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete all rules query",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		util.DeleteAllRules()
	},
}
