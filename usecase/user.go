package usecase

import (
	"context"
	"strconv"

	"github.com/miyamoto-jo/clean-architecture-go/port"
)

type UserInteractor struct {
	Output port.UserOutputPort
	Repo   port.UserRepository
}

func NewUserInteractor(out port.UserOutputPort, repo port.UserRepository) port.UserInputPort {
	return &UserInteractor{
		Output: out,
		Repo:   repo,
	}
}

func (m *UserInteractor) GetUserByID(ctx context.Context, inputUserID string) {
	userIDInt, err := strconv.Atoi(inputUserID)
	if err != nil {
		m.Output.RenderError(err)
		return
	}
	user, err := m.Repo.GetUserByID(ctx, userIDInt)
	if err != nil {
		m.Output.RenderError(err)
		return
	}
	m.Output.Render(user)
}

func (m *UserInteractor) AddUser(ctx context.Context, userName string) {
	users, err := m.Repo.GetAllUser(ctx)
	if err != nil {
		m.Output.RenderError(err)
		return
	}
	lastUserID := len(users)
	addUser, err := m.Repo.AddUser(ctx, lastUserID+1, userName)
	if err != nil {
		m.Output.RenderError(err)
		return
	}
	m.Output.Render(addUser)
}
