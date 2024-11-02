# Build stage
FROM golang:1.18.2-bullseye as deploy-builder

WORKDIR /app

# go.mod, go.sum 파일 복사 및 의존성 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 애플리케이션 소스 파일 복사, main만 복사
COPY chapter14/section56/main.go ./chapter14/section56/

# 바이너리 빌드
RUN go build -o app ./chapter14/section56/main.go

#--------------------------------
# Deployment stage
FROM debian:bullseye-slim as deploy

WORKDIR /app

# 빌드된 바이너리를 복사
COPY --from=deploy-builder /app/app .

# 실행 명령
CMD ["./app"]

#--------------------------------
# Development stage with Air for live reloading
FROM golang:1.23 as dev

WORKDIR /app
COPY --from=deploy-builder /app .
RUN go install github.com/air-verse/air@latest
CMD ["air"]
