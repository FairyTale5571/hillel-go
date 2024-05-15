package app

import (
	"github.com/fairytale5571/awesomeProject1/config"
	"github.com/fairytale5571/awesomeProject1/internal/repository"
	"github.com/fairytale5571/awesomeProject1/internal/service"
	"github.com/fairytale5571/awesomeProject1/pkg/yt"
)

type App struct {
	service *service.Service
}

func New() (*App, error) {
	repo, err := repository.NewRepository(
		yt.NewYoutube(),
	)
	if err != nil {
		return nil, err
	}

	instanceService, err := service.New(config.GetConfig().Token, repo)
	if err != nil {
		return nil, err
	}

	return &App{
		service: instanceService,
	}, nil
}

func (a *App) Run() error {
	return a.service.Run()
}

func (a *App) Stop() error {
	return nil
}
