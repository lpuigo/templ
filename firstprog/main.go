package main

import (
	"context"
	"os"
)

//func main() {
//	component := hello("Laurent")
//	http.Handle("/", templ.Handler(component))
//
//	fmt.Println("Listening on :3000")
//	http.ListenAndServe(":3000", nil)
//}

// To install templ package (if not already available locally)
//
// 		go install github.com/a-h/templ/cmd/templ@latest
//
// Add to current project
//
//		go get github.com/a-h/templ

func main() {
	//ba := booleanAttribute(false)
	//ba.Render(context.Background(), os.Stdout)
	//println()
	//
	//ca := conditionalAttribute(true)
	//ca.Render(context.Background(), os.Stdout)
	//println()
	//
	//sa := spreadAttributesUsage()
	//sa.Render(context.Background(), os.Stdout)
	//println()
	//
	//p1 := Person{
	//	URL:  "/github.com/lpuigo",
	//	Name: "Laurent",
	//}
	//uc := urlComponent(p1)
	//uc.Render(context.Background(), os.Stdout)
	//println()
	//
	//jsB := jsButton("du texte")
	//jsB.Render(context.Background(), os.Stdout)
	//println()

	//dynCssButton("Click me false", false).Render(context.Background(), os.Stdout)
	//println()
	//dynCssButton("Click me1 true", true).Render(context.Background(), os.Stdout)
	//println()
	//dynCssButton("Click me2 true", true).Render(context.Background(), os.Stdout)
	//println()

	cssCompButton("Click me2 false", false).Render(context.Background(), os.Stdout)
	println()
	cssCompButton("Click me2 true", true).Render(context.Background(), os.Stdout)
	println()
}

// continue with https://templ.guide/core-concepts/components
