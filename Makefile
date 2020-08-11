##########
# Building
##########

build-docker-dev-lint:
	docker build -f dev.lint.Dockerfile .
build-go:
	go get -v -t -d ./...
	go build -v .
	rm nuke

#########
# Linting
#########

lint-golangci:
	golangci-lint run
lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum
lint-goreleaser:
	goreleaser check
lint-hadolint:
	hadolint dev.lint.Dockerfile
lint-in-docker:
	docker build -f dev.lint.Dockerfile -t mattgleich/nuke:lint .
	docker run mattgleich/nuke:lint

##########
# Grouping
##########

# Linting
local-lint: lint-golangci lint-goreleaser lint-hadolint lint-gomod
docker-lint: lint-in-docker
# Build
local-build: build-docker-dev-lint
