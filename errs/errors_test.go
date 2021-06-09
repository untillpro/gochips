package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	var errs Errors
	require.Nil(t, errs)

	var newErr error
	errs.AddE(newErr)
	require.Nil(t, errs)

	errs.Add("error1")
	require.Equal(t, "error1", errs.Error())

	errs.Add("error2")
	require.Equal(t, "Multiple errors:\nerror1\nerror2", errs.Error())

	errs.AddE(errors.New("error3"))
	require.Equal(t, "Multiple errors:\nerror1\nerror2\nerror3", errs.Error())

	errs = append(errs, errors.New("error4"))
	require.Equal(t, "Multiple errors:\nerror1\nerror2\nerror3\nerror4", errs.Error())
	require.Equal(t, "error4", errs[3].Error())

	errs = errs.Add("error5")
	require.Equal(t, "Multiple errors:\nerror1\nerror2\nerror3\nerror4\nerror5", errs.Error())

	errsAsError := error(errs)
	require.Equal(t, "Multiple errors:\nerror1\nerror2\nerror3\nerror4\nerror5", errsAsError.Error())
}
