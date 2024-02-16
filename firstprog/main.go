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

func main() {
	ba := booleanAttribute(false)
	ba.Render(context.Background(), os.Stdout)
	println()

	ca := conditionalAttribute(true)
	ca.Render(context.Background(), os.Stdout)
	println()

	sa := spreadAttributesUsage()
	sa.Render(context.Background(), os.Stdout)
	println()

	p1 := Person{
		URL:  "/github.com/lpuigo",
		Name: "Laurent",
	}
	uc := urlComponent(p1)
	uc.Render(context.Background(), os.Stdout)
	println()

	jsB := jsButton("du texte")
	jsB.Render(context.Background(), os.Stdout)
	println()
}
