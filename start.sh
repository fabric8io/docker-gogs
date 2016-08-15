#!/bin/bash

ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key -N '' < /dev/null
ssh-keygen -t rsa -b 4096 -f /etc/ssh/ssh_host_rsa_key -N '' < /dev/null

sed -i \
  -e 's|^HostKey /etc/ssh/ssh_host_ecdsa_key|#HostKey /etc/ssh/ssh_host_ecdsa_key|' \
  -e 's|^PasswordAuthentication yes|PasswordAuthentication no|' \
  -e 's|^#PermitRootLogin yes|PermitRootLogin no|' \
  -e 's|^GSSAPIAuthentication yes|GSSAPIAuthentication no|' \
  /etc/ssh/sshd_config

echo AllowUsers git >> /etc/ssh/sshd_config

chown -R ${GOGS_RUN_USER}:${GOGS_RUN_USER} /opt/gogs/data /opt/gogs/custom/conf/ /opt/gogs/log

# lets default some environments using expressions if they are not passed in

if [ -z "$GOGS_SERVER__DOMAIN" ]; then
    if [ -z "$DOMAIN" ]; then
        echo "Environment variable DOMAIN is not set!!!"
        exit 1
    fi 
    export GOGS_SERVER__DOMAIN="gogs.${DOMAIN}"
fi 

if [ -z "$GOGS_SERVER__PROTOCOL" ]; then
    export GOGS_SERVER__PROTOCOL="http"
fi 


if [ -z "$GOGS_SERVER__ROOT_URL" ]; then
    export GOGS_SERVER__ROOT_URL="${GOGS_SERVER__PROTOCOL}://${GOGS_SERVER__DOMAIN}"
fi 

if [ -z "$ADMIN_USER_EMAIL" ]; then
    export ADMIN_USER_EMAIL="${JENKINS_GOGS_EMAIL}"
fi 
if [ -z "$ADMIN_USER_PASSWORD" ]; then
    export ADMIN_USER_PASSWORD="${JENKINS_GOGS_PASSWORD}"
fi 
if [ -z "$ADMIN_USER_NAME" ]; then
    export ADMIN_USER_NAME="${JENKINS_GOGS_USER}"
fi 

# openshift OAuth
if [ -z "$GOGS_OAUTH_OPENSHIFT__SCOPES" ]; then
    if [ -z "$DOMAIN" ]; then
        echo "Environment variable DOMAIN is not set!!!"
        exit 1
    fi 
    export GOGS_OAUTH_OPENSHIFT__SCOPES="https://${DOMAIN}:8443/console/user"
fi 
if [ -z "$GOGS_OAUTH_OPENSHIFT__AUTH_URL" ]; then
    if [ -z "$DOMAIN" ]; then
        echo "Environment variable DOMAIN is not set!!!"
        exit 1
    fi 
    export GOGS_OAUTH_OPENSHIFT__AUTH_URL="https://${DOMAIN}:8443/oauth/authorize"
fi 
if [ -z "$GOGS_OAUTH_OPENSHIFT__TOKEN_URL" ]; then
    if [ -z "$DOMAIN" ]; then
        echo "Environment variable DOMAIN is not set!!!"
        exit 1
    fi 
    export GOGS_OAUTH_OPENSHIFT__TOKEN_URL="https://${DOMAIN}:8443/oauth/token"
fi 


echo "GOGS_SERVER__PROTOCOL = ${GOGS_SERVER__PROTOCOL}"
echo "GOGS_SERVER__DOMAIN = ${GOGS_SERVER__DOMAIN}"
echo "GOGS_SERVER__ROOT_URL = ${GOGS_SERVER__ROOT_URL}"

echo "ADMIN_USER_EMAIL = ${ADMIN_USER_EMAIL}"
echo "ADMIN_USER_NAME = ${ADMIN_USER_NAME}"


exec supervisord -c /opt/gogs/supervisord.conf
