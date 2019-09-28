package sls

// const PluginInputType
const (
	PluginInputTypeDockerStdout = "service_docker_stdout"
	PPluginInputTypeCanal       = "service_canal"
)

type PluginInputItem struct {
	Type   string          `json:"type"`
	Detail PluginInterface `json:"detail"`
}

func CreatePluginInputItem(t string, detail PluginInterface) *PluginInputItem {
	return &PluginInputItem{
		Type:   t,
		Detail: detail,
	}
}

type LogConfigPluginInput struct {
	Inputs      []*PluginInputItem `json:"inputs"`
	Processors  []*PluginInputItem `json:"processors,omitempty"`
	Aggregators []*PluginInputItem `json:"aggregators,omitempty"`
	Flushers    []*PluginInputItem `json:"flushers,omitempty"`
}

type PluginInterface interface {
}

type ConfigPluginCanal struct {
	Host              string
	Port              int
	User              string
	Password          string
	Flavor            string
	ServerID          int
	IncludeTables     []string
	ExcludeTables     []string
	StartBinName      string
	StartBinLogPos    int
	HeartBeatPeriod   int
	ReadTimeout       int
	EnableDDL         bool
	EnableXID         bool
	EnableGTID        bool
	EnableInsert      bool
	EnableUpdate      bool
	EnableDelete      bool
	TextToString      bool
	StartFromBegining bool
	Charset           string
}

func CreateConfigPluginCanal() *ConfigPluginCanal {
	return &ConfigPluginCanal{
		Host:            "127.0.0.1",
		Port:            3306,
		User:            "root",
		Flavor:          "mysql",
		ServerID:        1205,
		HeartBeatPeriod: 60,
		ReadTimeout:     90,
		EnableGTID:      true,
		EnableInsert:    true,
		EnableUpdate:    true,
		EnableDelete:    true,
		Charset:         "utf8",
	}
}

type ConfigPluginDockerStdout struct {
	IncludeLabel         map[string]string
	ExcludeLabel         map[string]string
	IncludeEnv           map[string]string
	ExcludeEnv           map[string]string
	FlushIntervalMs      int
	TimeoutMs            int
	BeginLineRegex       string
	BeginLineTimeoutMs   int
	BeginLineCheckLength int
	MaxLogSize           int
	Stdout               bool
	Stderr               bool
}

func CreateConfigPluginDockerStdout() *ConfigPluginDockerStdout {
	return &ConfigPluginDockerStdout{
		FlushIntervalMs:      3000,
		TimeoutMs:            3000,
		Stdout:               true,
		Stderr:               true,
		BeginLineTimeoutMs:   3000,
		BeginLineCheckLength: 10 * 1024,
		MaxLogSize:           512 * 1024,
	}
}
