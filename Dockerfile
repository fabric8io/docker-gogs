FROM debian:stable

ENV GOGS_VERSION 0.5.8
ENV DB_TYPE mysql
ENV DB_HOST 127.0.0.1:3306
ENV DB_NAME gogs
ENV DB_USER gogs
ENV DB_PASSWD gogs

RUN apt-get update && \
    apt-get install -y curl apt-transport-https openssh-server && \
    curl https://deb.packager.io/key | apt-key add - && \
    echo "deb https://deb.packager.io/gh/pkgr/gogs wheezy pkgr" > /etc/apt/sources.list.d/gogs.list && \
    apt-get update && \
    apt-get install -y gogs=${GOGS_VERSION}* && \
    apt-get clean && \
    rm -f /etc/ssh/ssh_host_*_key_* && \
    mkdir /var/run/sshd

ADD start.sh /start.sh

# Expose our port
EXPOSE 22 6000

CMD ["/start.sh"]
