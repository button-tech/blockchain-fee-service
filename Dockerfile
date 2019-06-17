FROM golang:latest AS builder

# write your ssh private key into the id_rsa
ADD ./ssh_key/id_rsa /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o  bin/main .

FROM scratch

COPY --from=builder /build/bin /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
ENTRYPOINT ["/app/main"]
