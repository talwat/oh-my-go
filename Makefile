PREFIX = $(HOME)/.local

all:
	go build .

install:
	go build .
	mv -v oh-my-go $(PREFIX)/bin/oh-my-go