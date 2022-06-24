package port

import (
	"context"

	"github.com/miyamoto-jo/clean-architecture-go/entity"
)

// 入力ポート
type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
}

// 出力ポート
type UserOutputPort interface {
	Render(*entity.User)
	RenderError(error)
}

// データ取得ポート
type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
}
