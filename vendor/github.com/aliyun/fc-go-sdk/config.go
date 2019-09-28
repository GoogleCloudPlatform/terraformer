package fc

// Config defines fc config
type Config struct {
	Endpoint        string // fc地址
	APIVersion      string // API版本
	AccountID       string // Account ID
	AccessKeyID     string // accessKeyID
	AccessKeySecret string // accessKeySecret
	SecurityToken   string // STS securityToken
	UserAgent       string // SDK名称/版本/系统信息
	IsDebug         bool   // 是否开启调试模式，默认false
	Timeout         uint   // 超时时间，默认60秒
	host            string // Set host from endpoint
}

// NewConfig get default config
func NewConfig() *Config {
	config := Config{}
	config.Endpoint = ""
	config.AccessKeyID = ""
	config.AccessKeySecret = ""
	config.SecurityToken = ""
	config.IsDebug = false
	config.UserAgent = "go-sdk-0.1"
	config.Timeout = 60
	config.APIVersion = "2016-08-15"
	config.host = ""
	return &config
}
