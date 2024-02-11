package log

import "runtime"

var (
	checked     = false
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	YellowDark  = "\033[33m"
	YellowLight = "\033[93m"
	Blue        = "\033[34m"
	Purple      = "\033[35m"
	Cyan        = "\033[36m"
	GrayLigth   = "\033[37m"
	GrayDark    = "\033[90m"
	White       = "\033[97m"
)

func check() {
	if checked {
		return
	}
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		YellowDark = ""
		YellowLight = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		GrayDark = ""
		GrayLigth = ""
		White = ""
	}

	checked = true
}
