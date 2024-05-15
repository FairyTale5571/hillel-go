package service

import (
	"context"
	"fmt"
	"github.com/fairytale5571/awesomeProject1/internal/domain"
	"github.com/go-telegram/bot/models"

	"github.com/go-telegram/bot"
)

type Service struct {
	bot    *bot.Bot
	repo   Repository
	states map[int64]domain.State
}

func handler() bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		fmt.Printf("Update: %+v\n", update)
	}
}

func New(token string, repo Repository) (*Service, error) {
	res := &Service{
		states: make(map[int64]domain.State),
		repo:   repo,
	}

	instanceBot, err := bot.New(
		token,
		bot.WithDefaultHandler(handler()),
	)
	if err != nil {
		return nil, err
	}
	res.bot = instanceBot

	return res, nil
}

func (s *Service) Run() error {
	ctx := context.TODO()

	s.bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, s.handlerStart)
	s.bot.RegisterHandler(bot.HandlerTypeMessageText, "📈 Завантажити відео", bot.MatchTypeExact, s.handlerAskLinkYTVideo)
	s.bot.RegisterHandler(bot.HandlerTypeCallbackQueryData, "download_video:", bot.MatchTypePrefix, s.handlerUploadYTVideo)
	s.bot.RegisterHandlerMatchFunc(s.matchState(domain.StateWaitingVideo), s.handlerDownloadYTVideo)

	s.bot.Start(ctx)
	return nil
}

func (s *Service) Stop() error {
	return nil
}
