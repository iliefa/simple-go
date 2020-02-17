FROM golang

WORKDIR /workspace

# Assuming the source code is collocated to this Dockerfile, copy the whole
# directory into the container that is building the Docker image.
COPY . .
RUN go get github.com/gorilla/mux
RUN go build -o /myapp

# When a container is run from this image, run the binary
CMD /myapp