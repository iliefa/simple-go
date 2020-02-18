FROM golang
COPY . .
RUN go get github.com/gorilla/mux
RUN CGO_ENABLED=0 go build -o /myapp
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

FROM scratch
COPY --from=0 /myapp /myapp
WORKDIR /workspace
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /etc_passwd /etc/passwd
# Assuming the source code is collocated to this Dockerfile, copy the whole
# directory into the container that is building the Docker image.
EXPOSE 8080
USER nobody