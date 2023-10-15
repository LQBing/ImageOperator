@REM SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o bin/imageoperator-amd64-linux imageoperator.go

SET GOOS=darwin
SET GOARCH=amd64
go build -o bin/imageoperator-amd64-darwin imageoperator.go