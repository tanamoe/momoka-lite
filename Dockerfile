ARG  BUN_IMAGE=oven/bun:1.3.0
ARG  BUILDER_IMAGE=golang:1.25.3-alpine3.21
ARG  DISTROLESS_IMAGE=gcr.io/distroless/static

############################
# STEP 1 build the docs page
############################
FROM ${BUN_IMAGE} as docs-builder

WORKDIR /build/docs/

COPY ./docs/ .

RUN bun install --frozen-lockfile

RUN NODE_ENV=production bun run docs:build

############################
# STEP 2 build executable binary
############################
FROM ${BUILDER_IMAGE} as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

WORKDIR $GOPATH/src/mypackage/myapp/

# use modules
COPY ./go.mod .

RUN apk add git

ENV GO111MODULE=on
RUN go mod download
RUN go mod verify

COPY . .

# Copy our pre-built document
COPY --from=docs-builder /build/docs/dist ./docs/dist

# Build the static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	-ldflags='-w -s -extldflags "-static"' -a \
	-o /go/bin/momoka-lite ./cmd/serve.go

############################
# STEP 3 build a small image
############################
# using static nonroot image
# user:group is nobody:nobody, uid:gid = 65534:65534
FROM ${DISTROLESS_IMAGE}

# Copy our static executable
COPY --from=builder /go/bin/momoka-lite /go/bin/momoka-lite

# Run the hello binary.
ENTRYPOINT ["/go/bin/momoka-lite", "serve"]
