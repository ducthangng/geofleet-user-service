//go:build wireinject

package registry

import (
	"context"

	"github.com/ducthangng/geofleet/user-service/internal/handler"
	"github.com/google/wire"
)

func Initialize(ctx context.Context) (*handler.UserHandler, error) {
	wire.Build(ProviderSet)
	return nil, nil
}
