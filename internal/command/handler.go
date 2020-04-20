package command

import (
	"context"
	"github.com/alexey-zayats/claim-handler/internal/server"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

// Handler ...
type Handler struct {
	server *server.Server
}

// HandlerDI ...
type HandlerDI struct {
	dig.In
	Server *server.Server
}

// NewHandler ...
func NewHandler(di HandlerDI) Command {
	h := &Handler{
		server: di.Server,
	}
	return h
}

// Run - имплементация метода Run интерфейса Command
func (cmd *Handler) Run(ctx context.Context, args []string) error {

	if err := cmd.server.Start(ctx); err != nil {
		return errors.Wrap(err, "unable start watch")
	}

	return nil
}
