// Package os runs processes locally
package os

import (
	"github.com/focalsolution/micro-go-micro/runtime/process"
)

type Process struct{}

func NewProcess(opts ...process.Option) process.Process {
	return &Process{}
}
