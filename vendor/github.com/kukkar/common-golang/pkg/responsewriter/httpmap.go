package responsewriter

import "github.com/kukkar/common-golang/pkg/utils/rError"

type HTTPCode int

const (
	HTTPStatusSuccessCode             HTTPCode = 200
	HTTPStatusBadRequestCode          HTTPCode = 400
	HTTPStatusTokenExpiredCode        HTTPCode = 400
	HTTPStatusInternalServerErrorCode HTTPCode = 412
	HTTPFatalErrorCode                HTTPCode = 412
	HTTPStatusNotFound                HTTPCode = 404
	HTTPRateLimitExceeded             HTTPCode = 429
	HTTPForbidden                     HTTPCode = 403
	HTTPPreConditionErrorCode         HTTPCode = 412
	HTTPStatusPartialSuccess          HTTPCode = 206
)

var appErrorToHttpCodeMap = map[int]HTTPCode{
	rError.CustomCode:                 HTTPPreConditionErrorCode,
	rError.UnmarshalCode:              HTTPStatusBadRequestCode,
	rError.MarshalCode:                HTTPStatusBadRequestCode,
	rError.BadReqCode:                 HTTPStatusBadRequestCode,
	rError.ForbiddenCode:              HTTPForbidden,
	rError.UnauthoriseCode:            HTTPForbidden,
	rError.NotFoundCode:               HTTPStatusNotFound,
	rError.MiscCode:                   HTTPPreConditionErrorCode,
	rError.ConnCode:                   HTTPStatusInternalServerErrorCode,
	rError.TxCode:                     HTTPStatusInternalServerErrorCode,
	rError.CreateCode:                 HTTPStatusInternalServerErrorCode,
	rError.SelectCode:                 HTTPStatusInternalServerErrorCode,
	rError.InsertCode:                 HTTPStatusInternalServerErrorCode,
	rError.UpdateCode:                 HTTPStatusInternalServerErrorCode,
	rError.DeleteCode:                 HTTPStatusInternalServerErrorCode,
	rError.DropCode:                   HTTPStatusInternalServerErrorCode,
	rError.ExecCode:                   HTTPStatusInternalServerErrorCode,
	rError.RateLimitCode:              HTTPRateLimitExceeded,
	rError.MaxAttemptReachedCode:      HTTPRateLimitExceeded,
	rError.TokenExpiredCode:           HTTPStatusTokenExpiredCode,
	rError.AccountAlreadyExistCode:    HTTPStatusSuccessCode,
	rError.AccountAlreadyMappedCode:   HTTPStatusSuccessCode,
	rError.WrongBankDetailCode:        HTTPStatusSuccessCode,
	rError.DeviceAlreadyUsedCode:      HTTPStatusSuccessCode,
	rError.InvalidVPA:                 HTTPPreConditionErrorCode,
	rError.OrganizedUser:              HTTPPreConditionErrorCode,
	rError.PartialSuccess:             HTTPStatusPartialSuccess,
	rError.NoTxn:                      HTTPStatusSuccessCode,
	rError.InsufficientBalanceCode:    HTTPPreConditionErrorCode,
	rError.OrderStickerValidationCode: HTTPPreConditionErrorCode,
	rError.CancelOrderCode:            HTTPPreConditionErrorCode,
	rError.PanNameNotMatchCode:        HTTPPreConditionErrorCode,
}
