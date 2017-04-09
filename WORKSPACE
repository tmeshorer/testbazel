
# GOLANG INIT
load("//tools/go:go_configure.bzl", "go_configure")

go_configure()

bind(
    name = "go_package_prefix",
    actual = "//:go_package_prefix",
)


new_git_repository(
    name = "go_logrus",
    build_file = "third_party/go/Sirupsen_logrus.BUILD",
    commit = "4b6ea7319e214d98c938f12692336f7ca9348d6b",
    remote = "https://github.com/Sirupsen/logrus.git",
)