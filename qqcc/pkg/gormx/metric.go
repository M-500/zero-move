package gormx

import "github.com/zeromicro/go-zero/core/metric"

// @Description
// @Author 代码小学生王木木

const gormNamespace = "gorm_client"

var (
	// 直方图
	metricClientReqDur = metric.NewHistogramVec(&metric.HistogramVecOpts{
		Namespace: gormNamespace,
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "gorm client requests duration(ms).",
		Labels:    []string{"table", "method"},
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	})

	// 只能增加 不能减少
	metricClientReqErrTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: gormNamespace,
		Subsystem: "requests",
		Name:      "error_total",
		Help:      "gorm client requests error count.",
		Labels:    []string{"table", "method", "is_error"},
	})
)
