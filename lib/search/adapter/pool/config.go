package pool

import "time"

type Config struct {
	Cap      int
	InitCap  int
	Lifetime time.Duration
	Timeout  time.Duration
	Factory  Creator
}
