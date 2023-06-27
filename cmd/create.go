//go:build create

package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new entry",
	Run: func(cmd *cobra.Command, args []string) {
		e := cmd.Flag("entry").Value.String()
		v, _ := cmd.Flags().GetBool("verbose")
		if v {
			fmt.Println("Verbose mode on")
		}
		color.Green("entry is %s", e)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("entry", "e", "", "Entry name for the password")
}
