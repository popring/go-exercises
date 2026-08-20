package todo

import (
	"fmt"
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
	nextID int
}

func (l *List) Add(title string) {
	l.nextID++
	t := Todo{
		ID:        l.nextID,
		Title:     title,
		CreatedAt: time.Now(),
	}
	l.Items = append(l.Items, t)
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
