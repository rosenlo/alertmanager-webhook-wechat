include $(shell pwd)/build/Makefile

.PHONY:image
image:
	@cd $(WORKSPACE)/cmd/wecom-webhook/ && make VERSION=${VERSION}
	@docker build -t wecom-webhook:${VERSION} -f build/Dockerfile ./build

.PHONY:clean
clean:
	$(GOCLEAN)
	rm -rf $(BIN_PATH)
