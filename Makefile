UDEL  ?= userdel
MKDIR ?= mkdir -p
CHOWN ?= chown
RM    ?= rm -rf
CP    ?= cp

.PHONY: build
build: clean
	go run cssMyKaomoji.go

.PHONY: clean
clean:
	rm kaomoji.css
	rm test/index.html
