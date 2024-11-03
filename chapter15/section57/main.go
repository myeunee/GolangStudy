package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

/******** section 55 main.go 참고 *******/
func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":80", // 포트를 코드 내부에 하드 코딩
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	// 다른 고루틴에서 http 서버 실행
	eg.Go(func() error {
		// http.ErrServerClosed는
		// http.Server.Shutdown() 가 정상 종료된 것을 나타냄 -> 이상 처리가 아님
		if err := s.ListenAndServe(); err != nil &&
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
