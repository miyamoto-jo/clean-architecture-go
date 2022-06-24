package gateway

import (
	"context"
	"errors"
	"fmt"

	"github.com/miyamoto-jo/clean-architecture-go/entity"
	"github.com/miyamoto-jo/clean-architecture-go/port"
)

var mockUsersMaster = []*entity.User{
	{ID: 1, Name: "miyamoto-jo"},
	{ID: 2, Name: "miyamoto-ikumi"},
	{ID: 3, Name: "miyamoto-yote"},
	{ID: 4, Name: "miyamoto-jiru"},
	{ID: 5, Name: "miyamoto-cona"},
}

// モックユーザーデータを参照するリポジトリ
type MockUserRepository struct{}

func NewMockUserRepository() port.UserRepository {
	return &MockUserRepository{}
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, userID int) (*entity.User, error) {
	for _, user := range mockUsersMaster {
		if user.ID == userID {
			return user, nil
		}
	}
	return nil, fmt.Errorf(`unknown user ID: %d`, userID)
}

func (m *MockUserRepository) GetAllUser(ctx context.Context) ([]*entity.User, error) {
	if mockUsersMaster == nil {
		return nil, errors.New(`mockdata is nil`)
	}
	return mockUsersMaster, nil
}

func (m *MockUserRepository) AddUser(ctx context.Context, userID int, userName string) (*entity.User, error) {
	// すでに存在するユーザー
	if _, err := m.GetUserByID(ctx, userID); err == nil {
		return nil, fmt.Errorf(`it is a user who already exists. userID: %d`, userID)
	}
	addUser := &entity.User{
		ID:   userID,
		Name: userName,
	}
	mockUsersMaster = append(mockUsersMaster, addUser)
	return addUser, nil
}

// -------------------------------------------------------

// SQLのユーザーデータを参照するリポジトリ
type SQLUserRepository struct {
	// dbコネクションとか
}

func NewSQLUserRepository() port.UserRepository {
	return &SQLUserRepository{}
}

func (m *SQLUserRepository) GetUserByID(ctx context.Context, userID int) (*entity.User, error) {
	// sqlの処理
	return nil, errors.New(`SQLUserRepository: not implements`)
}

func (m *SQLUserRepository) GetAllUser(ctx context.Context) ([]*entity.User, error) {
	// sqlの処理
	return nil, errors.New(`SQLUserRepository: not implements`)
}

func (m *SQLUserRepository) AddUser(ctx context.Context, userID int, userName string) (*entity.User, error) {
	// sqlの処理
	return nil, errors.New(`SQLUserRepository: not implements`)
}
