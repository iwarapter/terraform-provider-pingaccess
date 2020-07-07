package backup

import "net/http"

type BackupAPI interface {
	BackupCommand() (resp *http.Response, err error)
}
