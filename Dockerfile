FROM golang:onbuild

ENV MESSAGR_APP_CONFIG=/local/config/messagr/production.yaml

EXPOSE 8080
