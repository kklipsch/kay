APP_NAME = kay

LOCALPATH= $(shell pwd)

#make outputs
BIN_PATH= $(LOCALPATH)/bin
EXEC_NAME= $(BIN_PATH)/$(APP_NAME)
PACK_NAME= $(BIN_PATH)/$(APP_NAME).zip

#local install locations
INSTALL_DIR= $(HOME)/bin
BASH_SOURCES= $(HOME)/.sources

$(EXEC_NAME): bin
	go build -o $(EXEC_NAME) *.go

$(PACK_NAME): $(EXEC_NAME)
	zip -rj $(PACK_NAME) $(EXEC_NAME) ./autocomplete/kay.bash

package: $(PACK_NAME)

test: 
	go test ./...

integration: clean $(PACK_NAME)
	ansible-playbook -i integration/local integration/test.yml --extra-vars "kay_package=$(PACK_NAME)"

install: $(PACK_NAME)
	ansible-playbook -i integration/local integration/install.yml --extra-vars '{"kay_package": "$(PACK_NAME)", "install_directory": {"stdout": "$(INSTALL_DIR)"}, "bash_sources": {"stdout": "$(BASH_SOURCES)"}}'

help: $(EXEC_NAME)
	$(EXEC_NAME) -h

help.%: clean $(EXEC_NAME)
	$(EXEC_NAME) help $(@:help.%=%)	

bin:
	mkdir -p bin

dep:
	go get github.com/kklipsch/cli

clean:
	go clean
	rm -rf bin
