FILES := $(wildcard sample/*)

gobuild:
	@mkdir -p ./build
	@CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o ./build/jdf .

run:
	@make gobuild && ./build/jdf

run-sample:
	@make gobuild
	@$(foreach file, $(FILES), \
		echo "Printing contents of: $(file)" ; cat $(file) | ./build/jdf ; echo ;)

