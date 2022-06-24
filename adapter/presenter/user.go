package presenter

import (
	"fmt"
	"net/http"

	"github.com/miyamoto-jo/clean-architecture-go/entity"
	"github.com/miyamoto-jo/clean-architecture-go/port"
)

// ユーザー構造体をプレーンテキストでアウトプットする構造体
type UserPlanTextOutput struct {
	w http.ResponseWriter
}

func NewUserPlanTextOutput(w http.ResponseWriter) port.UserOutputPort {
	return &UserPlanTextOutput{
		w: w,
	}
}

func (u *UserPlanTextOutput) Render(user *entity.User) {
	u.w.WriteHeader(http.StatusOK)
	fmt.Fprint(u.w, user)
}

func (u *UserPlanTextOutput) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(u.w, err)
}

// -------------------------------------------------------

// ユーザー構造体をjson形式でアウトプットする構造体
type UserJsonOutput struct {
	w http.ResponseWriter
}

func NewUserJsonTextOutput(w http.ResponseWriter) port.UserOutputPort {
	return &UserJsonOutput{
		w: w,
	}
}

func (u *UserJsonOutput) Render(user *entity.User) {
	// jsonでoutputする処理
}

func (u *UserJsonOutput) RenderError(err error) {
	// エラー処理
}
