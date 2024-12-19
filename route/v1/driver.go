package v1

import (
	"github.com/dappsteros-io/DappsterOS-Common/model"
	"github.com/dappsteros-io/DappsterOS-Common/utils/common_err"
	"github.com/dappsteros-io/DappsterOS-LocalStorage/internal/op"
	"github.com/labstack/echo/v4"
)

func ListDriverInfo(ctx echo.Context) error {
	return ctx.JSON(common_err.SUCCESS, model.Result{Success: common_err.SUCCESS, Message: common_err.GetMsg(common_err.SUCCESS), Data: op.GetDriverInfoMap()})
}
