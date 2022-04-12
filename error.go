package mobilepulsa

import (
	"encoding/json"
	"errors"
)

type ErrorCode error

var mapErrorCode = map[string]ErrorCode{
	"01":  ErrInvoicePaid,
	"02":  ErrBillUnpaid,
	"03":  ErrInvalidRefId,
	"04":  ErrBillExpired,
	"05":  ErrUndefined,
	"06":  ErrTransactionNotFound,
	"07":  ErrTransactionFailed,
	"08":  ErrBillBlocked,
	"09":  ErrInquiryFailed,
	"10":  ErrLimitTopUp,
	"11":  ErrDuplicateRefId,
	"12":  ErrBalanceMax,
	"13":  ErrCustomerNumberBlock,
	"14":  ErrDestinationNumber,
	"15":  ErrNumberNotSupported,
	"16":  ErrOperatorNumberNotMatch,
	"17":  ErrInsufficientDeposit,
	"20":  ErrCodeNotFound,
	"30":  ErrAlreadyPaidAtCounter,
	"31":  ErrTransactionRejectMax,
	"32":  ErrTransactionFailedPeriod,
	"33":  ErrTransactionCantProcess,
	"34":  ErrBillPaid,
	"35":  ErrTransactionRejectUnpaid,
	"36":  ErrExceedingDueDate,
	"37":  ErrPaymentFailed,
	"38":  ErrPaymentFailed,
	"40":  ErrTransactionRejectPaid,
	"41":  ErrCantPaidAtCounter,
	"42":  ErrPaymentRequest,
	"102": ErrInvalidIPAddress,
	"106": ErrProductOutOfService,
	"107": ErrXMLFormat,
	"117": ErrPageNotFound,
	"121": ErrLimitTopUpMonthly,
	"131": ErrRegionTopUpBlocked,
	"141": ErrInvalid,
	"142": ErrInvalidUser,
	"201": ErrUndefined,
	"202": ErrMaxNumber,
	"203": ErrNumberTooLong,
	"204": ErrAuthentication,
	"205": ErrCommand,
	"206": ErrDestinationBlocked,
	"207": ErrMaxNumberWithCode,
	"208": ErrUsername,
}

