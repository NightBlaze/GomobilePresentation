package enginewidgets

import (
	"context"
	"errors"

	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/types"
)

type Task struct {
	id       int64
	cancelFn context.CancelFunc
	ctx      context.Context
}

func NewTask(ctx context.Context, cancelFn context.CancelFunc, _ struct{}) *Task {
	return &Task{
		cancelFn: cancelFn,
		ctx:      ctx,
	}
}

func TaskWithID(storage types.UniqueObjects[int64, *Task], id int64, _ struct{}) *Task {
	result, ok := storage.GetWithKey(id)
	if !ok {
		return nil
	}
	return result
}

func (t *Task) ID(_ struct{}) int64 {
	if t == nil {
		return TaskID_InternalError
	}
	return t.id
}

func (t *Task) SetID(id int64, _ struct{}) {
	if t == nil {
		return
	}
	t.id = id
}

func (t *Task) Cancel(_ struct{}) {
	if t == nil || t.cancelFn == nil {
		return
	}
	t.cancelFn()
}

func (t *Task) IsCancelled(_ struct{}) bool {
	if t == nil || t.ctx == nil {
		return false
	}

	err := t.ctx.Err()
	return err != nil && errors.Is(err, context.Canceled)
}
