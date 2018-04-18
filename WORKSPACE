# Bazel Go rules.

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "f70c35a8c779bb92f7521ecb5a1c6604e9c3edd431e50b6376d7497abc8ad3c1",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.11.0/rules_go-0.11.0.tar.gz",
)

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_register_toolchains",
    "go_repository",
    "go_rules_dependencies",
)

go_rules_dependencies()

go_register_toolchains()

# Bazel Docker rules.

git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    tag = "v0.4.0",
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    go_image_repositories = "repositories",
)

go_image_repositories()

# Go dependencies.

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    tag = "v1.0.1",
)
