package nagiosplugin

// Exit constants.
// These are from Nagios.
const (
	ExitHostUp   int = 0
	ExitHostDown int = 2

	ExitServiceOK       int = 0
	ExitServiceWarning  int = 1
	ExitServiceCritical int = 2
	ExitServiceUnknown  int = 3
)

// Units constants.
// These are supported by the `nagiosgraph` plugin.
const (
	UnitsNone string = ""

	UnitsSeconds      string = "s"
	UnitsMilliseconds string = "ms"
	UnitsMicroseconds string = "us"

	UnitsPercentage string = "%"

	UnitsBytes     string = "B"
	UnitsKilobytes string = "KB"
	UnitsMegabytes string = "MB"
	UnitsGigabytes string = "GB"
	UnitsTerabytes string = "TB"
	UnitsPetabytes string = "PB"

	UnitsCounter string = "c"
)
