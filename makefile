VERSION ?= v0.0.1

all: build

install:
	pip install -r requirements.txt
	npm install staticrypt

build:
	mkdocs build

pdf:
	ENABLE_PDF_EXPORT=1 mkdocs build
	@echo "PDF generated at site/pdf/leetcode-golang.pdf"

start:
	mkdocs serve

# 使用方式: make gen ID=217
gen:
	go run cmd/leetcode-readme-gen.go $(ID)

tag:
	@echo "Version: $(VERSION)"
	git tag -a $(VERSION) -m "$(VERSION)"
	git push origin $(VERSION)

sync:
	git submodule update --recursive --remote

clean:
	rm -rf site/

.PHONY: clean build all start gen tag sync
