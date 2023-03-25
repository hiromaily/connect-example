import type { NextPage, GetStaticProps } from 'next'
import * as React from 'react'
import App from '../components/app'

type IndexProps = {
  message: string
}

const Index: NextPage<IndexProps> = ({ message }) => {
  return (
      <App />
  )
}

export default Index

// For SSG
export const getStaticProps: GetStaticProps<IndexProps> = async (context) => {
  const timestamp = new Date().toLocaleString()
  const message = `this page was rendered by calling getStaticProps() at ${timestamp}`
  return {
    props: {
      message,
    },
  }
}
