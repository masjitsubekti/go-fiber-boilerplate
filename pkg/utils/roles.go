package utils

import (
	"fmt"

	"github.com/create-go-app/fiber-go-template/pkg/constant"
)

// VerifyRole func for verifying a given role.
func VerifyRole(role string) (string, error) {
	// Switch given role.
	switch role {
	case constant.AdminRoleName:
		// Nothing to do, verified successfully.
	case constant.ModeratorRoleName:
		// Nothing to do, verified successfully.
	case constant.UserRoleName:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
