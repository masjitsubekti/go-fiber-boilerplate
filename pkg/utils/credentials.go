package utils

import (
	"fmt"

	"github.com/create-go-app/fiber-go-template/pkg/constant"
)

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role string) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// Switch given role.
	switch role {
	case constant.AdminRoleName:
		// Admin credentials (all access).
		credentials = []string{
			constant.BookCreateCredential,
			constant.BookUpdateCredential,
			constant.BookDeleteCredential,
		}
	case constant.ModeratorRoleName:
		// Moderator credentials (only book creation and update).
		credentials = []string{
			constant.BookCreateCredential,
			constant.BookUpdateCredential,
		}
	case constant.UserRoleName:
		// Simple user credentials (only book creation).
		credentials = []string{
			constant.BookCreateCredential,
		}
	default:
		// Return error message.
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
