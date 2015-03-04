#!/bin/bash

if [ ! -e /etc/ssh/ssh_host_rsa_key ]; then
  dpkg-reconfigure openssh-server
fi

/usr/sbin/sshd -D &

INI_FILE=/opt/gogs/custom/conf/app.ini

sed -i "s/^RUN_MODE =.*/RUN_MODE = ${RUN_MODE:-prod}/" ${INI_FILE}

sed -i "s/^OFFLINE_MODE =.*/OFFLINE_MODE= ${OFFLINE_MODE:-true}/" ${INI_FILE}

sed -i "s/^DB_TYPE =.*/DB_TYPE = ${DB_TYPE:-sqlite3}/" ${INI_FILE}
sed -i "s/^HOST =.*/HOST = $(eval echo $DB_HOST)/" ${INI_FILE}
sed -i "s/^NAME =.*/NAME = ${DB_NAME}/" ${INI_FILE}
sed -i "s/^USER =.*/USER = ${DB_USER}/" ${INI_FILE}
sed -i "s/^PASSWD =.*/PASSWD = ${DB_PASSWD}/" ${INI_FILE}
#sed -i "s/^INSTALL_LOCK =.*/INSTALL_LOCK = true/" ${INI_FILE}

exec sudo -u git -H sh -c "cd /opt/gogs; exec ./gogs web"
