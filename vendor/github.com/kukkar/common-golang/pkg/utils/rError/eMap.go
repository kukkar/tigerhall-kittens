package rError

var errMap map[int]string

//error code
const (
	CustomCode = iota + 1
	UnmarshalCode
	MarshalCode
	BadReqCode
	ForbiddenCode
	UnauthoriseCode
	NotFoundCode
	MiscCode
	ConnCode
	TxCode
	CreateCode
	SelectCode
	InsertCode
	UpdateCode
	DeleteCode
	DropCode
	ExecCode
	RateLimitCode
	MaxAttemptReachedCode
	TokenExpiredCode
	AccountAlreadyExistCode
	AccountAlreadyMappedCode
	PreConditionFailedCode
	WrongBankDetailCode
	DeviceAlreadyUsedCode
	InvalidVPA
	OrganizedUser
	PartialSuccess
	NoTxn
	NoSettlement
	InsufficientBalanceCode
	SameNotificationCode
	OrderStickerValidationCode
	CancelOrderCode
	PanNameNotMatchCode
)

//msg constant
const (
	InvalidReqMsg               = "Invalid Request, Please contact system adminstrator for further clarification."
	ForbiddenMsg                = "You are not allowed to perform this operation. Please contact system adminstrator."
	UnauthoriseMsg              = "You are not authorised"
	BlandReqMsg                 = "Blank request, Please provide input to process"
	NotFoundMsg                 = "Record not found"
	ServerErrorMsg              = "Sorry unable to process this request. Please Try Again"
	AccountAlreadyExistErrorMsg = "Account already exist with account number"
	AccountAlreadyMappedMsg     = "Account already mapped"
	PreConditionFailedMsg       = "Request can not be processed as conditions not satisfied"
	WrongBankDetailMessage      = "Bank Detail is not valid"
	TooManyRequest              = "Too many request in given time frame"
	PartialSuccessMessage       = "partial success"
	NoTxnMsg                    = "No UPI Txns. in last 15 days"
	NoSettlementMSG             = "No Settlement in Last 15 Days"
	InsufficientBalanceMsg      = "insufficient balance"
	SameNotificationMSg         = "Unable to Add Notification for same mobile"
	OrderStickerValidationMSG   = "order sticker validation to generate order sticker failed "
	OrderStickercacelMSG        = "not able to cancel order-qr"
	PanNameNotMatchMSG          = "pan card name not match with bank account beneficiary name"
)

func init() {
	errMap = map[int]string{
		UnmarshalCode:              InvalidReqMsg,
		MarshalCode:                InvalidReqMsg,
		BadReqCode:                 InvalidReqMsg,
		ForbiddenCode:              ForbiddenMsg,
		UnauthoriseCode:            UnauthoriseMsg,
		NotFoundCode:               NotFoundMsg,
		MiscCode:                   ServerErrorMsg,
		ConnCode:                   ServerErrorMsg,
		TxCode:                     ServerErrorMsg,
		CreateCode:                 ServerErrorMsg,
		SelectCode:                 ServerErrorMsg,
		InsertCode:                 ServerErrorMsg,
		UpdateCode:                 ServerErrorMsg,
		DeleteCode:                 ServerErrorMsg,
		DropCode:                   ServerErrorMsg,
		ExecCode:                   ServerErrorMsg,
		RateLimitCode:              TooManyRequest,
		MaxAttemptReachedCode:      ServerErrorMsg,
		TokenExpiredCode:           ServerErrorMsg,
		AccountAlreadyExistCode:    AccountAlreadyExistErrorMsg,
		AccountAlreadyMappedCode:   AccountAlreadyMappedMsg,
		PreConditionFailedCode:     PreConditionFailedMsg,
		WrongBankDetailCode:        WrongBankDetailMessage,
		PartialSuccess:             PartialSuccessMessage,
		NoTxn:                      NoTxnMsg,
		NoSettlement:               NoSettlementMSG,
		InsufficientBalanceCode:    InsufficientBalanceMsg,
		SameNotificationCode:       SameNotificationMSg,
		OrderStickerValidationCode: OrderStickerValidationMSG,
		CancelOrderCode:            OrderStickercacelMSG,
		PanNameNotMatchCode:        PanNameNotMatchMSG,
	}
}
