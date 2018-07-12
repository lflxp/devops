import http from 'axios'

const ip = 'http://localhost:9090'
const url = ip + '/v1/test/test'

export default {
  getDefault() {
    console.log('ok')
    return http.get(url)
  }
}
