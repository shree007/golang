
ARG ALPINE_VERSION
FROM alpine:${ALPINE_VERSION}


###### Start Install python3 ######
ARG PYTHON3_VERSION
ENV PYTHONUNBUFFERED=1
RUN apk update && apk add --update --no-cache python3=~${PYTHON3_VERSION}
###### End Install python3 #######

###### Start Install AWSCLI ######
ARG AWSCLI_VERSION
RUN apk add aws-cli=~${AWSCLI_VERSION}
###### End Install AWSCLI #######
