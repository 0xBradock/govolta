/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/textinput"
	"github.com/spf13/cobra"
)

// dashCmd represents the dash command
var dashCmd = &cobra.Command{
	Use:   "dash",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dashboardLayer, err := tcell.New()
		if err != nil {
			fmt.Println(err)
		}
		defer dashboardLayer.Close()
		upperContainer := container.Top(container.Border(linestyle.Light))

		input, err := textinput.New()
		if err != nil {
			panic(err)
		}
		bottomContainer := container.Bottom(container.Border(linestyle.Light), container.PlaceWidget(input))

		containerLayer, err := container.New(
			dashboardLayer, container.SplitHorizontal(
				upperContainer,
				bottomContainer,
				container.SplitPercent(90),
			),
		)
		if err != nil {
			fmt.Println(err)
		}

		ctx, cancel := context.WithCancel(context.Background())
		quitter := func(k *terminalapi.Keyboard) {
			if k.Key == keyboard.KeyEsc {
				cancel()
			}
		}

		if err := termdash.Run(
			ctx,
			dashboardLayer,
			containerLayer,
			termdash.KeyboardSubscriber(quitter)); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dashCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
