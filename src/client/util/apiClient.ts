import axios from 'axios'

const v1ApiClient = axios.create({ baseURL: `/api/v1` })

export default v1ApiClient
