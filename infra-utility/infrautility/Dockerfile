
ARG ALPINE_VERSION

FROM alpine:${ALPINE_VERSION}
ARG PYTHON_VERSION
###### Start Install python3 ######
RUN apk add \
    wget \
    gcc \
    make \
    zlib-dev \
    libffi-dev \
    openssl-dev \
    musl-dev
ENV PYTHONUNBUFFERED=1
RUN apk add --update --no-cache python3=~3.7 \
    python3 -m ensurepip \
    pip3 install --no-cache --upgrade pip setuptools

###### End Install python3 ######