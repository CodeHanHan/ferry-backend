package captcha

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DriverDigitFunc(t *testing.T) {
	id, b64s, err := DriverDigitFunc()
	require.NoError(t, err)
	require.NotEmpty(t, b64s)
	fmt.Printf("id: %s, b64s: %v\n", id, b64s)

	ok := store.Verify(id, "", true)
	require.Equal(t, false, ok)
}
