.PHONY: clean
clean:
	@rm -rf outline/ solid/

.PHONY: generate
generate: clean
	@go run .
	@templ fmt ./**
	@templ generate
