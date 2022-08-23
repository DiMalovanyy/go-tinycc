CURRENT_DIR=$(shell pwd)
TCC_LIBC_NAME=""
TCC_INCLUDE_DIR=$(CURRENT_DIR)/lib/tinycc/include

# TODO: 
# 1. Locate TccLib if it already installed 			

$(TCC_INCLUDE_DIR)/libtcc.a: submodule_update
	cd $(CURRENT_DIR)/lib/tinycc; \
	grep -rl CString . | xargs sed -i 's/CString/tinycc_CString/g'; \
	./configure; \
	make

submodule_update:
	@echo "Cloning https://github.com/TinyCC/tinycc.git into $(CURRENT_DIR)/lib/ ..."
	git submodule update --init

build:
	go build -ldflags \
		"-X 'github.com/DiMalovanyy/go-tinycc/internal.TccHeaders=$(TCC_INCLUDE_DIR)'"

test:
	go test -ldflags \
		"-X 'github.com/DiMalovanyy/go-tinycc/internal.TccHeaders=$(TCC_INCLUDE_DIR)'"

clean:
	cd $(CURRENT_DIR)/lib/tinycc; \
	git clean -fxd; \
	git checkout -- .

.PHONY: submodule_update clean build test
