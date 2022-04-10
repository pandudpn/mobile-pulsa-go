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
	ErrInvoicePaid             ErrorCode = errors.New("invoice already paid")                                  // 01
	ErrBillUnpaid              ErrorCode = errors.New("bill unpaid")                                           // 02
	ErrInvalidRefId            ErrorCode = errors.New("invalid reference id")                                  // 03
	ErrBillExpired             ErrorCode = errors.New("billing expired")                                       // 04
	ErrUndefined               ErrorCode = errors.New("error undefined")                                       // 05, 201
	ErrTransactionNotFound     ErrorCode = errors.New("transaction not found")                                 // 06
	ErrTransactionFailed       ErrorCode = errors.New("transaction failed")                                    // 07
	ErrBillBlocked             ErrorCode = errors.New("billing id blocked")                                    // 08
	ErrInquiryFailed           ErrorCode = errors.New("inquiry failed")                                        // 09
	ErrBillNotAvailable        ErrorCode = errors.New("billing not yet available")                             // 10
	ErrLimitTopUp              ErrorCode = errors.New("top up reach limit using same destination in 1 day")    // 10
	ErrDuplicateRefId          ErrorCode = errors.New("duplicate ref id")                                      // 11
	ErrBalanceMax              ErrorCode = errors.New("balance maximum limit")                                 // 12
	ErrCustomerNumberBlock     ErrorCode = errors.New("customer number blocked")                               // 13
	ErrDestinationNumber       ErrorCode = errors.New("incorrect destination number")                          // 14
	ErrNumberNotSupported      ErrorCode = errors.New("number not supported")                                  // 15
	ErrOperatorNumberNotMatch  ErrorCode = errors.New("number not match with operator")                        // 16
	ErrInsufficientDeposit     ErrorCode = errors.New("insufficient deposit")                                  // 17
	ErrCodeNotFound            ErrorCode = errors.New("product code not found")                                // 20
	ErrAlreadyPaidAtCounter    ErrorCode = errors.New("payment have to be done via counter")                   // 30
	ErrTransactionRejectMax    ErrorCode = errors.New("transaction rejected due to exceeding maximal bill")    // 31
	ErrTransactionFailedPeriod ErrorCode = errors.New("transaction failed, please pay bill before")            // 32
	ErrTransactionCantProcess  ErrorCode = errors.New("transaction cannot be process")                         // 33
	ErrBillPaid                ErrorCode = errors.New("bill has been paid")                                    // 34
	ErrTransactionRejectUnpaid ErrorCode = errors.New("transaction rejected due to another unpaid")            // 35
	ErrExceedingDueDate        ErrorCode = errors.New("exceeding due date, please pay at counter")             // 36
	ErrPaymentFailed           ErrorCode = errors.New("payment failed")                                        // 37, 38
	ErrTransactionRejectPaid   ErrorCode = errors.New("transaction rejected due to all invoice has been paid") // 40
	ErrCantPaidAtCounter       ErrorCode = errors.New("can't be paid in counter")                              // 41
	ErrPaymentRequest          ErrorCode = errors.New("payment request haven't been received")                 // 42
	ErrInvalidIPAddress        ErrorCode = errors.New("ip address not in whitelist")                           // 102
	ErrProductOutOfService     ErrorCode = errors.New("product is temporarily out of service")                 // 106
	ErrXMLFormat               ErrorCode = errors.New("error in xml format")                                   // 107
	ErrPageNotFound            ErrorCode = errors.New("page not found")                                        // 117
	ErrLimitTopUpMonthly       ErrorCode = errors.New("monthly top up is limit")                               // 121
	ErrRegionTopUpBlocked      ErrorCode = errors.New("top up region blocked")                                 // 131
	ErrInvalid                 ErrorCode = errors.New("invalid zone id, server id or rolename")                // 141
	ErrInvalidUser             ErrorCode = errors.New("invalid user id")                                       // 142
	ErrMaxNumber               ErrorCode = errors.New("maximum 1 number 1 time in 1 day")                      // 202
	ErrNumberTooLong           ErrorCode = errors.New("number too long")                                       // 203
	ErrAuthentication          ErrorCode = errors.New("wrong authentication")                                  // 204
	ErrCommand                 ErrorCode = errors.New("wrong command")                                         // 205
	ErrDestinationBlocked      ErrorCode = errors.New("this destination number has been blocked")              // 206
	ErrMaxNumberWithCode       ErrorCode = errors.New("maximum 1 number with any code 1 time in 1 day")        // 207
	ErrUsername                ErrorCode = errors.New("invalid username")                                      // 208
)

// getErrorCode will define ErrorCode from Response
func getErrorCode(v string) ErrorCode {
	// checking key of map is exists
	// for handle error go-routine
	errCode, ok := mapErrorCode[v]
	if !ok {
		return ErrParseFailed
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
