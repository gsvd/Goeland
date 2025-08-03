package errorsx

type ErrorCode string

const (
	ErrCodeEmptyAddress     ErrorCode = "ERR_EMPTY_ADDRESS"
	ErrCodePasswordRequired ErrorCode = "ERR_PASSWORD_REQUIRED"
	ErrCodeAccountExists    ErrorCode = "ERR_ACCOUNT_EXISTS"
	ErrCodeUnknown          ErrorCode = "ERR_UNKNOWN"
)

type AppError struct {
	Code string `json:"code"`
}

func (e *AppError) Error() string {
	return e.Code
}

func NewAppError(code ErrorCode) *AppError {
	return &AppError{Code: string(code)}
}
