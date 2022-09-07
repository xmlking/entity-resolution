lint:
	@echo "format and lint..."
	@buf format -w --exit-code
	@buf lint

generate: lint
	@echo "generate..."
	@rm -rf gen
	@buf generate

publish: generate
	@echo "publish to BSR..."
	@buf mod update proto/entityapis
	@buf push proto/entityapis

release: publish
	@git add .
	@git commit -am "build(proto): proto mod files updated"
	@cog bump --auto
