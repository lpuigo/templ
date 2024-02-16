package main

import (
	"context"
	"os"
)

func main() {
	component := hello("Laurent")
	component.Render(context.Background(), os.Stdout)
}
