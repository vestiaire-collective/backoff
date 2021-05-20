package main

import (
	"time"

	"github.com/vestiaire-collective/backoff"
	"github.com/vestiaire-collective/backoff/expbo"
	"github.com/vestiaire-collective/backoff/linbo"
)

func main() {
	initial := uint64(1000)  //<< start at 1 sec (1000ms)
	max := uint64(1000 * 10) //<< cap it at 10 seconds
	exp := uint64(2)         //<< exponent
	inc := uint64(1000)      //<< increment 1 second

	expBO, err := expbo.New(initial, max, exp)

	if err != nil {
		panic(err)
	}

	linBO, err := linbo.New(initial, max, inc)

	if err != nil {
		panic(err)
	}

	flatBO, err := backoff.New(initial, initial, initial, nil)

	if err != nil {
		panic(err)
	}

	go func() {
		for {
			println("exponential backoff start...")

			expBO.Backoff()

			println("...exponential backoff stop")
		}
	}()

	go func() {
		for {
			println("linear backoff start...")

			linBO.Backoff()

			println("...linear backoff stop")
		}
	}()

	go func() {
		for {
			println("flat backoff start...")

			flatBO.Backoff()

			println("...flat backoff stop")
		}
	}()

	for {
		time.Sleep(time.Duration(10) * time.Second) //<< reset after 10 seconds

		expBO.Reset()
		linBO.Reset()
		flatBO.Reset()

		println("reset all")
	}
}
