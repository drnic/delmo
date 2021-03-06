#!/bin/bash

set -x
set -e

echo "Building binary"
export GOPATH=$GOPATH:$PWD/delmo
cd delmo/src/github.com/bodymindarts/delmo
make dev
BINARY="bin/delmo"

echo
echo "Downloading machine info"
aws --region ${AWS_REGION} s3 cp s3://${AWS_BUCKET}/${machine_name}.zip ./

echo
echo "Importing ${machine_name}"
machine-import ${machine_name}.zip
# The permission isn't set properly on import
chmod 0600 /root/.docker/machine/machines/${machine_name}/id_rsa

echo
echo "Testing example/webapp"
${BINARY} -f example/webapp/delmo.yml -m ${machine_name}

echo
echo "Executing example/webapp in parralel"
${BINARY} -f example/webapp/delmo.yml -m ${machine_name} --parallel
