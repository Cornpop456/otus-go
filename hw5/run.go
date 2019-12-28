package main

import (
	"errors"
	"fmt"
	"time"
)

type task func() error

// RunStat contains stats of Run function
type RunStat struct {
	errors           int
	goroutinesSucced int
	completedTasks   int
}

func (r RunStat) String() string {
	return fmt.Sprintf("errors:%v\ngoroutines succed:%v\ntasks were done:%v\n", r.errors, r.goroutinesSucced, r.completedTasks)
}

func divideTasks(tasks []task, N int) [][]task {
	res := make([][]task, 0, N)
	partSize := (len(tasks) / N) + len(tasks)%N

	if len(tasks) <= N {
		for _, v := range tasks {
			res = append(res, []task{v})
		}
		return res
	}

	for i := 0; i < N; i++ {
		part := make([]task, 0, partSize)
		res = append(res, part)
	}

	for i := 0; i < len(tasks); i += len(res) {
		for j := range res {
			if i+j == len(tasks) {
				break
			}
			res[j] = append(res[j], tasks[i+j])
		}
	}

	return res
}

func runPart(arr []task, maxErrors int, finish chan<- int, errors chan<- error, currentErrorsNum <-chan int, completedTasks chan<- int) {
	errCount := 0
	tasksCount := 0
	for _, v := range arr {
		errCount = <-currentErrorsNum
		if errCount >= maxErrors {
			completedTasks <- tasksCount
			finish <- 0
			return
		}
		err := v()
		if err != nil {
			errors <- err
		}
		tasksCount++
	}
	completedTasks <- tasksCount
	finish <- 1
}

// Run executes tasks in N goroutines and finishes if M errors occured
func Run(tasks []task, N int, M int) (RunStat, error) {
	if N == 0 {
		return RunStat{}, errors.New("Can't run tasks in 0 goroutines")
	} else if M == 0 {
		return RunStat{}, errors.New("M should be greater than 0")
	}

	if N > len(tasks) {
		N = len(tasks)
	}

	errorCounter := 0
	finishedGoroutines := 0
	successedGoroutines := 0
	totalTasks := 0

	finish := make(chan int, N)
	errors := make(chan error, M)
	currentErrorsNum := make(chan int)
	completedTasks := make(chan int)

	arr := divideTasks(tasks, N)

	for _, v := range arr {
		go runPart(v, M, finish, errors, currentErrorsNum, completedTasks)
	}

	for finishedGoroutines < N {
		select {
		case tmpFin := <-finish:
			finishedGoroutines++
			successedGoroutines += tmpFin
		case <-errors:
			errorCounter++
		case tmpTasks := <-completedTasks:
			totalTasks += tmpTasks
		case currentErrorsNum <- errorCounter:
		}
	}

	close(finish)
	close(errors)
	close(currentErrorsNum)
	close(completedTasks)

	return RunStat{errorCounter, successedGoroutines, totalTasks}, nil
}

func main() {
	var task1, task2, task3, task4, task5, task6, task7, task8 func() error

	task1 = func() error {
		fmt.Println("Task 1")

		return errors.New("Error")
	}

	task2 = func() error {
		fmt.Println("Task 2")

		return errors.New("Error")
	}

	task3 = func() error {
		fmt.Println("Task 3")

		return nil
	}

	task4 = func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("Task 4")

		return nil
	}

	task5 = func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("Task 5")

		return nil
	}

	task6 = func() error {
		fmt.Println("Task 6")

		return nil
	}

	task7 = func() error {
		fmt.Println("Task 7")

		return nil
	}

	task8 = func() error {
		fmt.Println("Task 8")

		return nil
	}

	stat, _ := Run([]task{task1, task2, task3, task4, task5, task6, task7, task8}, 3, 2)

	fmt.Print(stat)
}
