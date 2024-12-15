package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("helllo world"))
	})
	ctx := context.Background()
	ctx = run(ctx)
	<-ctx.Done()
	fmt.Println("main exit...")
}
func run(ctx context.Context) context.Context {

	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":8080"

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("input any key to exit...")
		var input string
		fmt.Scanln(&input)
		fmt.Println("exit...")
		log.Println(srv.Shutdown(ctx))
		cancel()
	}()
	return ctx
}
