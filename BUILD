load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/whichwit/resume
gazelle(name = "gazelle")

go_library(
    name = "resume_lib",
    srcs = ["main.go"],
    data = [
        "resume.json"
    ],
    importpath = "github.com/whichwit/resume",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "resume",
    data = glob(["templates/*.tmpl"]),
    embed = [":resume_lib"],
    visibility = ["//visibility:public"],
)