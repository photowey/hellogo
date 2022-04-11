package cmd

import (
	"fmt"
	"os"

	"github.com/hellogo/app"
	"github.com/spf13/cobra"
)

var (
	conf string

	cmd = &cobra.Command{
		Use:   "hellogo",
		Short: "hellogo",
		Long:  "hellogo",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func init() {
	cobra.OnInitialize(start)
	cmd.PersistentFlags().StringVarP(&conf, "conf", "f", "", "配置文件路径")
	cmd.AddCommand(serverCmd)
}

func start() {
	if err := app.Run(conf); err != nil {
		cobra.CheckErr(err)
	}
}

func Run() {
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
