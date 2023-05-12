import { useCallback, useEffect, useState } from 'react'
import './App.styl'
import GoRTSPComponent from './components/rtspGo'
import { StreamInfo, Streams, infoType } from './interface'
import { Button, Divider, Row, Col } from 'antd'

function App() {
  const baseUrl = 'http://127.0.0.1:8083'

  const [infos, setInfos] = useState<infoType[]>([])
  const [names, setNames] = useState<string[]>([])
  const [button, setButton] = useState<string>('All')

  useEffect(() => {
    const fetchPromise = fetch(baseUrl + '/streams')
    fetchPromise
      .then((res: Response) => {
        return res.json()
      })
      .then((data) => {
        const streams = (data && data.streams) as any[]

        if (streams) {
          const composite_result: Map<keyof StreamInfo, string>[] = []
          for (const uuid in streams) {
            const result = new Map<keyof StreamInfo, string>()
            const uuid_v = streams[uuid] as Streams
            result.set('uuid', uuid)
            result.set('name', uuid_v.name)
            result.set('channel', '0')
            result.set('url', uuid_v.channels['0'].url)
            composite_result.push(result)
            setNames((prev: string[]) => {
              if (!prev.includes(uuid_v.name)) {
                return [...prev, uuid_v.name]
              }
              return prev
            })
          }

          setInfos(composite_result)
        }
      })
      .catch((e: Error) => {
        console.error(e)
      })
  }, [])

  const selectButton = useCallback(
    (btn: string, infos: infoType[]) => {
      if (btn === 'All') {
        return (
          <>
            {infos.map((item: infoType, idx: number): JSX.Element => {
              return (
                <div key={idx} className="video">
                  <Divider>{item.get('name')}</Divider>
                  <GoRTSPComponent infos={item} key={idx} />
                </div>
              )
            })}
          </>
        )
      } else {
        for (const idx in infos) {
          const info = infos[idx]
          if (btn === info.get('name')) {
            return (
              <>
                <GoRTSPComponent infos={info} key={idx} />
              </>
            )
          }
        }
      }
    },
    [button],
  )

  return (
    <div className="App">
      <div className="button-area">
        <title className="button-title">按钮区</title>
        <Divider></Divider>
        <Button
          className={button === 'All' ? 'active' : ''}
          type="primary"
          block
          style={{ width: '50%' }}
          onClick={() => setButton('All')}
        >
          All
        </Button>
        <Divider></Divider>
        <div className="button-row">
          {names.map((name: string, idx: number) => {
            return (
              <Button
                className={name === button ? 'active' : ''}
                type="primary"
                key={idx}
                block
                onClick={() => setButton(name)}
              >
                {name}
              </Button>
            )
          })}
        </div>
      </div>

      <div
        className="video-area"
        style={
          button === 'All'
            ? { justifyContent: 'space-between' }
            : { justifyContent: 'center' }
        }
      >
        {button && selectButton(button, infos)}
      </div>
    </div>
  )
}

export default App
