package counter

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

type Counter struct {
	mu       sync.Mutex
	filePath string
	value    int
}

// NewCounter creates a new Counter instance and initializes it with the value in the text file.
func NewCounter(filePath string) (*Counter, error) {
	counter := &Counter{filePath: filePath}

	err := counter.load()
	if err != nil {
		return nil, err
	}

	return counter, nil
}

// load reads the value from the text file.
func (c *Counter) load() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	file, err := os.Open(c.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			c.value = 0
			return nil // If the file doesn't exist, initialize the counter to 0.
		}
		return err
	}
	defer file.Close()

	var value int
	_, err = fmt.Fscanf(file, "%d", &value)
	if err != nil {
		return err
	}

	c.value = value
	return nil
}

// save writes the current value to the text file.
func (c *Counter) save() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	file, err := os.Create(c.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%d", c.value)
	if err != nil {
		return err
	}

	return nil
}

// Increment increases the counter value by 1 and saves it.
func (c *Counter) Increment() error {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
	return c.save()
}

// Decrement decreases the counter value by 1 and saves it.
func (c *Counter) Decrement() error {
	c.mu.Lock()
	c.value--
	c.mu.Unlock()
	return c.save()
}

// Value returns the current value of the counter.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
