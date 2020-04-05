package employee

import "fmt"

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e *employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, e.TotalLeaves-e.LeavesTaken)
}

func New(firstname string, lastname string, totalLeaves int, leavesTaken int) *employee {
	newE := new(employee)
	newE.FirstName = firstname
	newE.LastName = lastname
	newE.TotalLeaves = totalLeaves
	newE.LeavesTaken = leavesTaken
	return newE
}
