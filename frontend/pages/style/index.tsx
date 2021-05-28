// @ts-nocheck

import React, { useEffect } from 'react'
import { useRouter } from 'next/router'
import Link from 'next/link'
import Peer from 'skyway-js'

export const Home = (): JSX.Element => {
  const router = useRouter()
  const { id } = router.query

  const myVideo = React.createRef()
  const myPeerId = React.createRef()

  const targetVideo = React.createRef()
  const targetPeerId = React.createRef()

  const [agenda, setAgenda] = React.useState<Agenda[]>([
    {
      title: '自己紹介',
      name: 'ひでよし',
      time: 1,
    },
    {
      title: '自己紹介',
      name: 'うっしー',
      time: 1,
    },
    {
      title: '自己紹介',
      name: 'かわかみ',
      time: 1,
    },
    {
      title: 'プレゼンテーション',
      name: 'かわかみ',
      time: 1,
    },
    {
      title: '終わりのあいさつ',
      name: 'ひでよし',
      time: 1,
    },
  ])

  const peer = new Peer({
    key: '79268aaf-5013-43eb-9147-0b8d3c557bca',
    debug: 3,
  })

  useEffect(() => {
    let localStream: MediaStream

    // 開始処理
    peer.on('open', () => {
      myPeerId.current.innerText = `My Peer id: ${peer.id}`
    })

    // 着信処理
    peer.on('call', (mediaConnection) => {
      mediaConnection.answer(localStream)
      setEventListener(mediaConnection)
      mediaConnection.on('stream', async (stream) => {
        targetVideo.current.srcObject = stream
      })
    })

    // 自分のカメラ映像取得
    navigator.mediaDevices
      .getUserMedia({ video: true, audio: true })
      .then((stream) => {
        // 成功時にvideo要素にカメラ映像をセットし、再生
        myVideo.current.srcObject = stream
        // 着信時に相手にカメラ映像を返せるように、グローバル変数に保存しておく
        localStream = stream
      })
      .catch((error) => {
        // 失敗時にはエラーログを出力
        console.error('mediaDevice.getUserMedia() error:', error)
        return
      })

    document.getElementById('make-call').onclick = () => {
      const targetId = targetPeerId.current.value
      const mediaConnection = peer.call(targetId, localStream)
      mediaConnection.on('stream', async (stream) => {
        targetVideo.current.srcObject = stream
        await targetVideo.current.play().catch(console.error)
      })
    }
  })

  return (
    <div className="items-center justify-center min-h-screen p-5">
      <h1 className="text-5xl font-bold">Welcome to Room {id} </h1>
      <div className="flex flex-row mt-3">
        <div className="flex-1 bg-gray-100 mx-3">
          <video
            id="my_video"
            ref={myVideo}
            width="400px"
            autoPlay={true}
            muted
            playsInline
            className="mx-auto"
          />
          <video
            id="target_video"
            ref={targetVideo}
            width="400px"
            autoPlay={true}
            playsInline
            className="mx-auto"
          />
        </div>
        <div className="flex-1">
          <Link href="/">トップへ戻る</Link>
          <h3 className="text-xl font-bold mb-8">このルームのアジェンダ</h3>
          <div className="flex flex-col">
            <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
              <div className="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
                <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
                  <table className="min-w-full divide-y divide-gray-200">
                    <thead className="bg-gray-50">
                      <tr>
                        <th
                          scope="col"
                          className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                        >
                          Title
                        </th>
                        <th
                          scope="col"
                          className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                        >
                          スピーカー
                        </th>
                        <th
                          scope="col"
                          className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                        >
                          時間
                        </th>
                      </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-gray-200">
                      {agenda.map((a, index) => (
                        <tr key={index}>
                          <td className="px-6 py-4 whitespace-nowrap">
                            <div className="text-sm text-gray-900">
                              {a.title}
                            </div>
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap">
                            <div className="text-sm text-gray-900">
                              {a.name}
                            </div>
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap">
                            <div className="text-sm text-gray-500">
                              {a.time}分
                            </div>
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>
          <div className="mt-3">
            <p ref={myPeerId} />
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 mt-2 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              placeholder="相手のpeer idを入力してください"
              ref={targetPeerId}
            />
            <button
              id="make-call"
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 mt-2 rounded focus:outline-none focus:shadow-outline"
              type="button"
            >
              発信
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Home
