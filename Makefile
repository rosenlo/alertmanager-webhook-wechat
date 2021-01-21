include $(shell pwd)/build/Makefile

.PHONY:image
image:
	@cd $(shell pwd)/cmd/webhook-wechat/ && make VERSION=${VERSION}
	@docker build -t rosenpy/alertmanager-webhook-wechat:${VERSION} -f build/Dockerfile ./build

.PHONY:clean
clean:
	$(GOCLEAN)
	rm -rf $(BIN_PATH)
