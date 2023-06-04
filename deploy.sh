#!/bin/sh

./scripts/build.sh

DEST=/apps/page/bin
VERSION=$(./scripts/version.sh)
SERVER=root@mviana.dev

echo "Copy files to remote server ${SERVER}"
ssh "${SERVER}" "mkdir -p ${DEST} && id -u page &>/dev/null || useradd page"
scp -pr ./bin/www.linux.amd64-"${VERSION}" "${SERVER}":"$DEST"/www-n
scp -pr ./www.service "${SERVER}":/etc/systemd/system/www.service
ssh "${SERVER}" "mv ${DEST}/www-n ${DEST}/www && systemctl daemon-reload && systemctl start www.service && systemctl enable www.service && systemctl restart www.service"
