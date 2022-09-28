package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex

type Class1 struct {
	number float64
}

type Class2 struct {
	number float64
}

func (c1 *Class1) SetNumberClass1(number float64) {
	c1.number = number
}

func (c2 *Class2) SetNumberClass2(number float64) {
	c2.number = number
}

func (c1 Class1) GetNumberClass1() float64 {
	return c1.number
}

func (c2 Class2) GetNumberClass2() float64 {
	return c2.number
}

func f1(c1 *Class1, c2 *Class2, k1 int, wg *sync.WaitGroup) {
	mutex.Lock()
	for k := 0; k < k1; k++ {
		c1.SetNumberClass1(c1.GetNumberClass1() + rand.Float64())
		c2.SetNumberClass2(c2.GetNumberClass2() + rand.Float64())
	}
	mutex.Unlock()
	wg.Done()
}

func f2(c1 *Class1, c2 *Class2, k2 int, wg *sync.WaitGroup) {
	mutex.Lock()
	for k := 0; k < k2; k++ {
		c1.SetNumberClass1(c1.GetNumberClass1() + rand.Float64())
		c2.SetNumberClass2(c2.GetNumberClass2() + rand.Float64())
	}
	mutex.Unlock()
	wg.Done()
}

func main() {
	var c1 Class1
	var c2 Class2
	rand.Seed(time.Now().Unix())
	var wg sync.WaitGroup
	var numOfThreads = rand.Intn(10) + 10
	fmt.Println("Number of threads: ", numOfThreads)

	wg.Add(numOfThreads)

	var k1 = rand.Intn(10000) + 10000
	var k2 = rand.Intn(10000) + 10000
	fmt.Printf("к1 = %d\n", k1)
	fmt.Printf("к2 = %d\n", k2)

	for i := 0; i < numOfThreads/2; i++ {
		go f1(&c1, &c2, k1, &wg)
	}
	for i := 0; i < numOfThreads-numOfThreads/2; i++ {
		go f2(&c1, &c2, k2, &wg)
	}
	wg.Wait()
	fmt.Println("c1 =", c1.GetNumberClass1())
	fmt.Println("c2 =", c2.GetNumberClass2())
}
