package models

import (
	"errors"
	"time"
	"sync"
)

type Task struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ClosedAt    *time.Time `json:"closedAt"`
}

type TaskLib struct {
	Lib    map[int]*Task
	mu     sync.RWMutex
	nextID int
}

func CreateContainer() *TaskLib {
	return &TaskLib{
		Lib:    make(map[int]*Task),
		nextID: 1,
	}
}

func (c *TaskLib) Create(title string, description string, priority string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newTask := Task{
		ID:          c.nextID,
		Title:       title,
		Description: description,
		Status:      "in progress",
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		ClosedAt:    nil,
	}
	c.Lib[c.nextID] = &newTask
	c.nextID++
}

func (c *TaskLib) Get(id int) (*Task, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if exTask, ex := c.Lib[id]; ex {
		return exTask, nil
	}
	return nil, errors.New("task is not exist")
}

func (c *TaskLib) Update(id int, title string, description string, priority string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if exTask, ex := c.Lib[id]; ex {
		exTask.Title = title
		exTask.Description = description
		exTask.Priority = priority
		exTask.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("task is not exist")
}

func (c *TaskLib) Delete(id int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ex := c.Lib[id]; ex {
		delete(c.Lib, id)
		return nil
	}
	return errors.New("task is not exist")
}

func (c *TaskLib) Patch(id int, changes map[string]string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if exTask, ex := c.Lib[id]; ex {
		for key, value := range changes {
			switch key {
			case "title":
				exTask.Title = value
			case "description":
				exTask.Description = value
			case "status":
				exTask.Status = value
			case "priority":
				exTask.Priority = value
			}
		}
		return nil
	}
	return errors.New("task is not exist")
}

func (c *TaskLib) End(id int) error {
	if exTask, ex := c.Lib[id]; ex {
		if exTask.ClosedAt != nil {
			return errors.New("task already closed")
		}
		closeTime := time.Now()
		exTask.ClosedAt = &closeTime
		exTask.UpdatedAt = closeTime
		exTask.Status = "Completed"
		return nil
	}
	return errors.New("task is not exist")
}