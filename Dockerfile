FROM alpine:3.12 as base

# stage 1 - build go binary

FROM base as builder
# install build tools for go
RUN apk add --no-cache go
# set go env
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
# copy source code from repo folder proxy
COPY proxy /go/src/proxy
# build go binary
RUN cd /go/src/proxy && go build -o /go/bin/proxy

# stage 2 - build docker image
FROM base
# copy go binary from builder stage
COPY --from=builder /go/bin/proxy /usr/local/bin/proxy
# run go binary
CMD ["proxy"]