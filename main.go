package main

import (
	"fmt"
	"strings"

	"github.com/omniaura/agentflow-example/prompts"
)

func main() {
	fmt.Println("=== AgentFlow Example ===")
	fmt.Println("Demonstrating generated prompts from 5 .af template files")
	fmt.Println()

	// --- assistant.af ---

	printSection("assistant.af", "System Prompt (no variables)")
	fmt.Println(&prompts.SystemPrompt{})

	printSection("assistant.af", "Answer Question (string variables)")
	fmt.Println(&prompts.AnswerQuestion{
		Topic:    "Go generics",
		Question: "How do I write a generic function that works with both slices and maps?",
	})

	printSection("assistant.af", "Answer Question (with optional context)")
	fmt.Println(&prompts.AnswerQuestion{
		Topic:    "Go generics",
		Question: "How do I write a generic function that works with both slices and maps?",
		Context:  "I'm using Go 1.21 and have read the tutorial on type parameters.",
	})

	printSection("assistant.af", "Summarize Conversation")
	fmt.Println(&prompts.SummarizeConversation{
		Messages: strings.Join([]string{
			"User: What is a goroutine?",
			"Assistant: A goroutine is a lightweight thread managed by the Go runtime.",
			"User: How do I communicate between goroutines?",
			"Assistant: Use channels to send and receive values between goroutines.",
		}, "\n"),
	})

	// --- code-review.af ---

	printSection("code-review.af", "Code Review (nested structs)")
	fmt.Println(&prompts.ReviewCode{
		Author: "alice",
		Code: struct {
			Language    string
			Filename    string
			Content     string
			Description string
		}{
			Language:    "go",
			Filename:    "handler.go",
			Content:     "func Handle(w http.ResponseWriter, r *http.Request) {\n    fmt.Fprintln(w, \"ok\")\n}",
			Description: "New HTTP handler",
		},
		Pr: struct {
			IsDraft      bool
			LinesChanged int
		}{
			IsDraft:      false,
			LinesChanged: 42,
		},
	})

	printSection("code-review.af", "Code Review (draft PR, large change triggers gte 500)")
	fmt.Println(&prompts.ReviewCode{
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
	})

	// --- commit-message.af ---

	printSection("commit-message.af", "Generate Commit Message (simple)")
	fmt.Println(&prompts.GenerateCommitMessage{
		Diff: "- func oldName() {}\n+ func newName() {}",
	})

	printSection("commit-message.af", "Generate Commit Message (with ticket)")
	fmt.Println(&prompts.GenerateCommitMessage{
		Diff:     "- return nil\n+ return fmt.Errorf(\"invalid input: %w\", err)",
		TicketId: "PROJ-1234",
	})

	// --- explain-error.af ---

	printSection("explain-error.af", "Error Explanation System Prompt")
	fmt.Println(&prompts.ErrorExplanationSystem{})

	printSection("explain-error.af", "Explain Error (minimal)")
	fmt.Println(&prompts.ExplainError{
		Language:     "Go",
		ErrorMessage: "panic: runtime error: index out of range [5] with length 3",
	})

	printSection("explain-error.af", "Explain Error (with stack trace, file, and line)")
	fmt.Println(&prompts.ExplainError{
		Language:     "Python",
		ErrorMessage: "TypeError: 'NoneType' object is not subscriptable",
		StackTrace:   "  File \"app.py\", line 42, in process\n    result = data[\"key\"]",
		File: struct {
			Name string
			Line int
		}{
			Name: "app.py",
			Line: 42,
		},
	})

	// --- test-generator.af ---

	printSection("test-generator.af", "Test Generation System Prompt")
	fmt.Println(&prompts.TestGenerationSystem{})

	printSection("test-generator.af", "Generate Tests (high coverage target, gte 90)")
	fmt.Println(&prompts.GenerateTests{
		TestFramework:  "pytest",
		Language:       "python",
		SourceCode:     "def fibonacci(n):\n    if n <= 1:\n        return n\n    return fibonacci(n-1) + fibonacci(n-2)",
		FunctionName:   "fibonacci",
		CoverageTarget: 95,
	})

	printSection("test-generator.af", "Generate Tests (low coverage target, lte 50)")
	fmt.Println(&prompts.GenerateTests{
		TestFramework:  "go test",
		Language:       "go",
		SourceCode:     "func ProcessBatch(items []Item) error { ... }",
		CoverageTarget: 40,
	})

	printSection("test-generator.af", "Generate Tests (with existing tests)")
	fmt.Println(&prompts.GenerateTests{
		TestFramework: "jest",
		Language:      "typescript",
		SourceCode:    "export function parseConfig(raw: string): Config { ... }",
		FunctionName:  "parseConfig",
		ExistingTests: "test('parses valid config', () => {\n  expect(parseConfig('{}')).toEqual({})\n})",
	})
}

func printSection(file, title string) {
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf(">>> [%s] %s\n", file, title)
	fmt.Println(strings.Repeat("-", 60))
}
