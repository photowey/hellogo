package eventbus

type Channel chan Event

type Event struct {
	topic string
	data  any
}

func NewEvent(topic string, data any) Event {
	return Event{
		topic: topic, data: data,
	}
}
