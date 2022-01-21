update:
	go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle

build:
	bazel build //...

run:
	bazel run //app/cmd:cmd

test:
	bazel test //...