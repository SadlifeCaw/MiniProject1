package main

import (
	"sync"
)

type fork struct {
	sync.Mutex
	used      int
	beingUsed int
	id        int
}
