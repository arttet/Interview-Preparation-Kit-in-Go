DOC_DIR ?= docs
DOC_PORT ?= 2000

## ▸▸▸ Documentation commands ◂◂◂

.PHONY: doc-build
doc-build:		## Build the documentation site [env: DOC_DIR=]
	zola --root ${DOC_DIR} build

.PHONY: doc-serve
doc-serve:		## Serve the documentation site [env: DOC_PORT=]
	zola --root ${DOC_DIR} serve --open

.PHONY: doc-clean
doc-clean:		## Remove generated artifacts [env: DOC_DIR=]
	rm -rf ${DOC_DIR}/public
