package main

import (
	"context"
	"fmt"
	"os"

	"github.com/henomis/lingoose/llm/openai"
	"github.com/henomis/lingoose/thread"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("godotenv.Load err:", err)
		return
	}

	os.Getenv("OPENAI_API_KEY")

	newUserMessage := thread.NewUserMessage()
	newTextContent := thread.NewTextContent("Tell me a joke about geese")
	message := newUserMessage.AddContent(newTextContent)

	myThread := thread.New()
	myThread.AddMessage(message)

	myOpenAI := openai.New()
	err = myOpenAI.Generate(context.Background(), myThread)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\nmyThread: %v\n\n", myThread)
}
