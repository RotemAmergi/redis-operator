package v1alpha2

const (
	defaultRedisNumber          = 3
	defaultSentinelNumber       = 3
	defaultExporterImage        = "oliver006/redis_exporter"
	defaultExporterImageVersion = "v0.11.3"
	defaultRedisImage           = "redis"
	defaultRedisImageVersion    = "5.0.4-alpine"
)

var (
	defaultSentinelCustomConfig = []string{"down-after-milliseconds 5000", "failover-timeout 10000"}
)
