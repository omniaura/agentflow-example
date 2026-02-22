# agentflow-example

A standalone example project that imports and uses [agentflow](https://github.com/omniaura/agentflow) to generate type-safe Go code from `.af` prompt templates.

## What This Demonstrates

- **Simple prompts** — No variables, just static text (`SystemPrompt`)
- **String variables** — Dynamic interpolation (`AnswerQuestion` with `<!topic>`, `<!question>`)
- **Optional blocks** — Conditional content with `<?context>...</context>`
- **Nested structs** — Dot-notation variables like `<!code.language>` generate nested Go structs
- **Typed variables** — `bool` and `int` types (`<!pr.is_draft bool>`, `<!pr.lines_changed int>`)
- **Conditional comparisons** — `<?pr.lines_changed gte 500>` for numeric thresholds

## Project Structure

```
prompts/
  assistant.af          # Coding assistant prompt templates
  assistant_af.go       # Generated Go code
  code-review.af        # Code review prompt templates
  code-review_af.go     # Generated Go code
main.go                 # Example usage that prints all prompts
```

## Usage

### Prerequisites

Install agentflow:
```bash
go install github.com/omniaura/agentflow/cmd/af@latest
```

### Generate Code

```bash
af gen prompts -d prompts
```

### Run

```bash
go run .
```

## Regenerating

If you modify the `.af` files, regenerate with:

```bash
af gen prompts -d prompts
go build ./...
```
