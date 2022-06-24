package controller

import (
	"net/http"
	"strings"

	"github.com/miyamoto-jo/clean-architecture-go/usecase/port"
)

type User struct {
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	InputFactory  func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
	RepoFactory   func() port.UserRepository
}

func (m *User) GetUserController(w http.ResponseWriter, r *http.Request) {
	output := m.OutputFactory(w)
	repo := m.RepoFactory()
	input := m.InputFactory(output, repo)

	userID := strings.TrimPrefix(r.URL.Path, "/user/")
	input.GetUserByID(r.Context(), userID)
}
