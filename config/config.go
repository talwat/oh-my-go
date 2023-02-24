package config

import (
	"os"
	"strings"
)

var (
	Plugins = strings.Split(os.Getenv("OMGO_PLUGINS"), " ")
)
