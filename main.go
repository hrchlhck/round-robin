package main

import (
	"fmt"
	"time"
)

func main() {
	var rr *RoundRobin = createRoundRobin(1)

	rr.Add(createProcess(0, 0, 1))
	rr.Add(createProcess(1, 0, 2))
	rr.Add(createProcess(2, 0, 4))
	rr.Add(createProcess(3, 0, 6))
	rr.Add(createProcess(4, 0, 8))

	var lastP *Process = nil
	for !rr.IsEmpty() {
		var p *Process = rr.Head()

		if lastP == nil {
			lastP = p
		}

		for i := 0; i < rr.Quantum; i++ {
			rr.UpdateWaitTime(p)
			time.Sleep(1 * time.Second)

			p.RemainingTime--
			fmt.Printf("[ P%d ] Remaining time %d\n", p.ID, p.RemainingTime)

			if p.RemainingTime <= 0 {
				rr.Remove()
				break
			}
		}
		if lastP.ID != p.ID {
			rr.Switches++
			lastP = p
		}
	}
	rr.ShowMetrics()
}
