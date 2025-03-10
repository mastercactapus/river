package river

import (
	"context"

	"github.com/riverqueue/river/rivertype"
)

// HookDefaults should be embedded on any hook implementation. It helps
// guarantee forward compatibility in case additions are necessary to the Hook
// interface.
type HookDefaults struct{}

func (d *HookDefaults) IsHook() bool { return true }

// HookInsertBeginFunc is a convenience helper for implementing HookInsertBegin
// using a simple function instead of a struct.
type HookInsertBeginFunc func(ctx context.Context, params *rivertype.JobInsertParams) error

func (f HookInsertBeginFunc) InsertBegin(ctx context.Context, params *rivertype.JobInsertParams) error {
	return f(ctx, params)
}

func (f HookInsertBeginFunc) IsHook() bool { return true }

// HookWorkBeginFunc is a convenience helper for implementing HookworkBegin
// using a simple function instead of a struct.
type HookWorkBeginFunc func(ctx context.Context, job *rivertype.JobRow) error

func (f HookWorkBeginFunc) WorkBegin(ctx context.Context, job *rivertype.JobRow) error {
	return f(ctx, job)
}

func (f HookWorkBeginFunc) IsHook() bool { return true }
