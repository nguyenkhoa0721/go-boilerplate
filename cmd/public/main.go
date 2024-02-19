package main

import (
	"go-boilerplate/cmd/public/bootstrap"
	"go-boilerplate/cmd/public/module"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		module.PkgModule(),
		module.FeatureModule(),
		fx.Invoke(bootstrap.Bootstrap),
	).Run()
}
