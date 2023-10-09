// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/modernice/goes/aggregate/repository"
	"github.com/modernice/goes/command"
	"github.com/modernice/goes/command/handler"
	"github.com/zsmartex/zsmartex/cmd/user/config"
	"github.com/zsmartex/zsmartex/internal/user/handlers"
	"github.com/zsmartex/zsmartex/internal/user/projection"
	"github.com/zsmartex/zsmartex/pkg/mongodb"
	"github.com/zsmartex/zsmartex/pkg/setup"
)

// Injectors from wire.go:

func InitApp(ctx context.Context, cfg *config.Config) (*App, func(), error) {
	mongoDB := cfg.MongoDB
	client, err := mongodb.NewMongoClient(ctx, mongoDB)
	if err != nil {
		return nil, nil, err
	}
	user := projection.NewUser(ctx, client)
	eventStore := cfg.EventStore
	eventBus := cfg.EventBus
	eventRegistry := setup.NewEventRegistry()
	bus, cleanup, err := setup.NewEventBus(ctx, eventBus, eventRegistry)
	if err != nil {
		return nil, nil, err
	}
	store, err := setup.NewEventStore(ctx, eventStore, bus, eventRegistry)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	repository := setup.NewAggregate(store)
	commandRegistry := setup.NewCommandRegistry()
	commandBus := setup.NewCommandBus(commandRegistry, bus, eventRegistry)
	of := NewCommandHandler(repository, commandBus)
	app := NewApp(user, of, bus, store, eventRegistry, commandRegistry)
	return app, func() {
		cleanup()
	}, nil
}

// wire.go:

func NewCommandHandler(repo *repository.Repository, commandBus command.Bus) *handler.Of[*handlers.User] {
	return handler.New(handlers.NewUser, repo, commandBus)
}