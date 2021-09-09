package nagiosplugin

import (
	"fmt"
	"strings"
)

// Output is the ultimate output of a plugin.
type Output struct {
	Output          string             // This is the main output line (SERVICEOUTPUT).
	LongOutput      []string           // This is the list of additional output lines (LONGSERVICEOUTPUT).
	PerformanceData []*PerformanceData // This is any performance data.
}

// String returns the proper Nagios-formatted string for the output.
func (o *Output) String() string {
	var performanceStrings []string
	for _, performanceData := range o.PerformanceData {
		performanceStrings = append(performanceStrings, performanceData.String())
	}

	result := o.Output
	if len(performanceStrings) > 0 {
		result += " | " + performanceStrings[0]
		performanceStrings = performanceStrings[1:]
	}
	result += "\n"
	if len(o.LongOutput) > 0 {
		result += strings.Join(o.LongOutput, "\n")
		if len(performanceStrings) == 0 {
			result += "\n"
		} else {
			result += " | "
		}
	}
	if len(performanceStrings) > 0 {
		result += strings.Join(performanceStrings, "\n") + "\n"
	}

	return result
}

// PerformanceData is a data point for "nagiosgraph" to process.
type PerformanceData struct {
	Name  string  // This is the name of the metric.
	Value float64 // This is the value of the metric.
	Units string  // This is the units that it's measured in.

	WarningValue  *float64 // This is the warning value, if any.
	CriticalValue *float64 // This is the critical value, if any.
	MinimumValue  *float64 // This is the minimum value, if any.
	MaximumValue  *float64 // This is the maximum value, if any.
}

func (p *PerformanceData) SetWarningValue(v float64) {
	p.WarningValue = &v
}

func (p *PerformanceData) SetCriticalValue(v float64) {
	p.CriticalValue = &v
}

func (p *PerformanceData) SetMinimumValue(v float64) {
	p.MinimumValue = &v
}

func (p *PerformanceData) SetMaximumValue(v float64) {
	p.MaximumValue = &v
}

// formatFloat returns a string from a floating point value.
// If the value is a whole number, then no decimal place will be shown.
func formatFloat(f float64) string {
	v := fmt.Sprintf("%f", f)
	v = strings.TrimSuffix(v, ".000000")
	return v
}

// This serializes the performance data as a string in the format that `nagiosgraph` wants.
func (p *PerformanceData) String() string {
	valueString := formatFloat(p.Value)

	var warningString string
	if p.WarningValue != nil {
		warningString = formatFloat(*p.WarningValue)
	}
	var criticalString string
	if p.CriticalValue != nil {
		criticalString = formatFloat(*p.CriticalValue)
	}
	var minimumString string
	if p.MinimumValue != nil {
		minimumString = formatFloat(*p.MinimumValue)
	}
	var maximumString string
	if p.MaximumValue != nil {
		maximumString = formatFloat(*p.MaximumValue)
	}

	result := fmt.Sprintf("%s=%s%s;%s;%s;%s;%s", p.Name, valueString, p.Units, warningString, criticalString, minimumString, maximumString)

	return result
}
