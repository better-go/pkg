module github.com/better-go/pkg/os

go 1.18

//replace github.com/better-go/pkg/log => ../log

require github.com/better-go/pkg/log v0.0.0-20220923023940-c922e8210ef0

require (
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
)
