## Package
.PHONY: package.help

package.help:
	@echo '    package:'
	@echo ''
	@echo '        package                 show help.'
	@echo '        package.setup           setup dependence.'
	@echo '        package.build           build project.'
	@echo '        package.release         release project.'
	@echo ''

bin/goreleaser:
	@mkdir -p bin
	curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | BINARY=goreleaser bash -s
	@mv bin/goreleaser $@
.PHONY: bin/goreleaser

.PHONY: package
package:
	@if [ -z "${command}" ]; then \
		make package.help;\
	fi

.PHONY: package.setup
package.setup:
	@echo "----> setup package..."
	@echo ${MESSAGE_HAPPY}

.PHONY: package.build
package.build: bin/goreleaser
	@echo "----> build package..."
	bin/goreleaser build --snapshot --rm-dist
	@echo ${MESSAGE_HAPPY}

.PHONY: package.release
package.release:
	@echo "----> release package..."
	bin/goreleaser release --snapshot --rm-dist
	@echo ${MESSAGE_HAPPY}