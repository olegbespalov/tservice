FROM golang:1.14.3-alpine as builder

RUN apk update

ENV USER=appuser
ENV UID=10001

RUN adduser \    
   --disabled-password \    
   --gecos "" \    
   --home "/nonexistent" \    
   --shell "/sbin/nologin" \    
   --no-create-home \    
   --uid "${UID}" \    
   "${USER}"

WORKDIR /src

COPY go.mod .

ENV GO111MODULE=on
RUN go mod download
RUN go mod verify

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
   -ldflags='-w -s -extldflags "-static"' -a \
   -o /bin/app \
   cmd/tservice/main.go

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /bin/app /bin/app

USER appuser:appuser

VOLUME [ "/configs"]
VOLUME [ "/assets" ]

ENTRYPOINT ["/bin/app"]