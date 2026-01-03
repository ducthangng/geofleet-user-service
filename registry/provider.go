package registry

import (
	"context"
	"errors"

	"github.com/ducthangng/geofleet/user-service/internal/handler"
	"github.com/ducthangng/geofleet/user-service/internal/interface/postgresql"
	usecase "github.com/ducthangng/geofleet/user-service/internal/usercase"
	"github.com/ducthangng/geofleet/user-service/singleton"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ProvideDBPool provides the database connection
func ProvideDBPool(ctx context.Context) (*pgxpool.Pool, error) {
	conn := singleton.GetConn()
	if conn == nil || conn.DB == nil {
		return nil, errors.New("database connection is not initialized")
	}
	return conn.DB, nil
}

// ProvideRepository provides the SQLC queries
func ProvideRepository(db *pgxpool.Pool) *postgresql.Queries {
	return postgresql.New(db)
}

// ProvideUserUsecase provides the usecase struct
func ProvideUserUsecase(querier *postgresql.Queries) *usecase.UserUsecaseInteractor {
	return usecase.NewUserUsecaseInteractor(querier)
}

// ProviderSet groups these together (Optional, but clean)
var ProviderSet = wire.NewSet(
	ProvideDBPool,
	ProvideRepository,
	ProvideUserUsecase,
	wire.Bind(new(usecase.UserUsecaseService), new(*usecase.UserUsecaseInteractor)),
	handler.NewUserRestfulHandler,
)
