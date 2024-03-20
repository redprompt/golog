package config

import (
	"github.com/logrusorgru/aurora"
)

func getLabelAndColor(level Level, color bool) string {
	if level == Silent {
		return ""
	}
	aurora := aurora.NewAurora(color)

	tag := map[Level]string{
		Fatal:   aurora.Bold(aurora.Red("FTL")).String(),
		Error:   aurora.Red("ERR").String(),
		Info:    aurora.Blue("INF").String(),
		Warning: aurora.Yellow("WRN").String(),
		Debug:   aurora.Magenta("DBG").String(),
		Verbose: aurora.Green("VER").String(),
	}

	label := "[" + tag[level] + "] "
	return label
}

func Format(message string, level Level, color bool) string {
	label := getLabelAndColor(level, color)
	return label + message
}
