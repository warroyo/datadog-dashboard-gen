FROM ubuntu:latest

RUN locale-gen en_US.UTF-8
RUN dpkg-reconfigure locales
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8

RUN apt-get update; apt-get -y upgrade; apt-get clean
RUN apt-get install -y git curl; apt-get clean

RUN cd /tmp && \
    curl -O -L https://storage.googleapis.com/golang/go1.5.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go*.tar.gz && \
    rm go*.tar.gz
RUN cd /tmp && \
    curl -O -L https://github.com/pivotal-golang/golang-workspace-generator/releases/download/v0.0.1/workspace-generator-linux-amd64 && \
    mv workspace-generator-linux-amd64 /bin/workspace-generator && \
    chmod +x /bin/workspace-generator

ENV PATH /usr/local/go/bin:$PATH
