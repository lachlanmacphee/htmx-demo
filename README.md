# Getting Started

1. Clone the repo to your local machine
2. Install dependencies (golang, templ cli, etc)
3. Run `templ generate`
4. Run `go run main.go`
5. Hit Ctrl+C
6. Comment out the `db.Seed()` line in `main.go`
7. Run `wgo run main.go` to get reloads on `.go` file changes
8. Run `templ generate --watch` in a separate terminal to reload on template changes
9. Add other demos as you see fit and create a PR if you want them merged
