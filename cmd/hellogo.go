package cmd

import (
	"os"
	"os/signal"

	"github.com/hellogo/internal/web"
	"github.com/hellogo/pkg/logger"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  "Start the hellogo server app",
	Run: func(cmd *cobra.Command, args []string) {
		downCh := make(chan os.Signal, 1)
		signal.Notify(downCh, os.Interrupt)

		errCh := make(chan error, 1)
		handler := &web.Handler{
			Host: "192.168.1.101",
		}
		handler.Run(errCh)

		select {
		case <-downCh:
			destroy()
			return
		case err := <-errCh:
			logger.Error("The server occur inner error and exit now...", err)
			destroy()
			return
		}
	},
}

func destroy() {
	logger.Info("Stop Server...")
}
