package metric

import (
	"github.com/glorydeath/nux"
	"log"
)

func SocketStatSummaryMetrics() (L []*MetricValue) {
	ssMap, err := nux.SocketStatSummary()
	if err != nil {
		log.Println(err)
		return
	}

	for k, v := range ssMap {
		L = append(L, GaugeValue("ss."+k, v))
	}

	return
}
