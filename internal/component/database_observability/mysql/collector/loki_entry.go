package collector

import (
	"fmt"
	"time"

	"github.com/grafana/loki/v3/pkg/logproto"
	"github.com/prometheus/common/model"

	"github.com/grafana/alloy/internal/component/common/loki"
	"github.com/grafana/alloy/internal/component/database_observability"
	"github.com/grafana/alloy/internal/runtime/logging"
)

func buildLokiEntry(level logging.Level, op, instanceKey, line string) loki.Entry {
	return loki.Entry{
		Labels: model.LabelSet{
			"job":      database_observability.JobName,
			"op":       model.LabelValue(op),
			"instance": model.LabelValue(instanceKey),
		},
		Entry: logproto.Entry{
			Timestamp: time.Unix(0, time.Now().UnixNano()),
			Line:      fmt.Sprintf(`level="%s" %s`, level, line),
		},
	}
}
