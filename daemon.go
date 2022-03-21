package gtavd

import (
	"context"

	"github.com/kardianos/service"
	"github.com/maestre3d/gtavd/dlclist"
)

type Daemon struct {
	baseCtx       context.Context
	baseCtxCancel context.CancelFunc
}

var _ service.Interface = &Daemon{}

func (d *Daemon) Start(_ service.Service) error {
	d.baseCtx, d.baseCtxCancel = context.WithCancel(context.Background())
	go dlclist.Watch(d.baseCtx)
	return nil
}

func (d *Daemon) Stop(_ service.Service) error {
	d.baseCtxCancel()
	return nil
}
