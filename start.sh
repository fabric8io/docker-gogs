#!/bin/bash

if [ ! -e /etc/ssh/ssh_host_rsa_key ]; then
  dpkg-reconfigure openssh-server
fi

/usr/sbin/sshd -D &

INI_FILE=/opt/gogs/custom/conf/app.ini
SKIP_TLS_VERIFY=${SKIP_TLS_VERIFY:-false}

sed -i "s/^RUN_MODE =.*/RUN_MODE = ${RUN_MODE:-prod}/" ${INI_FILE}

sed -i "s/^OFFLINE_MODE =.*/OFFLINE_MODE= ${OFFLINE_MODE:-true}/" ${INI_FILE}

sed -i "s/^DB_TYPE =.*/DB_TYPE = ${DB_TYPE:-sqlite3}/" ${INI_FILE}
sed -i "s/^HOST =.*/HOST = $(eval echo $DB_HOST)/" ${INI_FILE}
sed -i "s/^NAME =.*/NAME = ${DB_NAME}/" ${INI_FILE}
sed -i "s/^USER =.*/USER = ${DB_USER}/" ${INI_FILE}
sed -i "s/^PASSWD =.*/PASSWD = ${DB_PASSWD}/" ${INI_FILE}
#sed -i "s/^INSTALL_LOCK =.*/INSTALL_LOCK = true/" ${INI_FILE}

sed -i "s/SKIP_TLS_VERIFY =.*/SKIP_TLS_VERIFY = ${SKIP_TLS_VERIFY}/" ${INI_FILE}

sed -i "s/^TASK_INTERVAL =.*/TASK_INTERVAL = ${TASK_INTERVAL:-0}/" ${INI_FILE}
sed -i "s/^DOMAIN =.*/DOMAIN = ${DOMAIN:-gogs.fabric8.local}/" ${INI_FILE}
sed -i "s/^HTTP_PORT =.*/HTTP_PORT = ${HTTP_PORT:-3000}/" ${INI_FILE}

sed -i "s/^PROTOCOL =.*/PROTOCOL = ${PROTOCOL:-http}/" ${INI_FILE}
sed -i "s|^ROOT_URL =.*|ROOT_URL = ${ROOT_URL:-${PROTOCOL:-http}://${DOMAIN:-gogs.fabric8.local}/}|" ${INI_FILE}

sed -i "s|^INSTALL_LOCK =.*|INSTALL_LOCK = true|" ${INI_FILE}

if [ "${PROTOCOL}" == "https" ]; then
  mkdir -p /opt/gogs/custom/https
  cd /opt/gogs/custom/https
  CERT_HOST=${CERT_HOST:-${DOMAIN-gogs.fabric8.local}}
  echo "Creating cert for host ${CERT_HOST}"
  /opt/gogs/gogs cert -host=${CERT_HOST},${DOMAIN:-gogs.fabric8.local},$(hostname -i),localhost
  chown git /opt/gogs/custom/https/*
fi

exec sudo -u git -H sh -c "cd /opt/gogs; exec ./gogs web"
