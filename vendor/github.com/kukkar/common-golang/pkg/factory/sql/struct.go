package sql

//
//Takesup the configuration for mysql connection.
//
type MysqlConfig struct {
	User               string `json:"User"`
	Password           string `json:"Password"`
	DBName             string `json:"DbName"`
	MaxOpenConnections int    `json:"MaxOpenConnections"`
	MaxIdleConnections int    `json:"MaxIdleConnections"`
	DefaultTimeZone    string `json:"DefaultTimeZone"`
	Host               string `json:"Host"`
	Port               string `json:"Port"`
}
