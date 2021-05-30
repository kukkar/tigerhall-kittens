package cache

type CacheConfig struct {
	Use   string `json:"Use"`
	Redis struct {
		Addr     string `json:"Addr"`
		PoolSize int    `json:"PoolSize"`
	} `json:"Redis"`
}
