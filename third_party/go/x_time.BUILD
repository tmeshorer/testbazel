package(default_visibility = ["@//visibility:public"])

load("@//third_party:go/build.bzl", "external_go_package")

external_go_package(
    base_pkg = "golang.org/x/time",
)

external_go_package(
    base_pkg = "golang.org/x/time",
    name = "rate",
    deps = [
        "@go_x_net//:context",
    ],
)

