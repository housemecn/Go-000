package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

// @Project: Go-000
// @Author: houseme
// @Description:
// @File: homework
// @Version: 1.0.0
// @Date: 2020/12/9 21:47
// @Package main

var ossi chan os.Signal

// server http.Server
type server struct {
	*http.Server

	name string
}

// handler handler http
type handler struct {
	name string
}

func main() {
	ossi = make(chan os.Signal, 1)
	group, ctx := errgroup.WithContext(context.Background())

	s1 := newServer("1", ":8088")
	s2 := newServer("2", ":9099")
	s3 := newServer("2", ":9098")

	group.Go(
		func() error {
			// If server is closed by close or shutdown, we don't return error so that err in waitGroup won't be polluted.
			if err := s1.ListenAndServe(); err != http.ErrServerClosed {
				return err
			}
			return nil
		},
	)
	group.Go(
		func() error {
			if err := s2.ListenAndServe(); err != http.ErrServerClosed {
				return err
			}
			return nil
		},
	)

	group.Go(
		func() error {
			if err := s3.ListenAndServe(); err != http.ErrServerClosed {
				return err
			}
			return nil
		},
	)

	group.Go(func() error {
		return listenSignal(ctx, s1, s2, s3)
	})

	if err := group.Wait(); err != nil {
		fmt.Println("1", err)
	}

}

// ServeHTTP Serve HTTP
func (h handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Server %s received request", h.name)
}

//  newServer new Server
func newServer(name, addr string) server {
	h := handler{name}

	return server{
		&http.Server{
			Addr:    addr,
			Handler: h,
		},
		name,
	}
}

// listenSignal listen Signal
func listenSignal(ctx context.Context, s1, s2, s3 server) error {
	signal.Notify(ossi)
	select {
	case sig := <-ossi:
		fmt.Printf("Recevied start signal %s\n", sig.String())
		// use context.Backgroud to shutdown the server because original ctx is canceled
		if err := s1.Shutdown(context.Background()); err != nil {
			fmt.Println("22", err)
		}
		if err := s2.Shutdown(context.Background()); err != nil {
			fmt.Println("33", err)
		}
		if err := s3.Shutdown(context.Background()); err != nil {
			fmt.Println("44", err)
		}
		return fmt.Errorf("Recevied signal %s", sig.String())
	case <-ctx.Done():
		return context.Canceled
	}
}
