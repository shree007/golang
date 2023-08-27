
ARG ALPINE_VERSION
FROM alpine:${ALPINE_VERSION}

RUN apk add curl

###### Start Install python3 ######
ARG PYTHON3_VERSION
ENV PYTHONUNBUFFERED=1
RUN apk update && apk add --update --no-cache python3=~${PYTHON3_VERSION}
###### End Install python3 #######

###### Start Install AWSCLI ######
ARG AWSCLI_VERSION
RUN apk add aws-cli=~${AWSCLI_VERSION}
###### End Install AWSCLI #######

###### Start Install Kubectl ######
ARG KUBECTL_VERSION
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl \
    && chmod +x ./kubectl \
    && mv ./kubectl /usr/bin/kubectl
###### End Install Kubectl #######