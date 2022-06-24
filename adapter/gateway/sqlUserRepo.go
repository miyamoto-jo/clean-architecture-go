package gateway

import (
	"context"
	"errors"

	"github.com/miyamoto-jo/clean-architecture-go/entity"
	"github.com/miyamoto-jo/clean-architecture-go/usecase/port"
)

type SQLUserRepository struct {
	// dbコネクションとか
}

func NewSQLUserRepository() port.UserRepository {
	return &SQLUserRepository{}
}

func (m *SQLUserRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	// sqlの処理
	return nil, errors.New(`SQLUserRepository: not implements`)
}
