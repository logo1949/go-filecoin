package metrics

import (
	"context"

	"gx/ipfs/QmNVpHFt7QmabuVQyguf8AbkLDZoFh7ifBYztqijYT1Sd2/go.opencensus.io/stats"
	"gx/ipfs/QmNVpHFt7QmabuVQyguf8AbkLDZoFh7ifBYztqijYT1Sd2/go.opencensus.io/stats/view"
)

type Int64Counter struct {
	measureCt *stats.Int64Measure
	view      *view.View
}

func NewInt64Counter(name, desc string) *Int64Counter {
	log.Infof("registering int64 counter: %s - %s", name, desc)
	iMeasure := stats.Int64(name, desc, stats.UnitDimensionless)
	iView := &view.View{
		Name:        name,
		Measure:     iMeasure,
		Description: desc,
		Aggregation: view.Count(),
	}
	if err := view.Register(iView); err != nil {
		// a panic here indicates a developer error when creating a view.
		// Since this method is called in init() methods, this panic when hit
		// will cause running the program to fail immediately.
		panic(err)
	}

	return &Int64Counter{
		measureCt: iMeasure,
		view:      iView,
	}
}

func (c *Int64Counter) Inc(ctx context.Context, v int64) {
	stats.Record(ctx, c.measureCt.M(v))
}
