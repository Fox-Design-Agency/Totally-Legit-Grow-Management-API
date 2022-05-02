# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:latest as builder

# Create and change to the app directory.
WORKDIR /app

# Copy local code to the container image.
COPY . ./
# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates


#ARGS to set the ENV 
ARG db_host
ENV DBHOST=$db_host
ARG db_name
ENV DBNAME=$db_name
ARG db_user
ENV DBUSER=$db_user
ARG db_pass
ENV DBPASS=$db_pass
ARG port
ENV PORT=$port
ARG environment
ENV ENVIRONMENT=$environment

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server
COPY --from=builder /app/pkg /pkg
COPY --from=builder /app/resources /resources

EXPOSE 80

# Run the web service on container startup.
CMD ["/server", "-prod=true"]
