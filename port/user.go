package port

import (
	"context"

	"github.com/miyamoto-jo/clean-architecture-go/entity"
)

// ユーザーのポート情報

// 入力ポート
type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
	AddUser(ctx context.Context, userName string)
}

// 出力ポート
type UserOutputPort interface {
	Render(*entity.User)
	RenderError(error)
}

// データ取得ポート
type UserRepository interface {
	GetUserByID(ctx context.Context, userID int) (*entity.User, error)
	GetAllUser(ctx context.Context) ([]*entity.User, error)
	AddUser(ctx context.Context, userID int, userName string) (*entity.User, error)
}
