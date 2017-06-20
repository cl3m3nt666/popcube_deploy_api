FROM golang:1.7

LABEL maintainer "clement@le-corre.eu" \
      description "API rest for deploy container easily for popcube organisation"

ARG DOCKER_VERSION=17.03*

ENV GOPATH=$GOPATH:/go/api \
    GOBIN=$GOPATH/bin \
    XTOKEN=1234 \
    DEFAULT_DOMAIN=popcube.xyz \
    BASE_NAME_HOST_DB=database \
    DEFAULT_DATABASE=popcube \
    DEFAULT_ORG_PATH=/organisation \
    ORGANISATION_TEMPLATE=/organisation_template


RUN apt-get update && \
    apt-get -y install \
      apt-transport-https \
      ca-certificates \
      curl \
      software-properties-common && \
    curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - && \
    add-apt-repository \
         "deb [arch=amd64] https://download.docker.com/linux/debian \
         $(lsb_release -cs) \
         stable" && \
    apt-get update && \
    apt-get -y install docker-ce=$DOCKER_VERSION && \
    rm -rf /var/lib/apt/lists/*

COPY base_org /organisation_template
COPY api/ /go/api
COPY run.sh /bin

WORKDIR /go/api
VOLUME ["/organisation"]

EXPOSE 80
CMD ["/bin/run.sh"]
