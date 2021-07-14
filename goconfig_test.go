package goconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	Option1 string
	Option2 int
}

func TestFlow(t *testing.T) {
	testStruct := TestStruct{
		Option1: "test",
		Option2: 10,
	}
	f, err := os.CreateTemp("", "")
	require.Nil(t, err, "temporary file could not be created")

	// Save the struct to file
	fname := f.Name()
	f.Close()

	// cleanup the file on exit
	defer os.Remove(fname)

	// marshal to temporary file
	err = Save(testStruct, fname)
	require.Nilf(t, err, "struct couldn't be saved: %s", err)

	// unmarshal to struct
	var teststructnew TestStruct
	err = Load(&teststructnew, fname)
	require.Nilf(t, err, "struct couldn't be loaded: %s", err)
	require.Equal(t, testStruct, teststructnew, "objects are different")
}
