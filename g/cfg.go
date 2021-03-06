package g

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/toolkits/file"
)

type HttpConfig struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
}

type WebSocketConfig struct {
	Enabled bool   `json:"enabled"`
	Port    string `json:"port"`
}

type RedisConfig struct {
	Addr                  string `json:"addr"`
	Passwd                string `json:"password"`
	MaxIdle               int    `json:"maxIdle"`
	MigrationMapQueue     string `json:"migrationMapQueue"`
	MigrationCountQueue   string `json:"migrationCountQueue"`
	RelationMapQueue      string `json:"relationMapQueue"`
	RelationMapPointQueue string `json:"relationMapPointQueue"`
	EnableSentinel        bool   `json:"enableSentinel"`
}

type RedisSentinelConfig struct {
	SentinelAddrs []string `json:"sentinelAddrs"`
	Db            int      `json:"db"`
	MasterName    string   `json:"masterName"`
}

type ClusterNode struct {
	Addrs []string `json:"addrs"`
}

type KafkaConfig struct {
	Enabled     bool                    `json:"enabled"`
	Batch       int                     `json:"batch"`
	ConnTimeout int                     `json:"connTimeout"`
	CallTimeout int                     `json:"callTimeout"`
	MaxConns    int                     `json:"maxConns"`
	MaxIdle     int                     `json:"maxIdle"`
	Replicas    int                     `json:"replicas"`
	Topics      map[string]string       `json:"topics"`
	Cluster     map[string]string       `json:"cluster"`
	ClusterList map[string]*ClusterNode `json:"clusterList"`
}

type ImageConfig struct {
	FilePath string `json:"filepath"`
}

type GlobalConfig struct {
	Debug         bool                 `json:"debug"`
	MinStep       int                  `json:"minStep"` //最小周期,单位sec
	Hosts         string               `json:"hosts"`
	Database      string               `json:"database"`
	MaxConns      int                  `json:"maxConns"`
	MaxIdle       int                  `json:"maxIdle"`
	Listen        string               `json:"listen"`
	Trustable     []string             `json:"trustable"`
	Http          *HttpConfig          `json:"http"`
	WebSocket     *WebSocketConfig     `json:"websocket"`
	Download      string               `json:"download"`
	Redis         *RedisConfig         `json:"redis"`
	Kafka         *KafkaConfig         `json:"kafka"`
	RedisSentinel *RedisSentinelConfig `json:"redisSentinel"`
	Image         *ImageConfig         `json:"image"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	configLock.Lock()
	defer configLock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
}