var (
	// Error Parameter not filled
	ErrInvalidAPIKey = errors.New("invalid api key")
	ErrAPIKeyNil     = errors.New("api-key is nil")
	ErrUsernameNil   = errors.New("username is nil")
	ErrParseFailed   = errors.New("failed parse error response")

	// Error from response
	// ErrInvoicePaid when customer check payment and the invoice already paid
	ErrInvoicePaid ErrorCode = errors.New("invoice already paid") // 01
	// ErrBillUnpaid error bill not yet paid
	ErrBillUnpaid ErrorCode = errors.New("bill unpaid") // 02
	// ErrInvalidRefId invalid reference id
	ErrInvalidRefId ErrorCode = errors.New("invalid reference id") // 03
	// ErrBillExpired bill expired
	ErrBillExpired ErrorCode = errors.New("billing expired") // 04
	// ErrUndefined error not yet defined
	ErrUndefined ErrorCode = errors.New("error undefined") // 05, 201
	// ErrTransactionNotFound transaction not found
	ErrTransactionNotFound ErrorCode = errors.New("transaction not found") // 06
	// ErrTransactionFailed the transaction is failed
	ErrTransactionFailed ErrorCode = errors.New("transaction failed") // 07
	// ErrBillBlocked billing id of customer get blocked
	ErrBillBlocked ErrorCode = errors.New("billing id blocked") // 08
	// ErrInquiryFailed inquiry failed
	ErrInquiryFailed ErrorCode = errors.New("inquiry failed") // 09
	// ErrBillNotAvailable billing not yet available
	ErrBillNotAvailable ErrorCode = errors.New("billing not yet available") // 10
	// ErrLimitTopUp limit top up in a day within same destination (customer_id)
	ErrLimitTopUp ErrorCode = errors.New("top up reach limit using same destination in 1 day") // 10
	// ErrDuplicateRefId reference id already created
	ErrDuplicateRefId ErrorCode = errors.New("duplicate ref id") // 11
	// ErrBalanceMax balance is reach maximum
	ErrBalanceMax ErrorCode = errors.New("balance maximum limit") // 12
	// ErrCustomerNumberBlock customer number already got blacklist
	ErrCustomerNumberBlock ErrorCode = errors.New("customer number blocked") // 13
	// ErrDestinationNumber destination_number and product code not match
	ErrDestinationNumber ErrorCode = errors.New("incorrect destination number") // 14
	// ErrNumberNotSupported number not support at this ProductCode
	ErrNumberNotSupported ErrorCode = errors.New("number not supported") // 15
	// ErrOperatorNumberNotMatch destination number not match with operator product
	ErrOperatorNumberNotMatch ErrorCode = errors.New("number not match with operator") // 16
	// ErrInsuffienceDeposit no balance on your deposit
	ErrInsufficientDeposit ErrorCode = errors.New("insufficient deposit") // 17
	// ErrCodeNotFound product not found
	ErrCodeNotFound ErrorCode = errors.New("product code not found") // 20
	// ErrAlreadyPaidAtCounter customer already paid at counter
	ErrAlreadyPaidAtCounter ErrorCode = errors.New("payment have to be done via counter") // 30
	// ErrTransactionRejectMax transaction has been reject due to exceeding maximal billing
	ErrTransactionRejectMax ErrorCode = errors.New("transaction rejected due to exceeding maximal bill") // 31
	// ErrTransactionFailedPeriod transaction got failed because customer need pay bill on other period
	ErrTransactionFailedPeriod ErrorCode = errors.New("transaction failed, please pay bill before") // 32
	// ErrTransactionCantProcess transaction cannot be process. something got error
	ErrTransactionCantProcess ErrorCode = errors.New("transaction cannot be process") // 33
	// ErrBillPaid billing already paid
	ErrBillPaid ErrorCode = errors.New("bill has been paid") // 34
	// ErrTransactionRejectUnpaid transaction got reject due to another unpaid
	ErrTransactionRejectUnpaid ErrorCode = errors.New("transaction rejected due to another unpaid") // 35
	// ErrExceedingDueDate bill is expired, need to pay at counter
	ErrExceedingDueDate ErrorCode = errors.New("exceeding due date, please pay at counter") // 36
	// ErrPaymentFailed payment got error, internal server error
	ErrPaymentFailed ErrorCode = errors.New("payment failed") // 37, 38
	// ErrTransactionRejectPaid transaction got reject because all invoice has been paid
	ErrTransactionRejectPaid ErrorCode = errors.New("transaction rejected due to all invoice has been paid") // 40
	// ErrCantPaidAtCounter customer can't paid at counter
	ErrCantPaidAtCounter ErrorCode = errors.New("can't be paid in counter") // 41
	// ErrPaymentRequest payment request haven't been received
	ErrPaymentRequest ErrorCode = errors.New("payment request haven't been received") // 42
	// ErrInvalidIPAddress ip address not in range whitelist
	ErrInvalidIPAddress ErrorCode = errors.New("ip address not in whitelist") // 102
	// ErrProductOutOfService product get out of stock or not yet ready
	ErrProductOutOfService ErrorCode = errors.New("product is temporarily out of service") // 106
	// ErrXMLFormat request format xml got error
	ErrXMLFormat ErrorCode = errors.New("error in xml format") // 107
	// ErrPageNotFound API not found
	ErrPageNotFound ErrorCode = errors.New("page not found") // 117
	// ErrLimitTopUpMonthly customer reach the limit at this month
	ErrLimitTopUpMonthly ErrorCode = errors.New("monthly top up is limit") // 121
	// ErrRegionTopUpBlocked region get blocked for top up
	ErrRegionTopUpBlocked ErrorCode = errors.New("top up region blocked") // 131
	// ErrInvalid invalid zone id, service id or role
	ErrInvalid ErrorCode = errors.New("invalid zone id, server id or rolename") // 141
	// ErrInvalidUser invalid user_id
	ErrInvalidUser ErrorCode = errors.New("invalid user id") // 142
	// ErrMaxNumber customer only can 1 transaction in a day
	ErrMaxNumber ErrorCode = errors.New("maximum 1 number 1 time in 1 day") // 202
	// ErrNumberTooLong customer_id too long. please check the documentation
	ErrNumberTooLong ErrorCode = errors.New("number too long") // 203
	// ErrAuthentication unauthenticated
	ErrAuthentication ErrorCode = errors.New("wrong authentication") // 204
	// ErrCommand wrong command or payload
	ErrCommand ErrorCode = errors.New("wrong command") // 205
	// ErrDestinationBlocked destination number got blocked on operator. please contact customer service
	ErrDestinationBlocked ErrorCode = errors.New("this destination number has been blocked") // 206
	// ErrMaxNumberWithCode 1 number only can 1 transaction for any code
	ErrMaxNumberWithCode ErrorCode = errors.New("maximum 1 number with any code 1 time in 1 day") // 207
	// ErrUsername error username options
	ErrUsername ErrorCode = errors.New("invalid username") // 208
)

// getErrorCode will define ErrorCode from Response
func getErrorCode(v string) ErrorCode {
	// checking key of map is exists
	// for handle error go-routine
	errCode, ok := mapErrorCode[v]
	if !ok {
		return nil
	}

	return errCode
}

// ErrorResponse is data model for error response from API Call
type ErrorResponse struct {
	Data DataError `json:"data"`
}

// DataError is data model for error message and status from API Call
type DataError struct {
	RC           string              `json:"rc,omitempty"`
	ResponseCode string              `json:"response_code,omitempty"`
	Message      string              `json:"message"`
	Status       int                 `json:"status,omitempty"`
	ErrorDetails map[string][]string `json:"error_details,omitempty"`
}

// ErrorHttp it will parse error message from Response "rc" into error type
func ErrorHttp(respBody []byte) ErrorCode {
	var errResponse ErrorResponse

	err := json.Unmarshal(respBody, &errResponse)
	if err != nil {
		return err
	}

	var rc = errResponse.Data.ResponseCode
	if rc == "" {
		rc = errResponse.Data.RC
	}

	return getErrorCode(rc)
}
