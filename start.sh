#!/bin/sh -

IFS=' 
	'
PATH=/bin:/usr/bin:/usr/local/bin
HOME=${HOME:?"need \$HOME variable"}
USER=$(whoami)
export USER HOME PATH

mkdir ${HOME}/git

cd "$(dirname $0)"

(
  mkdir -p data/ssh/ || true
  cd data/ssh/
  ../../ssh-hostkeygen
)

export GOGS_RUN_USER=${USER:-git}

exec ./start-gogs
