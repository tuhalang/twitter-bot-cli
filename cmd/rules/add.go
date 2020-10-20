package rules

import (
	"github.com/spf13/cobra"
	"github.com/tuhalang/twitter-bot/util"
)

var (
	value string
	tag   string
)

var addRuleCmd = &cobra.Command{
	Use:   "add",
	Short: "Add rule query",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var rules = []map[string]string{
			{"value": value, "tag": tag},
		}
		util.AddRules(rules)
	},
}
