// Copyright (C) 2017 ScyllaDB

package repair

import "errors"

// Repair errors
var (
	ErrActiveRepair = errors.New("repair already in progress")
	ErrDisabled     = errors.New("repair disabled")

	errAborted        = errors.New("repair aborted")
	errDoneWithErrors = errors.New("repair done with errors")
	errStopped        = errors.New("repair stopped")
)
