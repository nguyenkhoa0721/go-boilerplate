package main

import (
	"go-boilerplate/cmd/public/bootstrap"
	"go-boilerplate/cmd/public/module"
	authPublicHandler "go-boilerplate/internal/auth/presenter/public"
	userPublicHandler "go-boilerplate/internal/user/presenter/public"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		module.Module(),
		fx.Invoke(authPublicHandler.RegisterHandler),
		fx.Invoke(userPublicHandler.RegisterHandler),
		fx.Invoke(bootstrap.Bootstrap),
	).Run()
}
