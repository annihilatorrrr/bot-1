package dispatch

import (
	"context"
	"strings"

	"golang.org/x/xerrors"

	"github.com/gotd/td/tg"
)

// MessageHandler is a simple message event handler.
type MessageHandler interface {
	OnMessage(ctx context.Context, e MessageEvent) error
}

// MessageHandlerFunc is a functional adapter for Handler.
type MessageHandlerFunc func(ctx context.Context, e MessageEvent) error

// OnMessage implements MessageHandler.
func (h MessageHandlerFunc) OnMessage(ctx context.Context, e MessageEvent) error {
	return h(ctx, e)
}

type handle struct {
	MessageHandler
	description string
}

// MessageMux is message event router.
type MessageMux struct {
	prefixes map[string]handle
}

// NewMessageMux creates new MessageMux.
func NewMessageMux() MessageMux {
	return MessageMux{prefixes: map[string]handle{}}
}

// Handle adds given prefix and handler to the mux.
func (m MessageMux) Handle(prefix, description string, handler MessageHandler) {
	m.prefixes[prefix] = handle{
		MessageHandler: handler,
		description:    description,
	}
}

// HandleFunc adds given prefix and handler to the mux.
func (m MessageMux) HandleFunc(prefix, description string, handler func(ctx context.Context, e MessageEvent) error) {
	m.Handle(prefix, description, MessageHandlerFunc(handler))
}

// OnMessage implements MessageHandler.
func (m MessageMux) OnMessage(ctx context.Context, e MessageEvent) error {
	for prefix, handler := range m.prefixes {
		if strings.HasPrefix(e.Message.Message, prefix) {
			if err := handler.OnMessage(ctx, e); err != nil {
				return xerrors.Errorf("handle %q: %w", prefix, err)
			}
			return nil
		}
	}

	return nil
}

// RegisterCommands registers all mux commands using https://core.telegram.org/method/bots.setBotCommands.
func (m MessageMux) RegisterCommands(ctx context.Context, raw *tg.Client) error {
	commands := make([]tg.BotCommand, 0, len(m.prefixes))
	for prefix, handler := range m.prefixes {
		if handler.description == "" {
			continue
		}
		commands = append(commands, tg.BotCommand{
			Command:     strings.TrimPrefix(prefix, "/"),
			Description: handler.description,
		})
	}

	if _, err := raw.BotsSetBotCommands(ctx, commands); err != nil {
		return xerrors.Errorf("set commands: %w", err)
	}
	return nil
}
