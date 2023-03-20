package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const worker = 30
const randomnumbertotal = 100

type Data struct {
	index       int
	value       int
	resultvalue int
}

func main() {
	timestart := time.Now()

	randomnum := createrandomnumber()
	randomnummulti2 := []int{}

	chandataout := make(chan Data)
	chandatain := make(chan Data)

	// Job Distribution
	go func() {
		for i := 0; i < len(randomnum); i++ {
			chandatain <- Data{
				index: i,
				value: randomnum[i],
			}
		}
		close(chandatain)
	}()

	// Worker Dispatcher
	wg := new(sync.WaitGroup)

	go func() {
		for i := 0; i <= worker; i++ {
			wg.Add(1)
			go func(workeridx int) {
				for job := range chandatain {
					calculateval := job.value * 2
					time.Sleep(500 * time.Millisecond)
					fmt.Println("worker:", workeridx, " calculate data index:", job.index, " value:", job.value)
					chandataout <- Data{
						resultvalue: calculateval,
					}
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(chandataout)
	}()

	// Result
	for res := range chandataout {
		randomnummulti2 = append(randomnummulti2, res.resultvalue)
	}

	fmt.Println(randomnummulti2)
	fmt.Println(time.Since(timestart))
}

func createrandomnumber() []int {
	res := []int{}

	for i := 0; i < randomnumbertotal; i++ {
		res = append(res, rand.Intn(100))
	}

	return res
}
