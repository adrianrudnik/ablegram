package config

type Config struct {
	Version    int              `yaml:"version" json:"version"`
	About      About            `yaml:"-" json:"-"`
	Log        LogConfig        `yaml:"log" json:"log"`
	Behaviour  BehaviourConfig  `yaml:"behaviour" json:"behaviour"`
	Collector  CollectorConfig  `yaml:"collector" json:"collector"`
	Device     DeviceConfig     `yaml:"device" json:"device"`
	Indexer    IndexerConfig    `yaml:"indexer" json:"indexer"`
	Webservice WebserviceConfig `yaml:"webservice" json:"webservice"`

	IsDevelopmentEnv bool `yaml:"-" json:"is_dev_env"`
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
	Targets map[string]CollectorTarget `yaml:"targets" json:"targets"`
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

type DeviceConfig struct {
	Targets map[string]DeviceTarget `yaml:"targets" json:"targets"`
}

type DeviceTarget struct {
	ID   string `yaml:"id" json:"id"`
	Type string `yaml:"type" json:"type"`
	Uri  string `yaml:"uri" json:"uri"`
}

type IndexerConfig struct {
	WorkerDelayInMs int `yaml:"worker_delay_in_ms" json:"worker_delay_in_ms"`
}

type WebserviceConfig struct {
	TryPorts        []int  `yaml:"try_ports" json:"-"`
	ChosenPort      int    `yaml:"-" json:"-"`
	MasterPassword  string `yaml:"master_password" json:"master_password"`
	TrustedPlatform string `yaml:"trusted_platform" json:"trusted_platform"`
}
