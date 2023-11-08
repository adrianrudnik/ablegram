package config

type Config struct {
	About      About            `yaml:"about"`
	Log        LogConfig        `yaml:"log"`
	Behaviour  BehaviourConfig  `yaml:"behaviour"`
	Collector  CollectorConfig  `yaml:"collector"`
	Parser     ParserConfig     `yaml:"parser"`
	Indexer    IndexerConfig    `yaml:"indexer"`
	Webservice WebserviceConfig `yaml:"webservice"`
}

type About struct {
	Version string `yaml:"version"`
	Commit  string `yaml:"commit"`
	Date    string `yaml:"date"`
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
	WorkerCount          int      `yaml:"worker_count"`
	WorkerDelayInMs      int      `yaml:"worker_delay_in_milliseconds"`
	SearchablePaths      []string `yaml:"searchable_paths"`
	ExcludeSystemFolders bool     `yaml:"exclude_system_folders"`
}

type ParserConfig struct {
	WorkerCount     int `yaml:"worker_count"`
	WorkerDelayInMs int `yaml:"worker_delay_in_milliseconds"`
}

type IndexerConfig struct {
	WorkerDelayInMs int `yaml:"worker_delay_in_milliseconds"`
}

type WebserviceConfig struct {
	TryPorts   []int `yaml:"try_ports"`
	ChosenPort int   `yaml:"-"`
}
