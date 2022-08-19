package main

import "fmt"

type RoundRobin struct {
	Processes []*Process
	Counter   int
	Quantum   int
	WaitTime  int
	Switches  int
}

func createRoundRobin(quantum int) *RoundRobin {
	return &RoundRobin{
		Processes: make([]*Process, 0),
		Counter:   0,
		Quantum:   quantum,
		WaitTime:  0,
		Switches:  0,
	}
}

func (rr *RoundRobin) Add(p *Process) {
	rr.Processes = append(rr.Processes, p)
}

func (rr *RoundRobin) Remove() {
	if len(rr.Processes) > 0 {
		fmt.Printf("Removing process %d\n", rr.Processes[0].ID)
		rr.Processes = rr.Processes[1:]
		rr.Counter--
	}
}

func (rr *RoundRobin) getCounter() int {
	if rr.Counter >= len(rr.Processes) {
		rr.Counter = 0
	}

	return rr.Counter
}

func (rr *RoundRobin) Head() *Process {
	p := rr.Processes[rr.getCounter()]
	rr.Counter++
	return p
}

func (rr *RoundRobin) Print() {
	for _, p := range rr.Processes {
		fmt.Println(p.String())
	}
}

func (rr *RoundRobin) ShowMetrics() {
	fmt.Printf("Total wait time: %d units\n", rr.WaitTime)
	fmt.Printf("Total switches: %d\n", rr.Switches)
}

func (rr *RoundRobin) IsEmpty() bool {
	return len(rr.Processes) == 0
}

func (rr *RoundRobin) UpdateWaitTime(ignore *Process) {
	for _, p := range rr.Processes {
		if p.ID != ignore.ID {
			p.WaitTime++

			rr.WaitTime++
		}
	}
}
