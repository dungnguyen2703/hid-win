# HID Utility

HID Utility

go clean -modcache
go mod tidy
go get ...
go run ./main.go

# Create syso image

go run github.com/akavel/rsrc@latest -arch amd64 -ico icon.ico -o icon.syso

# Clean icon cache

taskkill /f /im explorer.exe
cd $env:localappdata
del IconCache.db -a
start explorer.exe

# Test

go test -v ./...

# Build

go build -o ../build/hid.exe -tags=release

go build -ldflags "-H windowsgui -s -w" -o ./build/hid.exe -tags=release
