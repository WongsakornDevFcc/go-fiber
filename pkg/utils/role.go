package utils

import (
	"fmt"

	"go-fiber/pkg/repository"
)

func VerifyRole(role string) (string, error) {
	switch role {
	case repository.AdminRoleName:
	case repository.ModeratorRoleName:
	case repository.UserRoleName:
	default:
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
