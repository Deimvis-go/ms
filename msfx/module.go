package msfx

import (
	"github.com/Deimvis-go/ms/ms"
	"github.com/Deimvis-go/ms/ms/msconfig"
	"go.uber.org/fx"
)

// FX Module.
//
// Requires:
// - *msconfig.HTTPSetings
//
// Provides:
// - xhttp.RoundTripWrapFn `name:"self_identity"`
var Module = fx.Module("ms", ModuleOptions...)

var ModuleOptions = []fx.Option{
	fx.Provide(
		fx.Private,
	),
	fx.Provide(
		ms.NewHTTPPathNormalizer,
		fx.Annotate(
			msconfig.NewRoundTripWrap,
			fx.ResultTags(`group:"round_trip_wraps"`),
		),
	),
}
