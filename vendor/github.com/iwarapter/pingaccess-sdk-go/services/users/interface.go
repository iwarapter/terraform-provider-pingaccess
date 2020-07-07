package users

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type UsersAPI interface {
	GetUsersCommand(input *GetUsersCommandInput) (output *models.UsersView, resp *http.Response, err error)
	GetUserCommand(input *GetUserCommandInput) (output *models.UserView, resp *http.Response, err error)
	UpdateUserCommand(input *UpdateUserCommandInput) (output *models.UserView, resp *http.Response, err error)
	UpdateUserPasswordCommand(input *UpdateUserPasswordCommandInput) (output *models.UserPasswordView, resp *http.Response, err error)
}
