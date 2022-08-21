CURRENT_DIR=$(shell pwd)

$(CURRENT_DIR)/lib/tinycc/libtcc.a: submodule_update
	cd $(CURRENT_DIR)/lib/tinycc; \
	grep -rl CString . | xargs sed -i 's/CString/tinycc_CString/g' \
	./configure; \
	make

submodule_update:
	@echo "Cloning https://github.com/TinyCC/tinycc.git into $(CURRENT_DIR)/lib/ ..."
	git submodule update --init

clean:
	cd $(CURRENT_DIR)/lib/tinycc; \
	git clean -fxd

build:
	go build 

generate:
	go generate

test:
	go test

.PHONY: submodule_update clean build generate test
