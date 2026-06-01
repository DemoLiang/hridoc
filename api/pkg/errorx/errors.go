package errorx

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	ErrSystem           = 10001
	ErrInvalidParam     = 10002
	ErrUserNotFound     = 20001
	ErrIdCardExists     = 20002
	ErrPassword         = 20003
	ErrTokenInvalid     = 20004
	ErrCategoryNotFound = 30001
	ErrCertNotFound     = 30002
	ErrExcelFormat      = 40001
	ErrExcelValidation  = 40002
	ErrPreviewExpired   = 40003
	ErrTaskNotFound     = 40004
	ErrTaskFailed       = 40005
	ErrUploadFailed     = 50001
	ErrInvalidFileType  = 50002
	ErrFileTooLarge     = 50003
	ErrMinIOFailed      = 50004
	ErrNoPermission     = 60001
)

func New(code int, msg string) error {
	return fmt.Errorf("%d: %s", code, msg)
}

func IsNotFound(err error) bool {
	return err == sqlx.ErrNotFound
}
