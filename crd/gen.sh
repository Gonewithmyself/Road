#!/bin/bash

bash ./hack/update-codegen.sh

mv crd/pkg/apis/hellocrd/v1/* pkg/apis/hellocrd/v1

rm -rf pkg/client

mv crd/pkg/client   pkg/

rm -rf crd