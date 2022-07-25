package user

import (
	"restaurant/types"

	"github.com/google/uuid"
)

func New(number int) *types.Table {
	id := uuid.New()
	return &types.Table{
		ID: id.String(),
		Number: number,
	}
}