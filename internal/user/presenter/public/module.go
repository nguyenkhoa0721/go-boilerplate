package public

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("user_public_presenter", fx.Invoke(RegisterHandler))
}
