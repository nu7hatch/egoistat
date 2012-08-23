GO=go
GIT=git
ECHO=echo
JAMMIT=jammit
BUNDLE=bundle
GEM_INSTALL=gem install --no-ri --no-rdoc
GUARD=$(BUNDLE) exec guard
PUBLIC_DIR?=public
COMPILED_ASSETS_DIR=$(PUBLIC_DIR)/assets
SERVER=foreman start
PRODUCTION_BRANCH=production
DEV_OPTS=
DEV?=0

_MSG_UNCOMMITED_CHANGES="Uncommited changes detected, commit or stash them before deploy!"

all: build precompile_assets

server: install
	$(SERVER)

install:
	$(GO) install .

build:
	$(GO) build .

test:
	$(GO) test ./...

precompile_assets:
	$(JAMMIT) -f -o $(COMPILED_ASSETS_DIR)

watch_assets:
	$(GUARD) start -i

deploy_require_clean_tree:
	@$(GIT) diff --quiet HEAD || ($(ECHO) $(_MSG_UNCOMMITED_CHANGES) && false)

deploy_merge_master:
	$(GIT) reset HEAD .
	$(GIT) checkout -f production
	$(GIT) merge master -q -m "Merged changes from master"

deploy: deploy_require_clean_tree deploy_merge_master build precompile_assets
	$(GIT) add $(COMPILED_ASSETS_DIR)
	$(GIT) commit -qm "Recompiled assets"
	$(GIT) push heroku '$(PRODUCTION_BRANCH):master'

prepare:
	$(GEM_INSTALL) bundler
	$(BUNDLE) install

help:
	@$(ECHO) "Usage: make [TARGET] [VARS...]"
	@$(ECHO) "The targets are: build, install, test, server, assets, guard, deploy, prepare"
	@$(ECHO) "Defaults: DEV=0; PUBLIC_DIR=public;"
