package prometheus

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	// Initial count.
	currentCount = 0

	// The Prometheus metric that will be exposed.
	httpHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ars2023_http_hit_total",
			Help: "Total number of http hits.",
		},
	)

	createConfigHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "create_config_http_hit_total",
			Help: "Total number of create config hits.",
		},
	)

	getAllConfigHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_all_config_http_hit_total",
			Help: "Total number of get all config hits.",
		},
	)

	getConfigHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_config_http_hit_total",
			Help: "Total number of get config hits.",
		},
	)

	delConfigHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "del_config_http_hit_total",
			Help: "Total number of del config hits.",
		},
	)

	createGroupHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "create_group_http_hit_total",
			Help: "Total number of create group hits.",
		},
	)

	getAllGroupHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_all_group_http_hit_total",
			Help: "Total number of get all group hits.",
		},
	)

	getGroupHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_group_http_hit_total",
			Help: "Total number of get group hits.",
		},
	)

	delGroupHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "del_group_http_hit_total",
			Help: "Total number of del group hits.",
		},
	)

	appendGroupHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "append_group_http_hit_total",
			Help: "Total number of append group hits.",
		},
	)

	getConfigByLabelsHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_config_by_labels_http_hit_total",
			Help: "Total number of get config by labels hits.",
		},
	)

	delConfigByLabelsHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "del_config_by_labels_http_hit_total",
			Help: "Total number of del config by labels hits.",
		},
	)

	// Add all metrics that will be resisted
	metricsList = []prometheus.Collector{
		httpHits,
		createConfigHits,
		getAllConfigHits,
		getConfigHits,
		delConfigHits,
		createGroupHits,
		getAllGroupHits,
		getGroupHits,
		delGroupHits,
		appendGroupHits,
		getConfigByLabelsHits,
		delConfigByLabelsHits,
	}

	// Prometheus Registry to register metrics.
	prometheusRegistry = prometheus.NewRegistry()
)

func init() {
	// Register metrics that will be exposed.
	prometheusRegistry.MustRegister(metricsList...)
}

func MetricsHandler() http.Handler {
	return promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{})
}

func Count(f func(context.Context, http.ResponseWriter, *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		f(ctx, w, r) // original function call
	}
}

func CountCreateConfig(f func(ctx context.Context, w http.ResponseWriter, req *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		createConfigHits.Inc()
		f(ctx, w, r) // original function call
	}
}

func CountGetAllConfig(f func(ctx context.Context, w http.ResponseWriter, req *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		getAllConfigHits.Inc()
		f(ctx, w, r) // original function call
	}
}

func CountGetConfig(f func(ctx context.Context, w http.ResponseWriter, req *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		getConfigHits.Inc()
		f(ctx, w, r) // original function call
	}
}

func CountDelConfig(f func(ctx context.Context, w http.ResponseWriter, req *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		delConfigHits.Inc()
		f(ctx, w, r) // original function call
	}
}

func CountCreateGroup(f func(ctx context.Context, w http.ResponseWriter, req *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		createGroupHits.Inc()
		f(ctx, w, r) // original function call
	}
}

func CountGetAllGroup(f func(ctx context.Context, w http.ResponseWriter, req *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		getAllGroupHits.Inc()
		f(ctx, w, r) // original function call
	}
}

func CountGetGroup(f func(ctx context.Context, w http.ResponseWriter, req *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		getGroupHits.Inc()
		f(ctx, w, r) // original function call
	}
}

func CountDelGroup(f func(ctx context.Context, w http.ResponseWriter, req *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		delGroupHits.Inc()
		f(ctx, w, r) // original function call
	}
}

func CountAppendGroup(f func(ctx context.Context, w http.ResponseWriter, req *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		appendGroupHits.Inc()
		f(ctx, w, r) // original function call
	}
}

func CountGetConfigByLabels(f func(ctx context.Context, w http.ResponseWriter, req *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		getConfigByLabelsHits.Inc()
		f(ctx, w, r) // original function call
	}
}

func CountDelConfigByLabels(f func(context.Context, http.ResponseWriter, *http.Request)) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		delConfigByLabelsHits.Inc()
		f(ctx, w, r) // original function call
	}
}
