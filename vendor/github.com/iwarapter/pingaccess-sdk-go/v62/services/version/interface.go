package version

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type VersionAPI interface {
	VersionCommand() (output *models.VersionDocClass, resp *http.Response, err error)
}
