package exception

import (
	"github.com/golibs-starter/golib/exception"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
)

var (
	GetArgumentException = exception.New(common.GeneralServiceUnavailable,
		common.GetErrorResponse(common.GeneralServiceUnavailable).Message)
	InternalServerErrorException = exception.New(common.GeneralServiceUnavailable, common.GetErrorResponse(common.GeneralServiceUnavailable).Message)
)
