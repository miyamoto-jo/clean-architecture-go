// アダプターの中のcontrollerなので、
// ゲームで例えると、その名のとおりコントローラーの役目がある
// inputした値を業務ロジックに流す役目がある
package controller

import (
	"net/http"
	"strings"

	"github.com/miyamoto-jo/clean-architecture-go/port"
)

type UserController struct {
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	InputFactory  func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
	RepoFactory   func() port.UserRepository
}

func (m *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	output := m.OutputFactory(w)
	repo := m.RepoFactory()
	input := m.InputFactory(output, repo)

	userID := strings.TrimPrefix(r.URL.Path, "/user/")
	input.GetUserByID(r.Context(), userID)
}

func (m *UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	output := m.OutputFactory(w)
	repo := m.RepoFactory()
	input := m.InputFactory(output, repo)

	userName := strings.TrimPrefix(r.URL.Path, "/useradd/")
	input.AddUser(r.Context(), userName)
}
