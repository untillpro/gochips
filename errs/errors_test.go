package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	var errs Errors
	require.Nil(t, errs)
	errs = errs.Add("error1")
	require.Equal(t, "error1", errs.Error())
	errs = errs.Add("error2")
	require.Equal(t, "Multiple errors:\nerror1\nerror2", errs.Error())
	errs = errs.AddE(errors.New("error3"))
	require.Equal(t, "Multiple errors:\nerror1\nerror2\nerror3", errs.Error())
}
