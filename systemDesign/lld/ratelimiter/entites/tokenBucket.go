package entites

import "sync"

type TokenBucket struct {
	Key       string
	Capacitiy int
	MU        sync.Mutex
}
