package main

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunWithoutErrors(t *testing.T) {
	var task1, task2, task3, task4 task
	task1 = func() error {
		fmt.Println("Task 1")

		return nil
	}

	task2 = func() error {
		fmt.Println("Task 2")

		return nil
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

	tasks := []task{task1, task2, task3, task4}
	tasksLength := len(tasks)

	tests := []struct {
		arr                  []task
		N                    int
		M                    int
		wantGoroutinesSucced int
		wantCompletedTasks   int
	}{
		{tasks, 4, 1, 4, tasksLength},
		{tasks, 3, 1, 3, tasksLength},
		{tasks, 2, 1, 2, tasksLength},
		{tasks, 1, 1, 1, tasksLength},
	}
	for _, tc := range tests {

		stat, err := Run(tc.arr, tc.N, tc.M)

		assert.Nil(t, err)
		assert.Equal(t, tc.wantGoroutinesSucced, stat.goroutinesSucced, "values should be equal")
		assert.Equal(t, tc.wantCompletedTasks, stat.completedTasks, "values should be equal")
	}
}

func TestRunWithErrors(t *testing.T) {
	var task1, task2, task3, task4, task5 task
	task1 = func() error {
		fmt.Println("Task 1")

		return errors.New("Error")
	}

	task2 = func() error {
		fmt.Println("Task 2")

		return errors.New("Error")
	}

	task3 = func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("Task 3")

		return nil
	}

	task4 = func() error {
		fmt.Println("Task 4")

		return nil
	}

	task5 = func() error {
		fmt.Println("Task 5")

		return nil
	}

	tasks := []task{task1, task2, task3, task4, task5}

	tests := []struct {
		arr            []task
		N              int
		M              int
		wantNoMoreThan int
	}{
		{tasks, 3, 2, 5},
		{tasks, 2, 2, 4},
		{tasks, 1, 2, 3},
		{tasks, 3, 1, 4},
		{tasks, 2, 1, 3},
		{tasks, 1, 1, 2},
	}
	for _, tc := range tests {

		stat, err := Run(tc.arr, tc.N, tc.M)

		assert.Nil(t, err)
		assert.True(t, stat.completedTasks <= tc.wantNoMoreThan, "should not be more than %v got %v", tc.wantNoMoreThan, stat.completedTasks)
	}
}
