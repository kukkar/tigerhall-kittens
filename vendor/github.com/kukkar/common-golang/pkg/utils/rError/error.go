package rError

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-pg/pg"
	"github.com/jinzhu/gorm"
)

//As will check given error is of rError type
func As(err error) bool {
	var e *Error
	return errors.As(err, &e)
}

//NeedWrap will check error is not nil and not of rError type
//i.e it need to wrap with rError
func NeedWrap(err error) bool {
	if err != nil && !As(err) {
		return true
	}
	return false
}

//CustomError : custom error
func CustomError(c context.Context, debugMsg ...string) *Error {
	return newError(c, nil, CustomCode, debugMsg...)
}

//UnmarshalError : error occured while unmarshal
func UnmarshalError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, UnmarshalCode, debugMsg...)
}

//MarshalError : error occured while unmarshal
func MarshalError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, MarshalCode, debugMsg...)
}

//MiscError : error occured while processing
func MiscError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, MiscCode, debugMsg...)
}

//ConnError : error occured while creating connection
func ConnError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, ConnCode, debugMsg...)
}

//SelectError : error occured while select query
func SelectError(c context.Context, err error, debugMsg ...string) *Error {
	if err == pg.ErrNoRows {
		return NotFoundError(c, debugMsg...)
	}
	return newError(c, err, SelectCode, debugMsg...)
}

//SelectIgnoreNoRow : error occured while select query by ignoring no row error
func SelectIgnoreNoRow(c context.Context, err error, debugMsg ...string) error {
	if err != sql.ErrNoRows && err != gorm.ErrRecordNotFound {
		return newError(c, err, SelectCode, debugMsg...)
	}
	return nil
}

//CreateError : error occured while creating table
func CreateError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, CreateCode, debugMsg...)
}

//InsertError : error occured while insert query
func InsertError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, InsertCode, debugMsg...)
}

//UpdateError : error occured while update query
func UpdateError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, UpdateCode, debugMsg...)
}

//DeleteError : error occured while delete query
func DeleteError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, DeleteCode, debugMsg...)
}

//DropError : error occured while drop query
func DropError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, DropCode, debugMsg...)
}

//ExecError : error occured while exec query
func ExecError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, ExecCode, debugMsg...)
}

//TxError : error occured while starting transaction
func TxError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, TxCode, debugMsg...)
}

//NotFoundError : error occured when id not found
func NotFoundError(c context.Context, debugMsg ...string) *Error {
	if len(debugMsg) == 0 {
		debugMsg = append(debugMsg, "Record not exists")
	}
	return newError(c, nil, NotFoundCode, debugMsg...)
}

//BadReqError : error occured while validating request
//like while typecasting request, fk in request dosn't exists
func BadReqError(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, BadReqCode, debugMsg...)
}

//ForbiddenErr : forbidden access
func ForbiddenErr(c context.Context, debugMsg ...string) *Error {
	return newError(c, nil, ForbiddenCode, debugMsg...)
}

//UnauthoriseErr : unauthorized access
func UnauthoriseErr(c context.Context, debugMsg ...string) *Error {
	return newError(c, nil, UnauthoriseCode, debugMsg...)
}

// RateLimitExceedErr : service limit exceeded
func RateLimitExceedErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, RateLimitCode, debugMsg...)
}

// MaxAttemptReachedErr : max retry attempt reached
func MaxAttemptReachedErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, MaxAttemptReachedCode, debugMsg...)
}

// TokenExpiredErr : token or otp expired
func TokenExpiredErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, TokenExpiredCode, debugMsg...)
}

func AccountAlreadyExistErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, AccountAlreadyExistCode, debugMsg...)
}

func AccountAlreadyMapErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, AccountAlreadyExistCode, debugMsg...)
}

func PreConditionFailedErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, PreConditionFailedCode, debugMsg...)
}
func WrongBankDetailErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, WrongBankDetailCode, debugMsg...)
}
func DeviceAlreadyUsedErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, DeviceAlreadyUsedCode, debugMsg...)
}
func InvalidVPAErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, InvalidVPA, debugMsg...)
}
func OrganizedUserErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, OrganizedUser, debugMsg...)
}

func PartialSuccessErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, PartialSuccess, debugMsg...)
}

func NoTxnErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, NoTxn, debugMsg...)
}
func NoSettlementErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, NoSettlement, debugMsg...)
}

//MiscError : error occured while processing
func InsufficientBalanceErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, InsufficientBalanceCode, debugMsg...)
}
func SameNotificationErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, SameNotificationCode, debugMsg...)
}

func OrderStickerValidationErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, OrderStickerValidationCode, debugMsg...)
}
func OrderQRCancelErr(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, CancelOrderCode, debugMsg...)
}

func PanNameNotMatch(c context.Context, err error, debugMsg ...string) *Error {
	return newError(c, err, PanNameNotMatchCode, debugMsg...)
}
