package timelocation

import (
	"fmt"
	"sync"
	"time"
)

var (
	once                 sync.Once
	washingtonDCLocation *time.Location
)

func Setup() {
	once.Do(func() {
		var err error

		washingtonDCLocation, err = time.LoadLocation("America/New_York")
		if err != nil {
			panic(fmt.Errorf("无法加载美国华盛顿特区时区, err: %w", err))
		}
	})
}

func GetWashingtonDCLocation() *time.Location {
	return washingtonDCLocation
}
