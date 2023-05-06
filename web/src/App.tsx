import { useEffect, useState } from 'react'
import './App.css'
import GoRTSPComponent from './components/rtspGo'

function App() {
  const baseUrl = 'http://127.0.0.1:8083'

  const [info, setInfo] = useState<any[]>([])

  useEffect(() => {
    const fetchPromise = fetch(baseUrl + '/streams')
    fetchPromise
      .then((res: Response) => {
        return res.json()
      })
      .then((data) => {
        console.log(data)
        const streams = (data && data.streams) as any[]

        const uuids = streams.reduce(() => {}, [])
        setInfo((prev: any[]) => [...prev, data])
      })
      .catch((e: Error) => {
        console.error(e)
      })
  }, [])

  return (
    <div className="App">
      <GoRTSPComponent />
    </div>
  )
}

export default App
