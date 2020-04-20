package cmd

import (
	"context"
	"github.com/alexey-zayats/claim-handler/internal/command"
	"github.com/alexey-zayats/claim-handler/internal/config"
	"github.com/alexey-zayats/claim-handler/internal/di"
	"github.com/alexey-zayats/claim-handler/internal/queue"
	"github.com/alexey-zayats/claim-handler/internal/server"
	"github.com/spf13/cobra"
)

var handleCmd = &cobra.Command{
	Use:   "handle",
	Short: "handle",
	Long:  "handle",
	Run:   handleMain,
}

func init() {
	rootCmd.AddCommand(handleCmd)
}

func handleMain(cmd *cobra.Command, args []string) {

	ctx := context.Background()

	di := &di.Runner{
		Provide: map[string]interface{}{
			"config":           config.NewConfig,
			"queue.Connection": queue.NewConnection,
			"queue.Queue":      queue.NewQueue,
			"server.Sever":     server.NewServer,
			"command.Handler":  command.NewHandler,
		},
		Invoke: func(ctx context.Context, args []string) interface{} {
			return func(i command.Command) error {
				return i.Run(ctx, args)
			}
		},
	}

	di.Run(ctx, args)
}
