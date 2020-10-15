package conf

var Conf = &config{}

type config struct {
	PrometheusUrl             string
	PrometheusMemoryMetrics   string
	PrometheusMemoryThreshold int
}
