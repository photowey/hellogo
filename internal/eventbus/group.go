package eventbus

import (
	"sync"
)

type Group struct {
	subscribers SubscriberGroup
	lock        sync.RWMutex
}

func newGroup() *Group {
	return &Group{
		subscribers: make(SubscriberGroup, 0),
		lock:        sync.RWMutex{},
	}
}

func (group *Group) Length() int {
	return len(group.subscribers)
}

func (group *Group) Useable() bool {
	return group.Length() > 0
}

func (group *Group) remove(sub Subscriber) {
	length := len(group.subscribers)
	group.lock.Lock()
	defer group.lock.Unlock()
	idx := group.index(sub)
	if idx < 0 {
		return
	}
	copy(group.subscribers[idx:], group.subscribers[idx+1:])
	group.subscribers[length-1] = Subscriber{}
	group.subscribers = group.subscribers[:length-1]
}

func (group *Group) index(src Subscriber) int {
	for idx, sub := range group.subscribers {
		if sub == src {
			return idx
		}
	}

	return -1
}
