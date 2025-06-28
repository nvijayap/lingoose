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

	myThread := thread.New().AddMessage(
		thread.NewUserMessage().AddContent(
			thread.NewTextContent("Tell me a joke about geese"),
		),
	)

	err = openai.New().Generate(context.Background(), myThread)

	if err != nil {
		panic(err)
	}

	fmt.Println(myThread)
}
