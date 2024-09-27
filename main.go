package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structPractice/note"
	"example.com/structPractice/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func printSomething(value any) {

	typedVal, ok := value.(int)

	fmt.Println(typedVal)
	fmt.Println(ok)

	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case string:
		fmt.Println("String: ", value)
	default:
		fmt.Print("Other")
	}
}

func main() {
	printSomething(1)

	printSomething("Hooo")

	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func outputData(data outputable) error {

	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func getTodoData() string {
	return getUserInput("Todo text: ")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter title:")

	content := getUserInput("Enter content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
