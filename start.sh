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

exec supervisord -c /opt/gogs/supervisord.conf
