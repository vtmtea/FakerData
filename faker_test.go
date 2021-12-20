package faker

import (
	"fmt"
	"testing"
	"time"
)

func initialize() *Engine {
	faker := New()
	return faker
}

func TestName(t *testing.T) {
	faker := initialize()
	for i := 0; i < 10; i++ {
		fmt.Println(faker.Name())
		time.Sleep(time.Nanosecond)
	}
}

func TestEngine_PhoneNumber(t *testing.T) {
	faker := initialize()
	fmt.Println(faker.PhoneNumber())
}
