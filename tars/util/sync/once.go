package sync

import (
	"sync"
	"sync/atomic"
)

// A Once must not be copied after first use.
type Once struct {
	done uint32
	m    sync.Mutex
}

func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 { // Outlined slow-path to allow inlining of the fast-path.
		return nil
	}
	return o.doSlow(f)
}

func (o *Once) doSlow(f func() error) (err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

func (o *Once) Done() bool {
	return atomic.LoadUint32(&o.done) == 1
}
