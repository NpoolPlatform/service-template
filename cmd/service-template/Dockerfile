FROM jrei/systemd-centos:7

COPY ServiceTemplate.viper.yaml /etc/ServiceTemplate/ServiceTemplate.viper.yaml
COPY service-template /usr/local/bin/service-template

CMD ["service-template", "run"]
