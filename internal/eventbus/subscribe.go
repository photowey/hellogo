package eventbus

type SubscriberGroup []Subscriber

type Subscriber struct {
	channel Channel
}

func newSubscriber(channel chan Event) Subscriber {
	return Subscriber{
		channel: channel,
	}
}

func (sub *Subscriber) receive(event Event) {
	sub.channel <- event
}

func (sub *Subscriber) Await() (event Event) {
	return <-sub.channel
}
