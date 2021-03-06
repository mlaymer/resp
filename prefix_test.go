package resp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefix_Validate(t *testing.T) {
	t.Parallel()

	// Arrange.
	tests := []struct {
		name        string
		prefix      Prefix
		expectedErr error
	}{
		{
			name:        `Префикс "простые строки"`,
			prefix:      PrefixSimpleStrings,
			expectedErr: nil,
		},
		{
			name:        `Префикс "ошибки"`,
			prefix:      PrefixErrors,
			expectedErr: nil,
		},
		{
			name:        `Префикс "числа"`,
			prefix:      PrefixIntegers,
			expectedErr: nil,
		},
		{
			name:        `Префикс "объёмные строки"`,
			prefix:      PrefixBulkStrings,
			expectedErr: nil,
		},
		{
			name:        `Префикс "массивы"`,
			prefix:      PrefixArrays,
			expectedErr: nil,
		},
		{
			name:        "Некорректный префикс",
			prefix:      '~',
			expectedErr: ErrConstraintsViolated,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Act.
			actualErr := test.prefix.Validate()

			// Assert.
			assert.Equal(t, test.expectedErr, actualErr)
		})
	}
}
