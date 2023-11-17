package config

type Config struct {
	About      About            `yaml:"-" json:"-"`
	Log        LogConfig        `yaml:"log" json:"log"`
	Behaviour  BehaviourConfig  `yaml:"behaviour" json:"behaviour"`
	Collector  CollectorConfig  `yaml:"collector" json:"collector"`
	Parser     ParserConfig     `yaml:"parser" json:"parser"`
	Indexer    IndexerConfig    `yaml:"indexer" json:"indexer"`
	Webservice WebserviceConfig `yaml:"webservice" json:"webservice"`
}

type About struct {
	Version string `yaml:"version" json:"version"`
	Commit  string `yaml:"commit" json:"commit"`
	Date    string `yaml:"date" json:"date"`
}

type LogConfig struct {
	Level                  string `yaml:"level" json:"level"`
	EnableRuntimeLogfile   bool   `yaml:"enable_runtime_logfile" json:"enable_runtime_logfile"`
	EnableProcessedLogfile bool   `yaml:"enable_processed_logfile" json:"enable_processed_logfile"`
	RuntimeLogfilePath     string `yaml:"-" json:"runtime_logfile_path"`
	ProcessLogfilePath     string `yaml:"-" json:"process_logfile_path"`
}

type BehaviourConfig struct {
	DemoMode            bool `yaml:"demo_mode" json:"demo_mode"`
	AutostartWebservice bool `yaml:"autostart_webservice" json:"autostart_webservice"`
	OpenBrowserOnStart  bool `yaml:"open_browser_on_start" json:"open_browser_on_start"`
	ShowServiceGui      bool `yaml:"show_service_gui" json:"show_service_gui"`
}

type CollectorConfig struct {
	WorkerCount          int      `yaml:"worker_count" json:"worker_count"`
	WorkerDelayInMs      int      `yaml:"worker_delay_in_milliseconds" json:"worker_delay_in_milliseconds"`
	SearchablePaths      []string `yaml:"searchable_paths" json:"searchable_paths"`
	ExcludeSystemFolders bool     `yaml:"exclude_system_folders" json:"exclude_system_folders"`
}

type ParserConfig struct {
	WorkerCount     int `yaml:"worker_count" json:"worker_count"`
	WorkerDelayInMs int `yaml:"worker_delay_in_milliseconds" json:"worker_delay_in_milliseconds"`
}

type IndexerConfig struct {
	WorkerDelayInMs int `yaml:"worker_delay_in_milliseconds" json:"worker_delay_in_milliseconds"`
}

type WebserviceConfig struct {
	TryPorts      []int  `yaml:"try_ports" json:"-"`
	ChosenPort    int    `yaml:"-" json:"-"`
	OwnerPassword string `yaml:"owner_password" json:"owner_password"`
}
