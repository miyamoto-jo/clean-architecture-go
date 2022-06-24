package interactor

import (
	"context"

	"github.com/miyamoto-jo/clean-architecture-go/usecase/port"
)

type User struct {
	Output port.UserOutputPort
	Repo   port.UserRepository
}

func NewUserInputPort(out port.UserOutputPort, repo port.UserRepository) port.UserInputPort {
	return &User{
		Output: out,
		Repo:   repo,
	}
}

// TODO: メソッド名おかしくない？ここでは取得したユーザー情報をレンダリングするためにoutputに流している。
func (m *User) GetUserByID(ctx context.Context, userID string) {
	user, err := m.Repo.GetUserByID(ctx, userID)
	if err != nil {
		m.Output.RenderError(err)
		return
	}
	m.Output.Render(user)
}
