package admin

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("user_admin_presenter", fx.Invoke(RegisterHandler))
}
