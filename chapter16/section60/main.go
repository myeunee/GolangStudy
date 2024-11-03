package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/myeunee/GolangStudy/chapter16/section60/config"
	"golang.org/x/sync/errgroup"
)

// run 함수만 호출하도록
func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	s := &http.Server{
		// 인수로 받은 net.Listener를 이용 -> Addr 필드 지정 X
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	// 다른 고루틴에서 http 서버 실행
	eg.Go(func() error {
		// ListenAndServe 가 아니라 Serve 메서드로 변경
		if err := s.Serve(l); err != nil &&
			// http.ErrServerClosed는
			// http.Server.Shutdown()가 정상 종료되었다고 표시하므로 문제 X。
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
