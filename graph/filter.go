package graph

type SensorReading struct {
	Name    string
	Value   int
	Quality int
}

type DAQFilter struct {
	Subscribers map[int64]*DAQSubscriber
}

type DAQSubscriber struct {
	SubChan chan SensorReading
	Filter  func(machine string, sensor string, value int, quality int) bool
}

// There is a go routine here, that listens to the sniffer receive channel, and
// forwards to subscribers.

func (f *DAQFilter) Start(udpStream chan []SensorReading) {
	for {
		v := <-udpStream
		for _, sub := range f.Subscribers {
			for i := range v {
				if sub.Filter(v[i].Name, v[i].Value, v[i].Quality) {
					sub.SubChan <- v[i]
				}
			}
		}
	}
}
