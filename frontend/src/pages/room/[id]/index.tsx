// @ts-nocheck

import React, { useEffect } from 'react'
import { useRouter } from 'next/router'
import Link from 'next/link'
import Peer from 'skyway-js'
import Agenda from '../../../component/agenda'

export const Home = (): JSX.Element => {
  const router = useRouter()
  const { id } = router.query

  const myVideo = React.createRef()
  const myPeerId = React.createRef()

  const targetVideo = React.createRef()
  const targetPeerId = React.createRef()

  const peer = new Peer({
    key: '79268aaf-5013-43eb-9147-0b8d3c557bca',
    debug: 3,
  })

  useEffect(() => {
    let localStream

    // 開始処理
    peer.on('open', async () => {
      const interval = setInterval(function() {
        if (myPeerId && myPeerId.current) {
          myPeerId.current.textContent = `My Peer id: ${peer.id}`
          clearInterval(interval)
        }
      }, 500)
    })

    // 着信処理
    peer.on('call', (mediaConnection) => {
      mediaConnection.answer(localStream)
      setEventListener(mediaConnection, 'target')
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
      setEventListener(mediaConnection, 'me')
    }

    const setEventListener = async (mediaConnection, from) => {
      mediaConnection.on('stream', async (stream) => {
        // video要素にカメラ映像をセットして再生
        targetVideo.current.srcObject = stream
        await targetVideo.current.play().catch(console.error)
      })
      if (from === 'me') {
        await setMyTerm()
        await timeout(3000)
      }
      await setTargetTerm()
      await timeout(3000)
      await setMyTerm()
      await timeout(3000)
      await setTargetTerm()
      await timeout(3000)
      await setMyTerm()
      await timeout(3000)
      await setTargetTerm()
      await timeout(3000)
      if (from === 'target') {
        await setMyTerm()
        await timeout(3000)
      }
      await clearTerm()
    }

    const timeout = function (ms) {
      return new Promise((resolve) => setTimeout(resolve, ms))
    }

    const setTargetTerm = () => {
      myVideo.current.style.opacity = '0.5'
      myVideo.current.style.border = 'none'
      targetVideo.current.style.opacity = '1'
      targetVideo.current.style.border = '5px solid #14c484'
      targetVideo.current.muted = false
    }

    const setMyTerm = () => {
      myVideo.current.style.opacity = '1'
      myVideo.current.style.border = '5px solid #14c484'
      targetVideo.current.style.opacity = '0.5'
      targetVideo.current.style.border = 'none'
      targetVideo.current.muted = true
    }

    const clearTerm = () => {
      myVideo.current.style.opacity = '1'
      myVideo.current.style.border = 'none'
      targetVideo.current.style.opacity = '1'
      targetVideo.current.style.border = 'none'
      targetVideo.current.muted = false
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
              <Agenda />
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
