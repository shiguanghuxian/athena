default:
	@echo 'Usage of make: [ build | clean ]'

build: 
	@go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD`" -o ./build/server ./apps/server

clean: 
	@rm -f ./build/monitorevent
	@rm -f ./build/logs/*.log

.PHONY: default build clean