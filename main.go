package main

import ("github.com/LuvArora443/MicroservicesInGo/handlers"
	"log"
	"os"
	"net/http"
	"time"
	"context"
	"os/signal"
	"syscall"
)


func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	go func(){
		l.Println("Server running on 9090")
		err:= s.ListenAndServe()
		if err!=nil{
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}