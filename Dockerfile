FROM golang:1.15.5

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      less \
      vim-gtk3 \
      && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /work
