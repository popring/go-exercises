package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/popring/go-todo-cli/todo"
)

func main() {
	list := &todo.List{}
	scanner := bufio.NewScanner(os.Stdin)

	err := list.Load()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	defer list.Save()

Loop:
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		args := strings.SplitN(line, " ", 2)
		cmd := args[0]
		arg2 := ""
		if len(args) >= 2 {
			arg2 = args[1]
		}

		switch cmd {
		case "add":
			if arg2 == "" {
				fmt.Println("title required")
				break
			}
			list.Add(arg2)
		case "list":
			list.List()
		case "done":
			id, err := strconv.Atoi(arg2)
			if err != nil {
				fmt.Printf("invalid id: %q\n", arg2)
				break
			}
			if err := list.Complete(id); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("[%d] done\n", id)
		case "del":
			id, err := strconv.Atoi(arg2)
			if err != nil {
				fmt.Printf("invalid id: %q\n", arg2)
				break
			}
			if err := list.Delete(id); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("[%d] deleted\n", id)
		case "quit":
			break Loop
		default:
			fmt.Println("unknown command:", cmd)
		}
	}
}
