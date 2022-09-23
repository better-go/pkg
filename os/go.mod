module github.com/better-go/pkg/os

go 1.18

//replace github.com/better-go/pkg/log => ../log

require github.com/better-go/pkg/log v0.0.0-20220923014918-d5f0ad7ca40e

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
)
