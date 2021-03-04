GC=go

all: win windebug

debug: windebug

win: main.go
	rsrc -manifest test.manifest -ico icon.ico -o rsrc.syso
	GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui

windebug: main.go
	GOOS=windows GOARCH=amd64 $(GC) build -o rickware_debug.exe

clean:
	rm *.syso *.exe