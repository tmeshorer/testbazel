
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


new_git_repository(
    name = "go_zap",
    build_file = "third_party/go/uber_zap.BUILD",
    commit = "a2773be06b9ac7c318a3a105b5c310af5730c6b4",	
    remote = "https://github.com/uber-go/zap.git",
)

