gobuild:
	@mkdir -p ./build
	@go build -o ./build/jdf

run:
	@make gobuild && ./build/jdf

test:
	@make gobuild
	@echo "Running tests..."
	@for file in tests/*; do \
		if [ -f "$$file" ]; then \
			cat "$$file" | ./build/jdf; \
		fi \
	done
