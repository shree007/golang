
ARG ALPINE_VERSION
FROM alpine:${ALPINE_VERSION}

RUN apk add curl wget

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


###### Start Install terraform ######
ARG TERRAFORM_VERSION
RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip  \
    &&  unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    &&  mv terraform /usr/bin/terraform \
    &&  chmod +x /usr/bin/terraform
###### Start Install terraform ######

###### Start Install terragrunt ######
ARG TERRAGRUNT_VERSION
RUN curl -L https://github.com/gruntwork-io/terragrunt/releases/download/${TERRAGRUNT_VERSION}/terragrunt_linux_amd64 -o /usr/local/bin/terragrunt \
    && chmod +x /usr/local/bin/terragrunt
###### Start Install terragrunt ######