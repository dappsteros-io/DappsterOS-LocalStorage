package v2

import (
	"github.com/dappsteros-io/DappsterOS-LocalStorage/codegen"
)

type LocalStorage struct{}

func NewLocalStorage() codegen.ServerInterface {
	return &LocalStorage{}
}
