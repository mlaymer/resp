package resp

import "errors"

// ErrConstraintsViolated указывает на нарушение ограничений
// структур протокола.
var ErrConstraintsViolated = errors.New("RESP: constrains violated error")

// Validator оборачивает метод Validate,
// который осуществляет проверку структур
// протокола на соответствие предопределённым
// значениям.
// Реализации обязаны возвращать ErrConstraintsViolated.
type Validator interface {
	Validate() error
}
