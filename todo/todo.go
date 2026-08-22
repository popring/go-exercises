package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

type List struct {
	Items  []Todo
	NextID int
}

func (l *List) Add(title string) {
	l.NextID++
	t := Todo{
		ID:        l.NextID,
		Title:     title,
		CreatedAt: time.Now(),
	}
	l.Items = append(l.Items, t)
	fmt.Println("added")
}

func (l *List) List() {
	if len(l.Items) == 0 {
		fmt.Println("no todos")
		return
	}
	for _, t := range l.Items {
		mark := " "
		if t.Done {
			mark = "✅"
		}
		fmt.Printf("[%d] %s %s\n", t.ID, mark, t.Title)
	}
}

func (l *List) Complete(id int) error {
	for i, t := range l.Items {
		if t.ID == id {
			l.Items[i].Done = true
			return nil
		}
	}

	return fmt.Errorf("todo [%d] not found", id)
}

func (l *List) Delete(id int) error {
	for i, t := range l.Items {
		if t.ID == id {
			l.Items = append(l.Items[:i], l.Items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo [%d] not found", id)
}

var fileName = "todo.json"

func (l *List) Load() error {
	fileByte, err := os.ReadFile(fileName)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("file error", err)
		return err
	}
	if os.IsNotExist(err) {
		return nil
	}
	err = json.Unmarshal(fileByte, l)
	return err
}

func (l *List) Save() error {
	data, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}
	err = os.WriteFile(fileName, data, 0644)
	return err
}
