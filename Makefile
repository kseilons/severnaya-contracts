BUF_VERSION=1.6.0

.PHONY: proto
proto: clean format gen lint git

.PHONY: gen
gen:
	@$(GOPATH)/bin/buf generate
		@for dir in $(CURDIR)/gen/go/*/; do \
		  cd $$dir && \
		  folder_name=$$(basename $$dir) && \
		  go mod init github.com/kseilons/severnaya-contracts/gen/go/$$folder_name && go mod tidy; \
		done

.PHONY: lint
lint:
	@$(GOPATH)/bin/buf lint

.PHONY: format
format:
	@$(GOPATH)/bin/buf format


.PHONY: clean
clean:
	@rm -rf ./gen || true

.PHONY: git
git: git_commit git_push

.PHONY: git_commit
git_commit:
	@git commit -am "Update proto files"  || true

.PHONY: git_push
git_push:
	@git push  || true