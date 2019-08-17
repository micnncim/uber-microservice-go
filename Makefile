.PHONY: build
build:
	bazel build //...

.PHONY: test
test:
	bazel test //...

.PHONY: dep
dep:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=repositories.bzl%go_repositories

.PHONY: gazelle
gazelle:
	bazel run //:gazelle

.PHONY: buildifier
buildifier:
	bazel run //:buildifier

.PHONY: clean
clean:
	bazel clean
