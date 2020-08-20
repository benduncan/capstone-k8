FROM golang:1.11-alpine AS build-env

# Build phase
RUN apk add build-base git

ENV API_PORT 8000
ENV GOPATH /workspace/capstone-k8/
ENV GOBIN /workspace/capstone-k8/bin

ADD ./ $GOPATH/
WORKDIR $GOPATH/

RUN make clean
RUN make build

RUN apk del build-base git

# Next, just copy the golang binary, create a lightweight environment

FROM alpine
WORKDIR /workspace/capstone-k8/
RUN apk add ca-certificates

COPY --from=build-env /workspace/capstone-k8/bin/ /workspace/capstone-k8/bin/

EXPOSE $API_PORT
ENTRYPOINT ["/workspace/capstone-k8/bin/capstone-k8"]