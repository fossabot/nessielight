#!/bin/env bash

rm -rf dest
mkdir dest

if [[ $? -ne 0 ]]; then
  printf "mkdir error"
  exit 1
fi

printf "compling...\n"
# CGO_ENABLED=0 go build -trimpath -v -ldflags='-s -w -extldflags "-static"' -o dest/nessielight ./main
go build -trimpath -v -ldflags='-s -w -extldflags "-static"' -o dest/nessielight ./main

upx -5 dest/nessielight # 压缩，可以不要

sftp root@blog.sshwy.name:nessielight-dest <<< $'put dest/* .'

printf "restart service...\n"

ssh  root@blog.sshwy.name 'cd nessielight-dest; ./quickstart.sh'