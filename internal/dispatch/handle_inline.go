package dispatch

import (
	"context"

	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/tg"
)

func (b *Bot) OnBotInlineQuery(ctx context.Context, e tg.Entities, u *tg.UpdateBotInlineQuery) error {
	b.logger.Info("Got inline query",
		zap.String("query", u.Query),
		zap.String("offset", u.Offset),
	)

	user, ok := e.Users[u.UserID]
	if !ok {
		return xerrors.Errorf("unknown user ID %d", u.UserID)
	}

	var geo *tg.GeoPoint
	if u.Geo != nil {
		geo, _ = u.Geo.AsNotEmpty()
	}
	if err := b.onInline.OnInline(ctx, InlineQuery{
		QueryID:   u.QueryID,
		Query:     u.Query,
		Offset:    u.Offset,
		Enquirer:  user.AsInput(),
		geo:       geo,
		user:      user,
		baseEvent: b.baseEvent(),
	}); err != nil {
		return xerrors.Errorf("handle inline: %w", err)
	}

	return nil
}
