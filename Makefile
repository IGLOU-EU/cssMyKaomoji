.PHONY: build
build: clean
	go run cssMyKaomoji.go

.PHONY: clean
clean:
	rm kaomoji.css
	rm test/index.html
