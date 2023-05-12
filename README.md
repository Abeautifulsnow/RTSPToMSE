# RTSPToMSE
RTSP Stream to WebBrowser. Which use [MediaSource](https://developer.mozilla.org/en-US/docs/Web/API/MediaSource) way to load the video stream for playback.

## How to Use

This is a Front-End separation project. Therefore you need to start service respectively.

### Go Service

```bash
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
