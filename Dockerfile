FROM golang:1.20-alpine
# Create working directory
WORKDIR /app
# Copy go.mod and go.sum into /app
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download
# Create directories
RUN mkdir -p ./cmd/web
RUN mkdir ./data
RUN mkdir ./internal
RUN mkdir -p ./ui/html/pages
RUN mkdir -p ui/static/css
RUN mkdir -p ui/static/pdf
# Copy the data into the directories
COPY ./cmd/web/*.go ./cmd/web
COPY ./data/data.go ./data
COPY ./internal/fetchPosts.go ./internal
COPY ./ui/html/base.tmpl ./ui/html
COPY ./ui/html/pages/home.tmpl ./ui/html/pages
COPY ./ui/html/pages/project.tmpl ./ui/html/pages
COPY ./ui/static/css/main.css ./ui/static/css
COPY ./ui/static/pdf/Brad-Preston-Resume-2023.pdf ./ui/static/pdf

RUN cd ./cmd/web && CGO_ENABLED=0 GOOS=linux go build -o /test-website

CMD ["/test-website"]
