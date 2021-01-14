package notice

import "sync"

const DEFAULT = "default"

var (
	mt sync.Mutex
	keyOfINotice = make(map[string]INotice)
)

func Default() INotice {
	inst, ok := keyOfINotice[DEFAULT]
	if !ok {
		panic("default INotice is not initialize ")
	}

	return inst
}

func SetDefault(key string, inst INotice) {
	mt.Lock()
	defer mt.Unlock()
	if inst == nil {
		return
	}
	keyOfINotice[key] = inst
}