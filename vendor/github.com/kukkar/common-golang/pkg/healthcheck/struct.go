package healthcheck

//
// HealthStatus represent response struct of health check
//
type HealthStatus struct {
	Service          string `json:"service"`
	Database         string `json:"DBHealthCheck"`
	ExternalServices string `json:"externalServices"`
}

//
// Config
//
type Config struct {
}
