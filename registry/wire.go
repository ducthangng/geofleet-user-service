//go:build wireinject
// +build wireinject

package registry

import (
	"context"

	"github.com/ducthangng/geofleet/user-service/internal/handler"
	"github.com/google/wire"
)

// InitializeUserService is the "Injector Stub".
// 1. It defines the Function Signature (what inputs we have, what output we want).
// 2. It tells Wire which Providers to use via wire.Build.
func InitializeUserService(ctx context.Context) (*handler.UserRestfulHandler, error) {
	wire.Build(
		ProviderSet, // <--- This pulls in all the logic from providers.go
	)

	// This return statement is a "Stub".
	// It is only here so the Go compiler doesn't complain about missing return types.
	// When Wire runs, it ignores this line and generates the real code in wire_gen.go.
	return &handler.UserRestfulHandler{}, nil
}
