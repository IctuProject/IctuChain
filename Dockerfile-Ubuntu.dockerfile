# FROM
FROM --platform=linux ubuntu:22.04
ARG BUILDARCH

# CONSTANTES CON LA VERSION DE LAS DEPENDENCIAS A DESCARGAR
ENV GO_VERSION=1.20.2
ENV IGNITE_VERSION=0.26.1
ENV NODE_VERSION=18.x

ENV LOCAL=/usr/local
ENV GOROOT=$LOCAL/go
ENV HOME=/root
ENV GOPATH=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/bin

# INSTALL LINUX TOOLS TO DEVELOP, BUILD AND RUN 

ENV PACKAGES curl gcc jq
RUN apt-get update
RUN apt-get install -y $PACKAGES

# DESCARGAR LAS DEPENDENCIAS

# GO - descargar la versio especifica de go que necesitem i la descomprimir al directori que volem
RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-$BUILDARCH.tar.gz | tar -C $LOCAL -xzf -

#IGNITE 
RUN curl -L https://get.ignite.com/cli@v${IGNITE_VERSION}! | bash

#NODE 
RUN curl -fsSL https://deb.nodesource.com/setup_${NODE_VERSION} | bash -
RUN apt-get install -y nodejs

EXPOSE 1317 3000 4500 5000 26657

WORKDIR /chain

#COPY go.mod /chain/go.mod
#RUN go mod download
#RUN rm /chain/go.mod