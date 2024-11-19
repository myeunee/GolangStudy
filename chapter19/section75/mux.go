package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/myeunee/GolangStudy/chapter19/section75/clock"
	"github.com/myeunee/GolangStudy/chapter19/section75/config"
	"github.com/myeunee/GolangStudy/chapter19/section75/handler"
	"github.com/myeunee/GolangStudy/chapter19/section75/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	// chi.Router의 HandlerFunc은 라우팅 패턴과, 핸들러 함수를 받음
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))

	})
	fmt.Print("핸들러 등록 성공")
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg) // connection 생성
	if err != nil {
		return nil, cleanup, err
	}
	r := store.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{DB: db, Repo: &r, Validator: v} // AddTask를 handler로 사용
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{DB: db, Repo: &r}
	mux.Get("/tasks", lt.ServeHTTP)
	return mux, cleanup, nil
}
