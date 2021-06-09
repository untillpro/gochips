package errs

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	var errs Errors
	println(cap(errs))
	require.Nil(t, errs)

	perr := &errs

	var newErr error
	errs.AddE(newErr)
	require.Nil(t, errs)

	errs.Add("error1")
	fmt.Printf("%p\n", perr)
	require.Equal(t, "error1", perr.Error())
	require.Equal(t, "error1", Errors(*perr)[0].Error())

	errs.Add("error2")
	fmt.Printf("%p\n", perr)
	require.Equal(t, "Multiple errors:\nerror1\nerror2", perr.Error())
	require.Equal(t, "error2", Errors(*perr)[1].Error())

	errs.AddE(errors.New("error3"))
	fmt.Printf("%p\n", perr)
	require.Equal(t, "Multiple errors:\nerror1\nerror2\nerror3", perr.Error())
	require.Equal(t, "error3", Errors(*perr)[2].Error())

	errs = append(errs, errors.New("error4"))
	require.Equal(t, "Multiple errors:\nerror1\nerror2\nerror3\nerror4", perr.Error())
	require.Equal(t, "error4", Errors(*perr)[3].Error())

	errs = errs.Add("error5")
	require.Equal(t, "Multiple errors:\nerror1\nerror2\nerror3\nerror4\nerror5", perr.Error())

	errsAsError := error(perr)
	require.Equal(t, "Multiple errors:\nerror1\nerror2\nerror3\nerror4\nerror5", errsAsError.Error())

	for i := 0; i < 1200; i++ {
		errs.Add("test" + strconv.Itoa(i))
	}
	fmt.Printf("%p\n", perr)
	require.Equal(t, 1205, len(*perr))
	require.Equal(t, "test1199", Errors(*perr)[1204].Error())
}
