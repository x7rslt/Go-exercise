package main

import (
	"context"
	"fmt"
)

func main() {
	type favContextkey string
	f := func(ctx context.Context, k favContextkey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		} else {
			fmt.Println("not found value.")
		}
	}
	k := favContextkey("language")
	ctx := context.WithValue(context.Background(), k, "Go")
	//f(ctx, k)
	f(ctx, favContextkey("language"))
}
