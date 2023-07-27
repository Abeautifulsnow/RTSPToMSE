# RTSPToMSE
RTSP Stream to WebBrowser. Which use [MediaSource](https://developer.mozilla.org/en-US/docs/Web/API/MediaSource) way to load the video stream for playback.

## Start service from the source code

This is a Front-End separation project. Therefore you need to start service respectively.

### Go Service

```bash
cd server
# Download package
go mod tidy

# Start server
go run ./...
```

### Web Service

```bash
# Enter the directory
cd web

# Install package
pnpm i

# Start service
pnpm run dev
```

And then, you can visit it by http://localhost:5173/

## Use Docker-compose to start service

- Step one - Build docker images manually
  - `docker compose -f Docker-compose.yml build`
- Step two - Run these containers
  - `docker compose -f Docker-compose.yml up -d`
- Step three - Visit it on your browser
  - `http://localhost:8081`

