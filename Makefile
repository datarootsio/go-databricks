all: deps validate generate build test

build:
	go build ./...

deps:
	npm install -g autorest

generate:
	autorest --clusters
	autorest --dbfs
	autorest --groups
	autorest --workspace
	autorest --secrets
	autorest --libraries
	autorest --jobs

test:
	go test ./...

validate:
	autorest \
	  --input-file=clusters.yaml \
	  --input-file=dbfs.yaml \
	  --input-file=groups.yaml \
	  --input-file=workspace.yaml \
	  --input-file=secrets.yaml \
	  --input-file=libraries.yaml \
	  --input-file=jobs.yaml
