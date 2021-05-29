import React from 'react'
import '../styles/globals.css'
import { AppProps } from 'next/app'
import Head from 'next/head'

const MyApp = ({ Component, pageProps }: AppProps): JSX.Element => {
  return (
    <>
      <Head>
        <title>Create Next App</title>
      </Head>
      <Component {...pageProps} />
    </>
  )
}

export default MyApp
