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

go build -ldflags "-H windowsgui -s -w" -o ./build/hid.exe -tags=release

# Notes

Debug builds (`go run ./main.go`, or `go build` without `-tags=release`) read `profile*.json` and `settings.json` from `./build`, and never touch the registry Run key — running from source will not register the app to start with Windows.
