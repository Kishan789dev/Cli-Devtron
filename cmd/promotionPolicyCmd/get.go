/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package promotionPolicyCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GetCmd represents the Get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get called")
	},
}

func init() {
	GetCmd.AddCommand(GetAppAndEnvListCmd)
	GetCmd.AddCommand(GetArtifactPromotionPolicyCmd)
	GetCmd.AddCommand(GetListOfPromotionPoliciesCmd)

}
