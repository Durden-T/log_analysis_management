package utils

import (
	"gin-vue-admin/global"
	"go.uber.org/zap"
	"time"
)

func SetInterval(duration time.Duration, f func() error) chan<- struct{} {
	done := make(chan struct{}, 1)
	go func() {
		timer := time.NewTicker(duration)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				err := f()
				if err != nil {
					global.GVA_LOG.Error("setInterval func error: ", zap.Error(err))
				}

			case <-done:
				return
			}
		}
	}()
	return done
}
