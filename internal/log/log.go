package log

import (
	"fmt"
	"os"
)

//nolint:forbidigo
func Log(msg string, params ...interface{}) {
	fmt.Printf("oh-my-go: %s\n", fmt.Sprintf(msg, params...))
}

func OutputLog(msg string, params ...interface{}) {
	fmt.Fprintf(os.Stdout, msg, params...)
}

func RawError(msg string, params ...interface{}) {
	Log("error: %s", fmt.Sprintf(msg, params...))
	os.Exit(1)
}

func Error(err error, msg string, params ...interface{}) {
	if err == nil {
		return
	}

	RawError("%s: %s", fmt.Sprintf(msg, params...), err)
	os.Exit(1)
}
