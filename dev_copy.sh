#!/usr/bin/env bash

# Script to copy build binaries to terraform plugins location

rm -rf ~/.terraform.d/plugins/delphix.com/local/dct/0.0-dev/$CP_TARGET/
mkdir -p ~/.terraform.d/plugins/delphix.com/local/dct/0.0-dev/$CP_TARGET/
cp $CP_PATH ~/.terraform.d/plugins/delphix.com/local/dct/0.0-dev/$CP_TARGET/