PREFIX = $(HOME)/.local

build:
	mkdir -vp build
	go build -o build

install: build/oh-my-go
	mkdir -p $(PREFIX)/bin
	mv -v build/oh-my-go $(PREFIX)/bin/oh-my-go

clean:
	rm -rvf build