load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@bazel_skylib//rules:run_binary.bzl", "run_binary")

# gazelle:prefix github.com/whichwit/resume
gazelle(name = "gazelle")

go_library(
    name = "resume_lib",
    srcs = ["main.go"],
    data = ["resume.json"],
    importpath = "github.com/whichwit/resume",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "resume",
    embed = [":resume_lib"],
    visibility = ["//visibility:public"],
)
