# 定义工作环境名称
# gazelle:repo bazel_gazelle
workspace(name = "com_github_lalifeier_vvgo_mall")

# 下载所需要的软件包
load("//:DOWNLOAD.bzl", "download_package")

download_package()

# Bazel Go语言 规则集 初始化
# 导入go_register_toolchains和go_rules_dependencies方法
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# 初始化go规则集的依赖项
go_rules_dependencies()

# 注册go 1.19.5版本的工具链，包含下载安装go环境
go_register_toolchains(version = "1.19.5")

# Bazel Gazelle 规则集 初始化
# 导入gazelle_dependencies和go_repository方法
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:repos.bzl", "go_dependencies")

# gazelle:repository_macro repos.bzl%go_dependencies
go_dependencies()

# 初始化Gazelle规则集的依赖项
gazelle_dependencies()

# Bazel Docker 规则集 初始化
# 导入container_repositories方法
load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")
container_repositories()

# 导入container_deps方法
load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

# 导入container_pull方法
load("@io_bazel_rules_docker//container:container.bzl", "container_pull")

# 拉取Alpine Linux
# 该发行版使用musl libc，并且缺乏一些调试工具。
container_pull(
    name = "alpine_linux_amd64",
    registry = "index.docker.io",
    repository = "library/alpine",
    tag = "latest",
)

# 拉取Debian-Slim Linux
container_pull(
     name = "slim_linux_amd64",
     registry = "index.docker.io",
     repository = "library/debian",
     tag = "stable-slim",
)

# Bazel Kubernetes 规则集 初始化

# load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories")

# k8s_repositories()

# Bazel pkg 规则集 初始化

# load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

# rules_pkg_dependencies()

# Bazel protoc 初始化

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# Bazel Buf 规则集 初始化

# load("@rules_buf//buf:repositories.bzl", "rules_buf_dependencies", "rules_buf_toolchains")
# load("@rules_buf//gazelle/buf:repositories.bzl", "gazelle_buf_dependencies")
# load("//:buf_deps.bzl", "buf_deps")

# # gazelle:repository_macro buf_deps.bzl%buf_deps
# buf_deps()

# rules_buf_dependencies()

# rules_buf_toolchains(version = "v1.12.0")

# gazelle_buf_dependencies()

# Bazel GPRC 规则集 初始化

# load("@rules_proto_grpc//:repositories.bzl", "rules_proto_grpc_repos", "rules_proto_grpc_toolchains")

# rules_proto_grpc_toolchains()

# rules_proto_grpc_repos()

# Bazel Protobuf 规则集 初始化

# load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

# rules_proto_dependencies()

# rules_proto_toolchains()
