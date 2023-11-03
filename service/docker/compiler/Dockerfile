FROM ghcr.io/goreleaser/goreleaser-cross:v1.21.3

RUN set -ex \
    && apt-get update \
    && apt-get install -y gcc g++ zip libgl1-mesa-dev xorg-dev zlib1g-dev libmpc-dev libmpfr-dev libgmp-dev icnsutils imagemagick \
    && rm -rf /var/lib/apt/lists/*

RUN set -ex \
    && mkdir -p /go/src/github.com/adrianrudnik/ablegram \
    && cd /go/src/github.com/adrianrudnik/ablegram \
    && go mod init github.com/adrianrudnik/ablegram \
    && go get fyne.io/fyne/v2@latest \
    && go install fyne.io/fyne/v2/cmd/fyne@latest \
    && mkdir -p /go/bin \
    && mv /root/go/bin/fyne /usr/local/bin

COPY release.sh /usr/local/bin/release

WORKDIR go/src/github.com/adrianrudnik/ablegram/service

ENTRYPOINT []

CMD ["/bin/bash"]
