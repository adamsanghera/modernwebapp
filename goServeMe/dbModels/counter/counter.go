package counter

import bus "github.com/adamsanghera/redisBus"

type Counter struct {
	Value int    `json:"Value"`
	ID    string `json:"ID"`
}

// NewCounter instantiates a Counter, given its ID.
func NewCounter(ID string) (*Counter, error) {
	// Create the new counter
	ctr := &Counter{
		ID: ID,
	}

	return ctr, bus.Client.SetNX(ctr.ID, 0, 0).Err()
}
