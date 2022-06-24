package driver

import (
	"log"
	"net/http"

	"github.com/miyamoto-jo/clean-architecture-go/adapter/controller"
	"github.com/miyamoto-jo/clean-architecture-go/adapter/gateway"
	"github.com/miyamoto-jo/clean-architecture-go/adapter/presenter"
	"github.com/miyamoto-jo/clean-architecture-go/usecase"
)

func Serve(addr string) {
	userController := controller.UserController{
		InputFactory:  usecase.NewUserInteractor,
		OutputFactory: presenter.NewUserPlanTextOutput,
		// OutputFactory: presenter.NewUserJsonTextOutput ← 出力をjsonに変更した例
		RepoFactory: gateway.NewMockUserRepository,
		// RepoFactory: gateway.NewSQLUserRepository, ← epositoryをsqlに変更した例
	}

	http.HandleFunc("/useradd/", userController.AddUser)
	http.HandleFunc("/user/", userController.GetUser)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
