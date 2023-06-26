#!/bin/bash

EXPIRE_DATETIME=$(openssl x509 -in "inschdev/web/cert/device.pem" -noout -dates | grep notAfter | awk -F '=' '{print $2}')
EXPIRE_TS=$(date -d "$EXPIRE_DATETIME" +%s)
NEXT_MON_TS=$(date -d '+1 month' +%s)

if [[ $EXPIRE_TS -lt $NEXT_MON_TS ]]; then
	# 一个月之内过期则重新生成证书
	log "重新生成证书中..."
	cat >nsupdate.key <<EOF
key "insport" {
	algorithm hmac-sha256;
	secret "yKDjdg/5OotK2yq9OGU/clI+5TeKOdfDx/wtXyr3heE=";
};
EOF
	# 需要设置代理
	docker run --rm \
		-v "$(pwd)/out":/acme.sh \
		-v "$(pwd)/nsupdate.key":/tmp/nsupdate.key \
		-e NSUPDATE_SERVER="zwszsports.com" \
		-e NSUPDATE_KEY="/tmp/nsupdate.key" \
		-e http_proxy=$http_proxy \
		-e https_proxy=$https_proxy \
		neilpang/acme.sh --issue --dns dns_nsupdate -d *.device.zwszsports.com --server letsencrypt --keylength 2048

	cat "out/*.device.zwszsports.com/*.device.zwszsports.com.key" >"inschdev/web/cert/device.key"
	check_error "证书生成失败...."
	cat "out/*.device.zwszsports.com/fullchain.cer" >"inschdev/web/cert/device.pem"
	check_error "证书生成失败...."
	log "证书重新生成完成..."
fi
