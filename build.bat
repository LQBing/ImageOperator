@REM SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o bin/ImageOperator_linux_amd64 imageoperator.go

SET GOOS=darwin
SET GOARCH=amd64
go build -o bin/ImageOperator_darwin_amd64 imageoperator.go

SET GOOS=windows
SET GOARCH=amd64
go build -o bin/ImageOperator_windows_amd64.exe imageoperator.go