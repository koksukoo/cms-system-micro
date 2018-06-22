const express = require('express')
const next = require('next')
const fetch = require('isomorphic-unfetch')

const dev = process.env.NODE_ENV !== 'production'
const app = next({ dev })
const handle = app.getRequestHandler()

app.prepare()
.then(() => {
  const server = express()
  // todo: change to manageable implementation
  server.get('/fetch/pages/', (req, res) => {
      fetch(`${process.env.ENGINE_URL}/projects/5b1c2e345d9b1d61551da093/pages/`).then(resp => {
            resp.json().then(r => res.send(r))
      })
  })

  server.get('*', (req, res) => {
    return handle(req, res)
  })

  server.listen(3000, (err) => {
    if (err) throw err
    console.log('> Ready on http://localhost:3000')
  })
})
.catch((ex) => {
  console.error(ex.stack)
  process.exit(1)
})