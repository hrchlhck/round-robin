package main

import "fmt"

type Process struct {
	WaitTime      int
	ArrivalTime   int
	BurstTime     int
	ID            int
	RemainingTime int
}

func createProcess(ID, arrivalTime, burstTime int) *Process {
	return &Process{
		ID:            ID,
		ArrivalTime:   arrivalTime,
		BurstTime:     burstTime,
		WaitTime:      0,
		RemainingTime: burstTime,
	}
}

func (p *Process) String() string {
	return fmt.Sprintf("<PID %d | ArrivalTime %d | BurstTime %d | WaitTime %d | Remaining Time %d>", p.ID, p.ArrivalTime, p.BurstTime, p.WaitTime, p.RemainingTime)
}
