package healthcheck

var _ HealthCheck = (*Impl)(nil)

//
// HealthCheck use to check service health using Health() method
//
type HealthCheck interface {
	Health() (HealthStatus, error)
}
