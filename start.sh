#!/bin/sh -

IFS=' 
	'
PATH=/bin:/usr/bin:/usr/local/bin
HOME=${HOME:?"need \$HOME variable"}
USER=$(whoami)
export USER HOME PATH

mkdir ${HOME}/git

cd "$(dirname $0)"

sed -i "s/RUN_USER = git/RUN_USER = ${USER}/" custom/conf/app.ini

exec ./gogs web -c custom/conf/app.ini
