JUICER=juicer
MERGE=$(JUICER) merge
GUARD=guard
ASSETS=./assets
PUBLIC=./public
JUICER_OPTS=-f -d $(PUBLIC)
CSS_OPTS=$(JUICER_OPTS)
CSS_EXTRA_FILES=
JS_OPTS=$(JUICER_OPTS) -s
JS_EXTRA_FILES=
IMG=$(PUBLIC)/img
IMG_PNG_OPTS=-o5
DEV_OPTS=
DEV?=0
GEM_OPTS=--no-ri --no-rdoc

ifeq ($(DEV), 1)
	DEV=1
	DEV_OPTS=-m none
endif

all: images assets

assets: js css

images: png

css:
	$(MERGE) $(CSS_EXTRA_FILES) $(ASSETS)/css/style.css \
	-o $(PUBLIC)/css/style.css -e data_uri $(CSS_OPTS) $(DEV_OPTS)

js:
	$(MERGE) $(JS_EXTRA_FILES) $(ASSETS)/js/app.js \
	-o $(PUBLIC)/js/app.js $(JS_OPTS) $(DEV_OPTS)

png:
	find $(IMG) -name '*.png' -exec optipng $(IMG_PNG_OPTS) {} \;

guard:
	$(GUARD) start -i

gems: install-guard install-juicer

install-juicer:
	gem install juicer $(GEM_OPTS)
	juicer install yui_compressor
	juicer install jslint

install-guard:
	gem install guard guard-shell guard-livereload $(GEM_OPTS)
