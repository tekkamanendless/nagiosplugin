package nagiosplugin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	fp := func(v float64) *float64 {
		return &v
	}

	t.Run("String", func(t *testing.T) {
		rows := []struct {
			description string
			output      Output
			result      string
		}{
			{
				description: "Empty",
				output:      Output{},
				result:      "\n",
			},
			{
				description: "Single line",
				output: Output{
					Output: "Line 1",
				},
				result: "Line 1\n",
			},
			{
				description: "Multiple lines",
				output: Output{
					Output:     "Line 1",
					LongOutput: []string{"Line 2", "Line 3"},
				},
				result: "Line 1\nLine 2\nLine 3\n",
			},
			{
				description: "Single line with single performance data",
				output: Output{
					Output: "Line 1",
					PerformanceData: []*PerformanceData{
						{
							Name:          "value-1",
							Value:         42,
							Units:         UnitsKilobytes,
							WarningValue:  fp(10),
							CriticalValue: fp(20),
							MinimumValue:  fp(0),
							MaximumValue:  fp(100),
						},
					},
				},
				result: "Line 1 | value-1=42KB;10;20;0;100\n",
			},
			{
				description: "Single line with multiple performance data",
				output: Output{
					Output: "Line 1",
					PerformanceData: []*PerformanceData{
						{
							Name:          "value-1",
							Value:         42,
							Units:         UnitsKilobytes,
							WarningValue:  fp(10),
							CriticalValue: fp(20),
							MinimumValue:  fp(0),
							MaximumValue:  fp(100),
						},
						{
							Name:          "value-2",
							Value:         50,
							Units:         UnitsKilobytes,
							WarningValue:  fp(10),
							CriticalValue: fp(20),
							MinimumValue:  fp(0),
							MaximumValue:  fp(100),
						},
					},
				},
				result: "Line 1 | value-1=42KB;10;20;0;100\nvalue-2=50KB;10;20;0;100\n",
			},
			{
				description: "Multiple lines with single performance data",
				output: Output{
					Output:     "Line 1",
					LongOutput: []string{"Line 2", "Line 3"},
					PerformanceData: []*PerformanceData{
						{
							Name:          "value-1",
							Value:         42,
							Units:         UnitsKilobytes,
							WarningValue:  fp(10),
							CriticalValue: fp(20),
							MinimumValue:  fp(0),
							MaximumValue:  fp(100),
						},
					},
				},
				result: "Line 1 | value-1=42KB;10;20;0;100\nLine 2\nLine 3\n",
			},
			{
				description: "Multiple lines with multiple performance data",
				output: Output{
					Output:     "Line 1",
					LongOutput: []string{"Line 2", "Line 3"},
					PerformanceData: []*PerformanceData{
						{
							Name:          "value-1",
							Value:         42,
							Units:         UnitsKilobytes,
							WarningValue:  fp(10),
							CriticalValue: fp(20),
							MinimumValue:  fp(0),
							MaximumValue:  fp(100),
						},
						{
							Name:          "value-2",
							Value:         50,
							Units:         UnitsKilobytes,
							WarningValue:  fp(10),
							CriticalValue: fp(20),
							MinimumValue:  fp(0),
							MaximumValue:  fp(100),
						},
					},
				},
				result: "Line 1 | value-1=42KB;10;20;0;100\nLine 2\nLine 3 | value-2=50KB;10;20;0;100\n",
			},
		}
		for rowIndex, row := range rows {
			t.Run(fmt.Sprintf("%d/%s", rowIndex, row.description), func(t *testing.T) {
				result := row.output.String()
				assert.Equal(t, row.result, result)
			})
		}
	})
}

func TestFormatFloat(t *testing.T) {
	rows := []struct {
		description string
		value       float64
		result      string
	}{
		{
			description: "Zero",
			value:       0,
			result:      "0",
		},
		{
			description: "One",
			value:       1,
			result:      "1",
		},
		{
			description: "Negative one",
			value:       -1,
			result:      "-1",
		},
		{
			description: "Medium number",
			value:       12345,
			result:      "12345",
		},
		{
			description: "Large number",
			value:       12345678910,
			result:      "12345678910",
		},
		{
			description: "Small number",
			value:       0.00001,
			result:      "0.00001",
		},
		{
			description: "Tiny number",
			value:       0.0000000001,
			result:      "0.0000000001",
		},
	}
	for rowIndex, row := range rows {
		t.Run(fmt.Sprintf("%d/%s", rowIndex, row.description), func(t *testing.T) {
			result := formatFloat(row.value)
			assert.Equal(t, row.result, result)
		})
	}
}

func TestPerformanceData(t *testing.T) {
	fp := func(v float64) *float64 {
		return &v
	}

	t.Run("SetWarningValue", func(t *testing.T) {
		var p PerformanceData
		assert.Nil(t, p.WarningValue)
		p.SetWarningValue(10)
		if assert.NotNil(t, p.WarningValue) {
			assert.Equal(t, float64(10), *p.WarningValue)
		}
	})
	t.Run("SetCriticalValue", func(t *testing.T) {
		var p PerformanceData
		assert.Nil(t, p.CriticalValue)
		p.SetCriticalValue(10)
		if assert.NotNil(t, p.CriticalValue) {
			assert.Equal(t, float64(10), *p.CriticalValue)
		}
	})
	t.Run("SetMinimumValue", func(t *testing.T) {
		var p PerformanceData
		assert.Nil(t, p.MinimumValue)
		p.SetMinimumValue(10)
		if assert.NotNil(t, p.MinimumValue) {
			assert.Equal(t, float64(10), *p.MinimumValue)
		}
	})
	t.Run("SetMaximumValue", func(t *testing.T) {
		var p PerformanceData
		assert.Nil(t, p.MaximumValue)
		p.SetMaximumValue(10)
		if assert.NotNil(t, p.MaximumValue) {
			assert.Equal(t, float64(10), *p.MaximumValue)
		}
	})
	t.Run("String", func(t *testing.T) {
		rows := []struct {
			description     string
			performanceData *PerformanceData
			result          string
		}{
			{
				description:     "Empty",
				performanceData: &PerformanceData{},
				result:          "=0;;;;",
			},
			{
				description: "All integer values",
				performanceData: &PerformanceData{
					Name:          "value-1",
					Value:         42,
					Units:         UnitsKilobytes,
					WarningValue:  fp(10),
					CriticalValue: fp(20),
					MinimumValue:  fp(0),
					MaximumValue:  fp(100),
				},
				result: "value-1=42KB;10;20;0;100",
			},
			{
				description: "All floating values",
				performanceData: &PerformanceData{
					Name:          "value-1",
					Value:         42.1,
					Units:         UnitsKilobytes,
					WarningValue:  fp(10.1),
					CriticalValue: fp(20.1),
					MinimumValue:  fp(0.1),
					MaximumValue:  fp(100.1),
				},
				result: "value-1=42.1KB;10.1;20.1;0.1;100.1",
			},
		}
		for rowIndex, row := range rows {
			t.Run(fmt.Sprintf("%d/%s", rowIndex, row.description), func(t *testing.T) {
				result := row.performanceData.String()
				assert.Equal(t, row.result, result)
			})
		}
	})
}
