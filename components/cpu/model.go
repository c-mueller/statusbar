package cpu

import "time"

var DefaultConfiguration = CPULoadConfiguration{
	UpdateInterval:   time.Second * 1,
	ShowAverageLoad:  true,
	LoadAverageCount: 120,
}

type CPUComponentBuilder struct{}

type CPULoadComponent struct {
	Config          *CPULoadConfiguration
	cpuUpdateTicker *time.Ticker
	cpuLoads        []float64
	currentValue    string
	updateTimestamp time.Time
	recentAverages  []float64
	currentAverage  float64
}

type CPULoadConfiguration struct {
	UpdateInterval   time.Duration `yaml:"update_interval"`
	ShowAverageLoad  bool          `yaml:"show_average_load"`
	LoadAverageCount int           `yaml:"load_average_count"`
}
