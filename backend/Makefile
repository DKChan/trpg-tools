.PHONY: test test-verbose test-cover test-cover-html clean-test

# 运行所有测试
test:
	go test -v ./...

# 运行所有测试（详细输出）
test-verbose:
	go test -v -race ./...

# 运行测试并生成覆盖率报告
test-cover:
	go test -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out

# 运行测试并生成 HTML 覆盖率报告
test-cover-html:
	go test -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html

# 清理测试覆盖率文件
clean-test:
	rm -f coverage.out coverage.html

# 运行特定包的测试
test-package:
	@if [ -z "$(pkg)" ]; then \
		echo "Usage: make test-package pkg=./path/to/package"; \
		exit 1; \
	fi
	go test -v $(pkg)

# 运行测试并检查覆盖率阈值（需要传入 threshold 参数）
test-cover-check:
	@if [ -z "$(threshold)" ]; then \
		echo "Usage: make test-cover-check threshold=70"; \
		exit 1; \
	fi
	@go test -coverprofile=coverage.out -covermode=atomic ./... && \
	total_coverage=$$(go tool cover -func=coverage.out | grep total | awk '{print $$3}' | sed 's/%//') && \
	echo "Total coverage: $$total_coverage%" && \
	if [ $$(echo "$$total_coverage < $(threshold)" | bc -l) -eq 1 ]; then \
		echo "Coverage $${total_coverage}% is below threshold $(threshold)%"; \
		exit 1; \
	else \
		echo "Coverage $${total_coverage}% meets threshold $(threshold)%"; \
	fi
