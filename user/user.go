package user

import (
	"restaurant/types"

	"github.com/google/uuid"
)

func New(fisrtName, phoneNumber string) *types.Client {
	id := uuid.New()
	return &types.Client{
		ID: id.String(),
		FullName: fisrtName,
		PhoneNumber: phoneNumber,
	}
}