package config

type Config struct {
	Log        LogConfig        `yaml:"log"`
	Behaviour  BehaviourConfig  `yaml:"behaviour"`
	Collector  CollectorConfig  `yaml:"collector"`
	Webservice WebserviceConfig `yaml:"webservice"`
}

type LogConfig struct {
	Level          string `yaml:"level"`
	ToFiles        bool   `yaml:"to_files"`
	ScannedFolders bool   `yaml:"scanned_folders"`
}

type BehaviourConfig struct {
	WebserviceAutostart bool `yaml:"autostart_webservice"`
	BrowserAutostart    bool `yaml:"autostart_browser"`
	ShowGui             bool `yaml:"show_gui"`
}

type CollectorConfig struct {
	SearchablePaths      []string `yaml:"searchable_paths"`
	ExcludeSystemFolders bool     `yaml:"exclude_system_folders"`
}

type WebserviceConfig struct {
	TryPorts   []int `yaml:"try_ports"`
	ChosenPort int   `yaml:"-"`
}
