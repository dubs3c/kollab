# Stage 1: Build the Svelte app
FROM node:18-bookworm-slim as build-svelte
WORKDIR /app
COPY frontend/ ./
RUN npm install
RUN npm run build

# Stage 2: Build the Go application
FROM golang:1.21.1-bookworm as build-go
WORKDIR /app

# Create a clean GOPATH
ENV GOPATH /app/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV GO111MODULE on

COPY go.mod go.sum ./
RUN go mod download
COPY *.go .
RUN go build -o kollab *.go

# Stage 3: Serve the Svelte app using NGINX
FROM nginx:latest
COPY --from=build-svelte /app/build /etc/nginx/html/mgmt
COPY --from=build-go /app/kollab /bin/kollab
COPY ./nginx-config/* /etc/nginx/
RUN apt-get update && apt-get install -y procps net-tools supervisor
EXPOSE 80

# Copy the supervisor configuration file
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Start supervisor
CMD ["/usr/bin/supervisord", "-n", "-c", "/etc/supervisor/conf.d/supervisord.conf"]