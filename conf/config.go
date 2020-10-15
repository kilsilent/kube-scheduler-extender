package conf

var Conf *config

type config struct {
	PrometheusUrl             string
	PrometheusMemoryMetrics   string
	PrometheusMemoryThreshold int
}

func NewConfig(PrometheusUrl, PrometheusMemoryMetrics string, PrometheusMemoryThreshold int) {
	Conf = &config{
		PrometheusUrl:             PrometheusUrl,
		PrometheusMemoryMetrics:   PrometheusMemoryMetrics,
		PrometheusMemoryThreshold: PrometheusMemoryThreshold,
	}

}
