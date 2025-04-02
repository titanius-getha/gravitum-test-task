package transport

import "errors"

var (
	ErrBadRequest = errors.New("неверные данные запроса")
	ErrInternal   = errors.New("внутренняя ошибка сервиса")
)
