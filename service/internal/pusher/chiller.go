package pusher

import "time"

type ChannelChiller struct {
	msg     interface{}
	Channel chan interface{}
	to      chan<- interface{}
	ticker  *time.Ticker
}

func NewChanChiller(interval time.Duration, to chan<- interface{}) *ChannelChiller {
	v := &ChannelChiller{
		Channel: make(chan interface{}, 10000),
		to:      to,
		ticker:  time.NewTicker(interval),
	}

	go v.Chill()

	return v
}

func (c *ChannelChiller) Chill() {
	for {
		select {
		case v := <-c.Channel:
			c.msg = v
		case <-c.ticker.C:
			if c.msg == nil {
				continue
			}

			c.to <- c.msg
			c.msg = nil
		}
	}
}
