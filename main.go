package main

import (
	"fmt"
	"strings"

	"github.com/omniaura/agentflow-example/prompts"
)

func main() {
	fmt.Println("=== AgentFlow Example ===")
	fmt.Println()

	// 1. Simple prompt with no variables
	printSection("System Prompt (no variables)")
	sys := &prompts.SystemPrompt{}
	fmt.Println(sys)

	// 2. Prompt with simple string variables
	printSection("Answer Question (string variables)")
	q := &prompts.AnswerQuestion{
		Topic:    "Go generics",
		Question: "How do I write a generic function that works with both slices and maps?",
	}
	fmt.Println(q)

	// 3. Same prompt with optional context filled in
	printSection("Answer Question (with optional context)")
	qCtx := &prompts.AnswerQuestion{
		Topic:    "Go generics",
		Question: "How do I write a generic function that works with both slices and maps?",
		Context:  "I'm using Go 1.21 and have read the tutorial on type parameters.",
	}
	fmt.Println(qCtx)

	// 4. Prompt with conversation messages
	printSection("Summarize Conversation")
	messages := strings.Join([]string{
		"User: What is a goroutine?",
		"Assistant: A goroutine is a lightweight thread of execution managed by the Go runtime.",
		"User: How do I communicate between goroutines?",
		"Assistant: You can use channels to send and receive values between goroutines.",
	}, "\n")
	s := &prompts.SummarizeConversation{Messages: messages}
	fmt.Println(s)

	// 5. Code review with nested structs and typed variables
	printSection("Code Review (nested structs, typed variables)")
	review := &prompts.ReviewCode{
		Author: "alice",
		Code: struct {
			Language    string
			Filename    string
			Content     string
			Description string
		}{
			Language: "go",
			Filename: "handler.go",
			Content: `func HandleRequest(w http.ResponseWriter, r *http.Request) {
    data, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    fmt.Fprintf(w, "Received: %s", data)
}`,
			Description: "New HTTP handler for processing incoming requests",
		},
		Pr: struct {
			IsDraft      bool
			LinesChanged int
		}{
			IsDraft:      false,
			LinesChanged: 42,
		},
	}
	fmt.Println(review)

	// 6. Code review with draft flag and large change warning
	printSection("Code Review (draft PR, large change)")
	bigReview := &prompts.ReviewCode{
		Author: "bob",
		Code: struct {
			Language    string
			Filename    string
			Content     string
			Description string
		}{
			Language: "python",
			Filename: "refactor.py",
			Content:  "# ... 600 lines of refactored code ...",
		},
		Pr: struct {
			IsDraft      bool
			LinesChanged int
		}{
			IsDraft:      true,
			LinesChanged: 600,
		},
	}
	fmt.Println(bigReview)

	// 7. Code review system prompt
	printSection("Code Review System Prompt")
	crSys := &prompts.CodeReviewSystem{}
	fmt.Println(crSys)
}

func printSection(title string) {
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf(">>> %s\n", title)
	fmt.Println(strings.Repeat("-", 60))
}
