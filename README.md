[![Go Report Card](https://goreportcard.com/badge/github.com/tekkamanendless/iaqualink)](https://goreportcard.com/report/github.com/tekkamanendless/iaqualink)
[![GoDoc](https://godoc.org/github.com/tekkamanendless/iaqualink?status.svg)](https://godoc.org/github.com/tekkamanendless/iaqualink)

# nagiosplugin
A Go package to facilitate creating Nagios plugins.

This package provides some constants as well as a structure to use to generate parseable output for use with Nagios and `nagiosgraph` for performance data.

## Examples

### Trivial success

```
func main() {
	var output nagiosplugin.Output
	output.Output = "Everything is OK"

	fmt.Printf("%s", output.String())
	os.Exit(nagiosplugin.ExitServiceOK)
}
```

### Performance data

```
func main() {
	var output nagiosplugin.Output
	output.Output = "The value is 42"

	p := &nagiosplugin.PerformanceData{
		Name: "my-value",
		Value: 42,
		Units: nagiosplugin.UnitsNone,
	}
	p.SetWarningValue(30)
	p.SetCriticalValue(50)
	output.PerformanceData = append(output.PerformanceData, p)

	fmt.Printf("%s", output.String())
	os.Exit(nagiosplugin.ExitServiceWarning)
}
```
