package main

import (
	"context"
	"net/http/httputil"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
	"golang.org/x/xerrors"

	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/message/styling"
	"github.com/gotd/td/tg"
)

const (
	githubOwner    = "gotd"
	githubRepo     = "td"
	githubRef      = "main"
	githubWorkflow = "bot.yml"
)

func (b *Bot) answerGH(
	ctx context.Context,
	send *message.Builder,
	peer tg.InputPeerClass,
	m *tg.Message,
) error {
	return b.getReply(ctx, send, peer, m, func(msg *tg.Message) error {
		gh := b.github
		if gh == nil {
			if _, err := send.Text(ctx, "Github intergration disabled"); err != nil {
				return xerrors.Errorf("send: %w", err)
			}

			return nil
		}

		// Create client with short-lived repository installation token.
		//
		inst, _, err := gh.Apps.FindRepositoryInstallation(ctx, githubOwner, githubRepo)
		if err != nil {
			return xerrors.Errorf("find repository installation: %w", err)
		}
		tok, _, err := gh.Apps.CreateInstallationToken(ctx, inst.GetID(), nil)
		if err != nil {
			return xerrors.Errorf("create installation token: %w", err)
		}
		gh = github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
			&oauth2.Token{
				AccessToken: tok.GetToken(),
			},
		)))

		// Reply with response header.
		//
		resp, err := gh.Actions.CreateWorkflowDispatchEventByFileName(
			ctx, githubOwner, githubRepo, githubWorkflow,
			github.CreateWorkflowDispatchEventRequest{
				Ref: githubRef,
				Inputs: map[string]interface{}{
					"telegram": true,
					"message":  m,
					"replyto":  msg,
				},
			},
		)
		if err != nil {
			return xerrors.Errorf("dispatch workflow: %w", err)
		}
		data, err := httputil.DumpResponse(resp.Response, true)
		if err != nil {
			return xerrors.Errorf("dump response: %w", err)
		}
		if _, err := send.StyledText(ctx, styling.Pre(string(data), "")); err != nil {
			return xerrors.Errorf("send: %w", err)
		}
		return nil
	})
}