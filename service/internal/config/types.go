package config

type Config struct {
	About      About            `yaml:"-" json:"-"`
	Log        LogConfig        `yaml:"log" json:"log"`
	Behaviour  BehaviourConfig  `yaml:"behaviour" json:"behaviour"`
	Collector  CollectorConfig  `yaml:"collector" json:"collector"`
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
	Targets []CollectorTarget `yaml:"targets" json:"targets"`
}

type IndexerConfig struct {
	WorkerDelayInMs int `yaml:"worker_delay_in_ms" json:"worker_delay_in_ms"`
}

type WebserviceConfig struct {
	TryPorts      []int  `yaml:"try_ports" json:"-"`
	ChosenPort    int    `yaml:"-" json:"-"`
	OwnerPassword string `yaml:"owner_password" json:"owner_password"`
}

type CollectorTarget struct {
	ID                   string `yaml:"id" json:"id"`
	Type                 string `yaml:"type" json:"type"`
	Uri                  string `yaml:"uri" json:"uri"`
	ParserPerformance    string `yaml:"parser_performance" json:"parser_performance"`
	ParserWorkerDelay    int    `yaml:"parser_delay" json:"parser_delay"`
	ExcludeSystemFolders bool   `yaml:"exclude_system_folders" json:"exclude_system_folders"`
	ExcludeDotFolders    bool   `yaml:"exclude_dot_folders" json:"exclude_dot_folders"`
}
