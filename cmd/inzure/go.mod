module github.com/CarveSystems/inzure/cmd/inzure

go 1.12

replace github.com/CarveSystems/inzure/pkg/inzure => ../../pkg/inzure

require (
	github.com/CarveSystems/inzure/pkg/inzure v0.0.1
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/urfave/cli v1.22.5
)
