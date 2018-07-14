import App, {Container} from 'next/app'
import dynamic from 'next/dynamic'
import React from 'react'

/* VENDORS */
import 'codemirror/lib/codemirror.css'

// As a workararound next.js dynamic bug we import dynamics on every page
// todo: this will be fixed when webpack 4 arrives with next
dynamic(import('components/Codearea'), { ssr: false })


export default class MyApp extends App {
  static async getInitialProps ({ Component, router, ctx }) {
    let pageProps = {}

    if (Component.getInitialProps) {
      pageProps = await Component.getInitialProps(ctx)
    }

    return {pageProps}
  }

  render () {
    const {Component, pageProps} = this.props
    return <Container>
      <Component {...pageProps} />
    </Container>
  }
}