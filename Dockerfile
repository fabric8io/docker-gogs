FROM gogs/gogs:0.9.71

ENTRYPOINT ["/app/gogs/start.sh"]

COPY app.ini /app/gogs/custom/conf/app.ini
COPY start.sh /app/gogs/start.sh

COPY sshkeygen /app/gogs/sshkeygen
COPY ssh-keygen /usr/bin/ssh-keygen

RUN mkdir -p /app/gogs/data && chmod 777 /app/gogs/data /app/gogs/custom/conf

USER git

ENV HOME=/app/gogs/data PATH=/app/gogs:$PATH
