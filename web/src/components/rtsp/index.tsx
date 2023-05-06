import React, { useEffect, useState } from 'react'
import { io } from 'socket.io-client'
import { message, Divider } from 'antd'

export default function RTSPCom(props: any) {
  const host = 'wss://local.device.zwszsports.com:8444'
  const sio = io(host, { autoConnect: true, retries: 3 })

  useEffect(() => {
    sio.connect()

    return () => {
      sio.disconnect()
    }
  }, [])

  const [eventName, setEventName] = useState('')
  const [frame, setFrame] = useState('')

  const getInputValue = (e: any) => {
    e && e.target && setEventName((e.target as any).value)
  }

  const emitEvent = () => {
    if (eventName) {
      try {
        sio.emit(eventName, () => {
          console.log(`trigger event: ${eventName}`)
        })
        message.success(`${eventName}事件发送成功`)
      } catch (error: any) {
        message.error(error)
      }
    }
  }
  function renderFrame(frame: any) {
    setFrame(frame)
  }

  useEffect(() => {
    const canvas = document.getElementById('canvas') as HTMLCanvasElement
    const ctx = canvas && canvas.getContext('2d')

    async function renderFrameNext() {
      if (frame) {
        const imageblob = new Blob([new Uint8Array(frame as any)])
        const imagebitmap = await createImageBitmap(imageblob)
        ctx!.drawImage(imagebitmap, 0, 0, canvas.width, canvas.height)
        setFrame('')
      }
      requestAnimationFrame(renderFrameNext)
    }

    sio.on('composed_frame', (payload: any) => {
      if (payload) {
        renderFrame(payload)
      }
    })

    requestAnimationFrame(renderFrameNext)
  })

  return (
    <>
      <Divider>socket调用分割线</Divider>
      {frame && <canvas id="canvas" width="1440" height="720"></canvas>}
    </>
  )
}
