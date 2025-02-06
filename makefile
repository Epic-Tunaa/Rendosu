BUILD_FLAGS = -Idflags="-s -w" -trimpath
OUTPUT = rendosu

all: build

build:
	go build $(BUILD_FLAGS) -o $(OUTPUT) ./cmd.main.go

release:
	GOOS=windows GOARCH=and64 go build $(BUILD_FLAGS) -o $(OUTPUT)-win.exe
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o $(OUTPUT)-linux

install: build
	cp $(OUTPUT) /user/local/bin

clean:
	rm -f $(OUTPUT)