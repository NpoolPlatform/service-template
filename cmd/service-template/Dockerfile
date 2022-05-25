FROM jrei/systemd-centos:7

WORKDIR /app

COPY service-template /app
COPY ServiceTemplate.viper.yaml /app

ENTRYPOINT ["./service-template"]

EXPOSE 32700 32701 32702

CMD ["run"]
