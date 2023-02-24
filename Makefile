PREFIX = $(HOME)/.local

all:
	go build .

install:
	go build .
	mkdir -p $(PREFIX)/bin
	mv -v oh-my-go $(PREFIX)/bin/oh-my-go