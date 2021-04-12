
GOPKG ?=	moul.io/sgtm
DOCKER_IMAGE ?=	moul/sgtm
GOBINS ?=	./cmd/sgtm

PRE_INSTALL_STEPS += gen.sum
PRE_UNITTEST_STEPS += gen.sum
PRE_TEST_STEPS += gen.sum
PRE_BUILD_STEPS += gen.sum
PRE_LINT_STEPsS += gen.sum
PRE_TIDY_STEPS += gen.sum
PRE_BUMPDEPS_STEPS += gen.sum

include rules.mk

VCS_REF = `git rev-parse --short HEAD`
BUILD_DATE = `date +%s`
VERSION = `git describe --tags --always`

LDFLAGS ?= -X moul.io/sgtm/internal/sgtmversion.VcsRef=$(VCS_REF) -X moul.io/sgtm/internal/sgtmversion.Version=$(VERSION) -X moul.io/sgtm/internal/sgtmversion.BuildTime=$(BUILD_DATE)

COMPILEDAEMON_OPTIONS ?= -exclude-dir=.git -color=true -build=go\ install -build-dir=./cmd/sgtm
run: generate
	go install github.com/githubnemo/CompileDaemon
	CompileDaemon $(COMPILEDAEMON_OPTIONS) -command="sgtm --dev-mode --enable-server --enable-discord --enable-processing-worker run"
.PHONY: run

run-discord: generate
	go install github.com/githubnemo/CompileDaemon
	CompileDaemon $(COMPILEDAEMON_OPTIONS) -command="sgtm --dev-mode --enable-discord run"
.PHONY: run-discord

run-server: generate
	go install github.com/githubnemo/CompileDaemon
	CompileDaemon $(COMPILEDAEMON_OPTIONS) -command="sgtm --dev-mode --enable-server run"
.PHONY: run-server

run-processing-worker: generate
	go install github.com/githubnemo/CompileDaemon
	CompileDaemon $(COMPILEDAEMON_OPTIONS) -command="sgtm --dev-mode --enable-processing-worker run"
.PHONY: run-processing-worker

packr:
	(cd static; git clean -fxd)
	cd pkg/sgtm && packr2
.PHONY: packr

flushdb:
	rm -f /tmp/sgtm.db
.PHONY: flushdb

docker.push: tidy generate docker.build
	docker push $(DOCKER_IMAGE)
.PHONY: docker.push

# prod

PROD_HOST = zrwf.m.42.am
PROD_PATH = infra/projects/sgtm.club

prod.deploy.full: docker.push
	ssh $(PROD_HOST) make -C $(PROD_PATH) re
.PHONY: prod.deploy.full

prod.logs:
	ssh $(PROD_HOST) make -C $(PROD_PATH) logs
.PHONY: prod.logs

# FIXME: add deps
sgtm-linux-static:
	GOOS=linux GOARCH=amd64 $(GO) build -ldflags "-linkmode external -extldflags -static $(LDFLAGS)" -o $@ ./cmd/sgtm
.PHONY: sgtm-linux-static

prod.build: generate packr sgtm-linux-static
	rm -rf ./pkg/sgtm/packrd ./pkg/sgtm/sgtm-packr.go
	docker build -f Dockerfile.fast -t $(DOCKER_IMAGE) .
.PHONY: prod.build

prod.deploy: prod.build
	docker push $(DOCKER_IMAGE)
	ssh $(PROD_HOST) make -C $(PROD_PATH) re
.PHONY: prod.deploy

prod.syncdb:
	rsync -avze ssh $(PROD_HOST):$(PROD_PATH)/sgtm.db /tmp/sgtm.db
.PHONY: prod.syncdb

dbshare:
	@echo "== EASY SHARING WITH IPFS =="
	@ls -lh /tmp/sgtm.db
	@echo https://ipfs.io/ipfs/`ipfs add -q /tmp/sgtm.db`?filename=sgtm.sqlite3
	@echo
.PHONY: dbshare

prod.dbdump: