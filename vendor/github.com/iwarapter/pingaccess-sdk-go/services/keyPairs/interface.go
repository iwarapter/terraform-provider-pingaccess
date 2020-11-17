package keyPairs

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type KeyPairsAPI interface {
	GetKeyPairsCommand(input *GetKeyPairsCommandInput) (output *models.KeyPairsView, resp *http.Response, err error)
	GenerateKeyPairCommand(input *GenerateKeyPairCommandInput) (output *models.KeyPairView, resp *http.Response, err error)
	ImportKeyPairCommand(input *ImportKeyPairCommandInput) (output *models.KeyPairView, resp *http.Response, err error)
	KeyAlgorithms() (output *models.KeyAlgorithmsView, resp *http.Response, err error)
	GetKeypairsCreatableGeneralNamesCommand() (output *models.SanTypes, resp *http.Response, err error)
	DeleteKeyPairCommand(input *DeleteKeyPairCommandInput) (resp *http.Response, err error)
	GetKeyPairCommand(input *GetKeyPairCommandInput) (output *models.KeyPairView, resp *http.Response, err error)
	PatchKeyPairCommand(input *PatchKeyPairCommandInput) (output *models.KeyPairView, resp *http.Response, err error)
	UpdateKeyPairCommand(input *UpdateKeyPairCommandInput) (output *models.KeyPairView, resp *http.Response, err error)
	ExportKeyPairCert(input *ExportKeyPairCertInput) (output *string, resp *http.Response, err error)
	GenerateCsrCommand(input *GenerateCsrCommandInput) (output *string, resp *http.Response, err error)
	ImportCSRResponseCommand(input *ImportCSRResponseCommandInput) (output *models.KeyPairView, resp *http.Response, err error)
	ExportKeyPair(input *ExportKeyPairInput) (resp *http.Response, err error)
	DeleteChainCertificateCommand(input *DeleteChainCertificateCommandInput) (resp *http.Response, err error)
}
