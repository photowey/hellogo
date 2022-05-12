package eventbus

import (
	"errors"
	"sync"
)

type EventBus struct {
	topicGroup map[string]*Group
	lock       sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		topicGroup: make(map[string]*Group),
		lock:       sync.RWMutex{},
	}
}

func (bus *EventBus) Length(topic string) (int, error) {
	bus.lock.Lock()
	if group, ok := bus.topicGroup[topic]; ok {
		bus.lock.Unlock()
		group.lock.RLock()
		defer group.lock.RUnlock()
		return group.Length(), nil
	} else {
		defer bus.lock.Unlock()
		return 0, errors.New("eventbus: topic not exist")
	}
}

func (bus *EventBus) Publish(topic string, data any) error {
	bus.lock.RLock()

	if group, ok := bus.topicGroup[topic]; ok {
		bus.lock.Unlock()
		group.lock.Lock()
		defer group.lock.Unlock()

		makeGroup := make(SubscriberGroup, 0)
		newGroupx := append(makeGroup, group.subscribers...)

		event := NewEvent(topic, data)
		go func(data Event, group SubscriberGroup) {
			for _, sub := range group {
				sub.receive(data)
			}
		}(event, newGroupx)

		return nil
	} else {
		defer bus.lock.RUnlock()
		return errors.New("eventbus: topic not exist")
	}
}

func (bus *EventBus) Subscribe(topic string, ch Channel) {
	bus.lock.Lock()
	sub := newSubscriber(ch)
	if oldGroup, ok := bus.topicGroup[topic]; ok {
		bus.lock.Unlock()
		oldGroup.lock.Lock()
		defer oldGroup.lock.Unlock()
		oldGroup.subscribers = append(oldGroup.subscribers, sub)
	} else {
		defer bus.lock.Unlock()
		group := newGroup()
		bus.topicGroup[topic] = group
		group.subscribers = append(group.subscribers, sub)
	}
}

func (bus *EventBus) UnSubscribe(topic string, sub Subscriber) {
	bus.lock.Lock()
	if group, ok := bus.topicGroup[topic]; ok && group.Useable() {
		bus.lock.Unlock()
		bus.topicGroup[topic].remove(sub)
	} else {
		defer bus.lock.Unlock()
		return
	}
}

type TopicPublisher func(data any) error

func (bus *EventBus) TopicFunc(topic string) TopicPublisher {
	return func(data any) error {
		return bus.Publish(topic, data)
	}
}
