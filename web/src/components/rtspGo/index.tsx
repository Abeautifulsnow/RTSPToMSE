import { Divider } from 'antd'
import { useEffect, useRef } from 'react'
import './index.styl'

function Utf8ArrayToStr(array: Uint8Array) {
  let out, i, c
  let char2, char3
  out = ''
  const len = array.length
  i = 0
  while (i < len) {
    c = array[i++]
    switch (c >> 4) {
      case 7:
        out += String.fromCharCode(c)
        break
      case 13:
        char2 = array[i++]
        out += String.fromCharCode(((c & 0x1f) << 6) | (char2 & 0x3f))
        break
      case 14:
        char2 = array[i++]
        char3 = array[i++]
        out += String.fromCharCode(
          ((c & 0x0f) << 12) | ((char2 & 0x3f) << 6) | ((char3 & 0x3f) << 0),
        )
        break
    }
  }
  return out
}

export default function GoRTSPComponent(props: any) {
  const videoRef = useRef<HTMLVideoElement | null>(null)

  let mseSourceBuffer: SourceBuffer,
    mseStreamingStarted = false,
    videoSound = false

  const mseQueue: any[] = []

  const pushPacket = () => {
    if (!mseSourceBuffer.updating) {
      if (mseQueue.length > 0) {
        const packet = mseQueue.shift()
        mseSourceBuffer.appendBuffer(packet)
      } else {
        mseStreamingStarted = false
      }
    }

    if (videoRef.current) {
      if (videoRef.current.buffered.length > 0) {
        if (
          typeof document.hidden !== 'undefined' &&
          document.hidden &&
          !videoSound
        ) {
          //no sound, browser paused video without sound in background
          videoRef.current.currentTime =
            videoRef.current.buffered.end(
              videoRef.current.buffered.length - 1,
            ) - 0.5
        }
      }
    }
  }

  const readPacket = (packet: ArrayBuffer) => {
    if (!mseStreamingStarted) {
      mseSourceBuffer.appendBuffer(packet)
      mseStreamingStarted = true
      return
    }
    mseQueue.push(packet)
    if (!mseSourceBuffer.updating) {
      pushPacket()
    }
  }

  const startPlay = () => {
    let protocol, port
    location.protocol == 'https:' ? (protocol = 'wss') : (protocol = 'ws')
    location.protocol == 'https:' ? (port = '8444') : (port = '8083')
    const url = `${protocol}://${location.hostname}:${port}/stream/27aec28e-6181-4753-9acd-0456a75f0289/channel/0/mse`

    const mse = new MediaSource()
    if (videoRef.current) {
      videoRef.current.src = window.URL.createObjectURL(mse)
    }
    mse.addEventListener(
      'sourceopen',
      function () {
        const ws = new WebSocket(url)
        ws.binaryType = 'arraybuffer'
        ws.onopen = function (event) {
          console.log('Connect to wss')
        }
        ws.onmessage = function (event) {
          const data = new Uint8Array(event.data)
          if (data[0] == 9) {
            const decoded_arr = data.slice(1)
            let mimeCodec
            if (window.TextDecoder) {
              mimeCodec = new TextDecoder('utf-8').decode(decoded_arr)
            } else {
              mimeCodec = Utf8ArrayToStr(decoded_arr)
            }
            if (mimeCodec.indexOf(',') > 0) {
              videoSound = true
            }
            mseSourceBuffer = mse.addSourceBuffer(
              'video/mp4; codecs="' + mimeCodec + '"',
            )
            mseSourceBuffer.mode = 'segments'
            mseSourceBuffer.addEventListener('updateend', pushPacket)
          } else {
            readPacket(event.data)
          }
        }
        ws.onerror = (ev: Event) => {
          console.error('ws error:', ev)
        }
      },
      false,
    )
  }

  useEffect(() => {
    if (videoRef.current) {
      startPlay()
    }
  }, [])

  useEffect(() => {
    if (videoRef.current) {
      videoRef.current.addEventListener('loadeddata', () => {
        videoRef.current && videoRef.current.play()
      })
      //fix stalled video in safari
      videoRef.current.addEventListener('pause', () => {
        if (
          videoRef.current &&
          videoRef.current.currentTime >
            videoRef.current.buffered.end(videoRef.current.buffered.length - 1)
        ) {
          videoRef.current.currentTime =
            videoRef.current.buffered.end(
              videoRef.current.buffered.length - 1,
            ) - 0.1
          videoRef.current.play()
        }
      })

      videoRef.current.addEventListener('error', () => {
        console.log('video_error')
      })
    }
  }, [])

  return (
    <>
      <Divider>socket调用分割线</Divider>
      <video
        ref={videoRef}
        id="videoplayer"
        autoPlay
        controls
        playsInline
        muted
      ></video>
    </>
  )
}
