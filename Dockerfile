FROM golang:1.14 as builder

RUN  apt-get update && apt-get -y install bash git make zip bzr && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*
ADD . /go/src/github.com/yamamoto-febc/terraform-provider-gmailfilter
WORKDIR /go/src/github.com/yamamoto-febc/terraform-provider-gmailfilter
ENV GOPROXY=https://proxy.golang.org
RUN ["make", "tools", "build"]

###

FROM hashicorp/terraform:0.12.21

COPY --from=builder /go/src/github.com/yamamoto-febc/terraform-provider-gmailfilter/bin/* /bin/

VOLUME ["/workdir"]
WORKDIR /workdir

ENTRYPOINT ["/bin/terraform"]
CMD ["--help"]

