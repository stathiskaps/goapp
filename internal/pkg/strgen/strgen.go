package strgen

import (
	"goapp/pkg/util"
	"log"
	"sync"
	"time"
)

type StringGenerator struct {
	strChan     chan<- string  // String output channel.
	quitChannel chan struct{}  // Quit.
	running     sync.WaitGroup // Running.
}

func New(strChan chan<- string) *StringGenerator {
	s := StringGenerator{}
	s.strChan = strChan
	s.quitChannel = make(chan struct{})
	s.running = sync.WaitGroup{}
	return &s
}

// Start string generator. Stop() must be called at the end.
func (s *StringGenerator) Start() error {
	s.running.Add(1)
	go s.mainLoop()

	return nil
}

func (s *StringGenerator) Stop() {
	close(s.quitChannel)
	s.running.Wait()
}

func (s *StringGenerator) mainLoop() {
	defer s.running.Done()

	for {
		hexStr, err := util.RandomHex(5)
		if err != nil {
			log.Printf("Error while generating random hex string (%v), skipping a beat\n", err)
			continue
		}
		select {
		case s.strChan <- hexStr:
		case <-s.quitChannel:
			return
		default:
		}
		time.Sleep(1 * time.Second)
	}
}
