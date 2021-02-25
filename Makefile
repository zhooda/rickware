GC=go

all: win windebug

debug: windebug

win: main.go
	GOOS=windows GOARSH=amd64 go build -ldflags -H=windowsgui
	# fyne package -os windows -icon icon.png

windebug: main.go
	GOOS=windows GOARCH=amd64 $(GC) build -o rickware_debug.exe