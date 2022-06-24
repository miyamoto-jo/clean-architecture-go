package gateway

import (
	"context"
	"errors"

	"github.com/miyamoto-jo/clean-architecture-go/entity"
	"github.com/miyamoto-jo/clean-architecture-go/usecase/port"
)

type MockUserRepository struct {
	mockdata map[string]*entity.User
}

func NewMockUserRepository() port.UserRepository {
	return &MockUserRepository{
		mockdata: map[string]*entity.User{
			"1": {
				ID:   "D03JJsdaGF",
				Name: "miyamoto-jo",
			},
			"2": {
				ID:   "D01IKUMISDAD",
				Name: "miyamoto-ikumi",
			},
			"3": {
				ID:   "D01dsahJJISDWQ",
				Name: "miyamoto-yote",
			},
			"4": {
				ID:   "D01dsahJJIJIRU",
				Name: "miyamoto-jiru",
			},
			"5": {
				ID:   "D01dsahJJICONA",
				Name: "miyamoto-cona",
			},
		},
	}
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	u, ok := m.mockdata[userID]
	if !ok {
		return nil, errors.New(`unknown user ID: ` + userID)
	}
	return u, nil
}
