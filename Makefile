FILES := $(wildcard sample/*)

gobuild:
	@mkdir -p ./build
	@go build -ldflags "-X 'main.Version=$(shell git describe --tags)'" -o ./build/jdf

run:
	@make gobuild && ./build/jdf

run-sample:
	@make gobuild
	@$(foreach file, $(FILES), \
		echo "Printing contents of: $(file)" ; cat $(file) | ./build/jdf ; echo ;)

