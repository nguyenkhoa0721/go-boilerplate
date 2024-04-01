package main

import (
	"go-boilerplate/cmd/public/bootstrap"
	"go-boilerplate/fxmodule"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fxmodule.PkgModule(),
		fxmodule.FeatureModule(),
		fx.Invoke(bootstrap.Bootstrap),
	).Run()
}
