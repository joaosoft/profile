############################
# STEP 1 build binary
############################
FROM golang:alpine AS builder

LABEL maintainer="João Ribeiro <joaosoft@gmail.com>"

RUN apk update && apk add --no-cache \
	dep \
	git

WORKDIR /go/src/profile
COPY . .

RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o profile .



############################
# STEP 2 run binary
############################
FROM scratch
COPY --from=builder /go/src/profile/profile .

ENTRYPOINT ./profile