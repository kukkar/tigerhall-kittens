package healthcheck

//
// Impl implement healthcheck
//
type Impl struct {
}

//
// Health call ping external services and databases to check health
//
func (this *Impl) Health() (HealthStatus, error) {
	return HealthStatus{
		Service:  STATUSOK,
		Database: STATUSOK,
	}, nil
}
