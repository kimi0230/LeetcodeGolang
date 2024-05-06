all: build

build: summary
	gitbook build

start: summary
	gitbook serve .

summary:
	book sm -i _book node_modules 

deploy_firebase: summary build
	firebase login
	firebase deploy

tag: summary
	@echo "Version: $(VERSION)"
	git tag -a $(VERSION) -m "$(VERSION)"
	git push origin $(VERSION)

sync:
	git submodule update --recursive --remote

.PHONY: clean build all