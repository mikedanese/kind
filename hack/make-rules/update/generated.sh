#!/usr/bin/env bash
# Copyright 2018 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# 'go generate's kind, using tools from vendor (go-bindata)
set -o errexit -o nounset -o pipefail

# cd to the repo root and setup go
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../.." && pwd -P)"
cd "${REPO_ROOT}"
source hack/build/setup-go.sh

# build the generators using the tools module
cd "${REPO_ROOT}/hack/tools"
go build -o "${REPO_ROOT}/bin/deepcopy-gen" k8s.io/code-generator/cmd/deepcopy-gen
go build -o "${REPO_ROOT}/bin/protoc-gen-go" google.golang.org/protobuf/cmd/protoc-gen-go
# go back to the root
cd "${REPO_ROOT}"

export PATH="${REPO_ROOT}/bin/:${PATH}"

# turn off module mode before running the generators
# https://github.com/kubernetes/code-generator/issues/69
# we also need to populate vendor

# run the generators
# TODO: -o "${REPO_ROOT}/../.." is a weird work-around ...
deepcopy-gen -i ./pkg/internal/apis/config/ -o "${REPO_ROOT}/../.." -O zz_generated.deepcopy --go-header-file hack/tools/boilerplate.go.txt
deepcopy-gen -i ./pkg/apis/config/v1alpha4 -o "${REPO_ROOT}/../.." -O zz_generated.deepcopy --go-header-file hack/tools/boilerplate.go.txt

protoc --go_out=paths=source_relative:. pkg/build/nodeimage/internal/bazelapi/analysis.proto


# set module mode back, return to repo root and gofmt to ensure we format generated code
make gofmt
