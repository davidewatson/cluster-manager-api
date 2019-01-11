FROM debian:stretch-20181226

ARG GOLANG_VERSION=1.11.4
ARG HELM_VERSION=2.12.1

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git \
                       apt-transport-https \
                       ca-certificates \
                       curl \
                       gnupg2 \
                       software-properties-common

# Install docker
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -

RUN add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/debian \
   $(lsb_release -cs) \
   stable"

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y docker-ce

# Install golang
ADD https://dl.google.com/go/go${GOLANG_VERSION}.linux-amd64.tar.gz /tmp
RUN tar -C /usr/local -xzf /tmp/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm -f /tmp/go${GOLANG_VERSION}.linux-64.tar.gz

ENV PATH /usr/local/go/bin:/go/bin:$PATH
ENV GOPATH /go

# Install kubectl
RUN curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
RUN echo "deb https://apt.kubernetes.io/ kubernetes-stretch main" | tee -a /etc/apt/sources.list.d/kubernetes.list
RUN apt-get update && \
    apt-get install -y kubectl

# Install helm
ADD https://storage.googleapis.com/kubernetes-helm/helm-v${HELM_VERSION}-linux-amd64.tar.gz /tmp
RUN cd /tmp && \
    tar -xzf /tmp/helm-v${HELM_VERSION}-linux-amd64.tar.gz && \
    mv /tmp/linux-amd64/helm /usr/local/bin && \
    rm -rf /tmp/linux-amd64 /tmp/helm-v${HELM_VERSION}-linux-amd64.tar.gz

# Install kind
RUN mkdir -p /go/src/sigs.k8s.io && \
    cd /go/src/sigs.k8s.io/ && \
    git clone https://github.com/kubernetes-sigs/kind.git && \
    go install sigs.k8s.io/kind && \
    rm -rf /go/src/sigs.k8s.io/

# Cleanup unnecessary files
RUN rm -rf /var/lib/apt/lists/*
