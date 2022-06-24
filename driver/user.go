package driver

import (
	"log"
	"net/http"

	"github.com/miyamoto-jo/clean-architecture-go/adapter/controller"
	"github.com/miyamoto-jo/clean-architecture-go/adapter/gateway"
	"github.com/miyamoto-jo/clean-architecture-go/adapter/presenter"
	"github.com/miyamoto-jo/clean-architecture-go/usecase/interactor"
)

func Serve(addr string) {
	user := controller.User{
		InputFactory:  interactor.NewUserInputPort,
		OutputFactory: presenter.NewUserOutputPort,
		RepoFactory:   gateway.NewMockUserRepository,
		// RepoFactory: gateway.NewSQLUserRepository, repositoryを差し替えた例
	}

	http.HandleFunc("/user/", user.GetUserController)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
