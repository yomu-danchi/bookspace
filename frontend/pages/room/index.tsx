import React from 'react'
import Link from 'next/link'
import Agenda from '../../component/agenda'

export const Home = (): JSX.Element => {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-2">
      <main className="flex flex-col items-center justify-center w-full flex-1 px-20 text-center">
        <h3 className="text-xl font-bold mb-8">このルームのアジェンダ</h3>
        <div className="flex flex-col">
          <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
            <Agenda />
          </div>
        </div>
        <div>
          <Link href="/room/[id]" as={`/room/${1}`}>
            <a className="group relative w-full my-8 flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              ルームを作成する
            </a>
          </Link>
        </div>
        <div>
          <Link href="/">
            <a className="group relative w-full my-8 flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-gray-500 hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              もどる
            </a>
          </Link>
        </div>
      </main>
    </div>
  )
}

export default Home
