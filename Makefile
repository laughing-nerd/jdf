gobuild:
	@mkdir -p ./build
	@go build -o ./build/jdf

run:
	@make gobuild && ./build/jdf
