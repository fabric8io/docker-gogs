#!/bin/sh -

IFS=' 
	'

export USER_ID=$(id -u)
export GROUP_ID=$(id -g)
envsubst < /tmp/passwd.template > /tmp/passwd
export LD_PRELOAD=libnss_wrapper.so
export NSS_WRAPPER_PASSWD=/tmp/passwd
export NSS_WRAPPER_GROUP=/etc/group

PATH=/bin:/usr/bin:/usr/local/bin
HOME=${HOME:?"need \$HOME variable"}
USER=$(whoami)
export USER HOME PATH

mkdir ${HOME}/git

cd "$(dirname $0)"

(
  mkdir -p ${HOME}/ssh/ || true
  cd ${HOME}/ssh/
  /opt/gogs/ssh-hostkeygen
)

export GOGS_RUN_USER=${USER:-git}

exec ./start-gogs
