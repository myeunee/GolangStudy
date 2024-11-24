package service

import (
	"context"

	"github.com/myeunee/GolangStudy/chapter19/section75/entity"
	"github.com/myeunee/GolangStudy/chapter19/section75/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister
type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}
type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}
