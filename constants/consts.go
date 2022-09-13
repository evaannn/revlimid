package constants

import (
	"os"
	"sync"

	"github.com/gookit/color"
)

var (
	White  = color.FgWhite.Render
	Red    = color.FgRed.Render
	Green  = color.FgGreen.Render
	Yellow = color.FgYellow.Render

	Tokens  []string
	Proxies []string
	Wg      sync.WaitGroup
	Logging bool
	Proxy   bool
	LogFile *os.File
)
