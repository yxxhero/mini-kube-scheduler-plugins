#!/bin/bash -
#===============================================================================
#
#          FILE: build.sh
#
#         USAGE: ./build.sh
#
#   DESCRIPTION: 
#
#       OPTIONS: ---
#  REQUIREMENTS: ---
#          BUGS: ---
#         NOTES: ---
#        AUTHOR: xiongxiong.yuan , 
#  ORGANIZATION: 
#       CREATED: 10/09/2021 07:40:24 PM
#      REVISION:  ---
#===============================================================================

set -o nounset                                  # Treat unset variables as an error

build(){
    go build -ldflags='-X k8s.io/component-base/version.gitVersion=v1.18.9 -X k8s.io/client-go/pkg/version.gitVersion=v1.18.9' -o kube-scheduler cmd/cks/main.go
}


build $@
