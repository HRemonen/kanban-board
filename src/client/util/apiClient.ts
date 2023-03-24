import axios from 'axios'

const apiClient = axios.create({ baseURL: `localhost:8080/api/v1` })

export default apiClient
