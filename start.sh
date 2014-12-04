#!/bin/bash

if [ ! -e /etc/ssh/ssh_host_rsa_key ]; then
  dpkg-reconfigure openssh-server
fi

/usr/sbin/sshd -D &

sed -i "s/^DB_TYPE =.*/DB_TYPE = ${DB_TYPE}/" /etc/gogs/conf/app.ini
sed -i "s/^HOST =.*/HOST = $(eval echo $DB_HOST)/" /etc/gogs/conf/app.ini
sed -i "s/^NAME =.*/NAME = ${DB_NAME}/" /etc/gogs/conf/app.ini
sed -i "s/^USER =.*/USER = ${DB_USER}/" /etc/gogs/conf/app.ini
sed -i "s/^PASSWD =.*/PASSWD = ${DB_PASSWD}/" /etc/gogs/conf/app.ini
sed -i "s/^INSTALL_LOCK =.*/INSTALL_LOCK = true/" /etc/gogs/conf/app.ini

exec gogs run web
