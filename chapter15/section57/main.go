package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}
	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", p, err)
	}
	// l.Addr() 메서드로 네트워크 연결의 주소를 가져옴
	url := fmt.Sprintf("http://%s", l.Addr().String())

	// 주소를 로그로 출력
	log.Printf("start with: %v", url)

	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, l net.Listener) error {
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	// 다른 고루틴에서 http 서버 실행
	eg.Go(func() error {
		if err := s.Serve(l); err != nil &&
			// http.ErrServerClosed는
			// http.Server.Shutdown() 가 정상 종료된 것을 나타냄 -> 이상 처리가 아님
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})
	// 채널로부터 알림(종료 알림)을 기다림
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// Go메서드로 실행한 다른 고루틴의 종료를 기다림
	return eg.Wait()
}
