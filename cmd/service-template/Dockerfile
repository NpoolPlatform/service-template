FROM jrei/systemd-centos:7

WORKDIR /app

COPY service-template /app

ENTRYPOINT ["service-template"]

EXPOSE 32700 32701 32702

CMD ["run"]
