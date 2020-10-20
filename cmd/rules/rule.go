package rules

import (
	"github.com/spf13/cobra"
)

var ruleCmd = &cobra.Command{
	Use:   "rules",
	Short: "Add, Get and Delete rules query",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// RegisterRootCmd cmd
func RegisterRootCmd(cmd *cobra.Command) {
	cmd.AddCommand(ruleCmd)
}

func init() {
	addRuleCmd.Flags().StringVarP(&value, "value", "v", "", "value of rule")
	addRuleCmd.Flags().StringVarP(&tag, "tag", "t", "", "tag of rule")
	ruleCmd.AddCommand(addRuleCmd)
	ruleCmd.AddCommand(getRuleCmd)
	ruleCmd.AddCommand(deleteRuleCmd)
}
