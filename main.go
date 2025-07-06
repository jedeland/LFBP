package main

import (
	"example"
)
//
// The import path for a local package named "example" (with its own main.go file)
// located in the same directory as your main.go is just "example".
// If "example" is a subdirectory of your project root, use "example" as the import path.
// If you are using Go modules, the import path would be "your-module-name/example".
// For local development without modules, "example" works if the folder structure is:
// /home/jeeland/vscode/LFBP/
// ├── main.go
// └── example/
//     └── main.go
//
// So your import is correct as: import "example"
func main() {
	example.Run()
}