# Build stage
# 최신 버전으로 수정 ******
FROM golang:1.22-bullseye as deploy-builder

WORKDIR /app

# go.mod, go.sum 파일 복사 및 의존성 다운로드
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# 바이너리 빌드
RUN go build -trimpath -ldflags "-w -s" -o app

#--------------------------------
# Deployment stage

FROM debian:bullseye-slim as deploy

RUN apt-get update

# 빌드된 바이너리를 복사
COPY --from=deploy-builder /app/app .

# 실행 명령
CMD ["./app"]

#--------------------------------
# Development stage with Air for live reloading

# ******** 최신 버전으로 수정 ***********
FROM golang:1.23.1 as dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]
