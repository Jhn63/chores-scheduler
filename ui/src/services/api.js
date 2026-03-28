import axios from 'axios'

const ax = axios.create({
  baseURL: 'http://localhost:8000',
  timeout: 1000,
})

const api = {
    getTasks: async () => {
        const response = await ax.get('/tasks')
        return response.data
    },

    createTask: async (task) => {
        const response = await ax.post('/tasks', task)
        return response.data
    },

    queueTask: async (id) => {
        const response = await ax.patch(`/tasks/${id}/queue`)
        return response.data
    },

    setDone: async (id) => {
        const response = await ax.patch(`/tasks/${id}/done`)
        return response.data
    }
}

export default api