GO=go
JAMMIT=jammit
BUNDLE=bundle
GEM_INSTALL=gem install --no-ri --no-rdoc
GUARD=$(BUNDLE) exec guard
PUBLIC_DIR?=public
SERVER=foreman start
DEV_OPTS=
DEV?=0

all: build precompile_assets

server: build
	$(SERVER)

build:
	$(GO) build .

test:
	$(GO) test ./...

precompile_assets:
	$(JAMMIT) -f -o $(PUBLIC_DIR)/assets

watch_assets:
	$(GUARD) start -i

deploy: all
	-git add $(PUBLIC_DIR) && git commit -qm "Recompiled assets"
	git push heroku master

prepare:
	$(GEM_INSTALL) bundler
	$(BUNDLE) install

help:
	@echo "Usage: make [TARGET] [VARS...]"
	@echo "The targets are: build, test, server, assets, guard, deploy, prepare"
	@echo "Defaults: DEV=0; PUBLIC_DIR=public;"
