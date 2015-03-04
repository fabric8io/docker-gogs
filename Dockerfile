FROM debian:stable

ENV GOGS_VERSION 0.5.13

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y curl apt-transport-https openssh-server sqlite3 unzip sudo git && \
    apt-get clean && \
    rm -f /etc/ssh/ssh_host_*_key_* && \
    mkdir /var/run/sshd && \
    curl -o /tmp/gogs.zip http://gogs.dn.qbox.me/gogs_v0.5.13_linux_amd64.zip && \
    cd /opt && unzip /tmp/gogs.zip && \
    rm -rf /tmp/gogs.zip /opt/__MACOSX $(find /opt/gogs -name .DS_Store) $(find /opt/gogs -name .idea)

RUN chmod +x /opt/gogs/gogs && \
    useradd -mr git && \
    mkdir -p /opt/gogs/data /opt/gogs/custom/conf /opt/gogs/log && \
    cp /opt/gogs/conf/app.ini /opt/gogs/custom/conf/app.ini && \
    chown -R git:git /opt/gogs/data /opt/gogs/conf/ /opt/gogs/custom/conf/ /opt/gogs/log

ADD start.sh /start.sh

ENV DB_TYPE sqlite3

# Expose our port
EXPOSE 22 3000

WORKDIR /opt/gogs

CMD ["/start.sh"]
