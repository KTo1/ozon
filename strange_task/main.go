package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Sayer interface {
	Say()
}

type Base struct {
	Name string
}

type Child struct {
	Base
	LastName string
}

func (b *Base) Say() {
	fmt.Println("Hello,", b.Name)
}

func (c *Child) Say() {
	fmt.Println("Hello,", c.LastName, c.Name)
}

func NewObject(obj string, name, lastName string) Sayer {
	switch {
	case obj == "base":
		return &Base{Name: name}
	case obj == "child":
		return &Child{
			Base:     Base{Name: name},
			LastName: lastName,
		}
	default:
		return nil
	}
}

func Generator(ctx context.Context, n int) chan Sayer {
	result := make(chan Sayer)
	wg := sync.WaitGroup{}

	oneTimer := time.NewTicker(time.Second)
	twoTimer := time.NewTicker(2 * time.Second)

	wg.Add(2)
	go func() {
		defer wg.Done()

		for i := 0; i <= n; i++ {
			select {
			case <-ctx.Done():
				return
			case <-oneTimer.C:
				result <- NewObject("base", fmt.Sprintf("%v %v", "base", i), "")
			}
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i <= n; i++ {
			select {
			case <-ctx.Done():
			case <-twoTimer.C:
				result <- NewObject("child", fmt.Sprintf("%v %v", "child", i), "inherited")
			}
		}
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	b1 := Base{
		Name: "parent",
	}

	c1 := Child{
		Base:     Base{Name: "Child"},
		LastName: "Inherited",
	}

	b1.Say()
	c1.Say()

	c := []Sayer{&b1, &c1}

	for _, v := range c {
		v.Say()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Second)
	defer cancel()

	start := time.Now()
	ch := Generator(ctx, 5)

	for v := range ch {
		v.Say()
	}

	fmt.Println(time.Since(start))
}
