import React from 'react'
import Document, { Html, Head, Main, NextScript } from 'next/document'

class AppDocument extends Document {
  render(): JSX.Element {
    return (
      <Html>
        <Head></Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    )
  }
}

export default AppDocument
