import axios from 'axios'

const ax = axios.create({
  baseURL: 'http://localhost:8000',
  timeout: 100000,
})

const api = {
    getTasks: async () => {
        const response = await ax.get('/tasks')
        return response.data
    },

    postTask: async (task) => {
        const response = await ax.post('/tasks', task)
        return response.data
    },
}

export default api