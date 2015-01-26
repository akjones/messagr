FROM golang:onbuild

#VOLUME /local/config/messagr
ENV MESSAGR_APP_CONFIG=/local/config/messagr/production.yaml

EXPOSE 8080
