import api from './api'

const taskService = {
    loadTasks: async (setTasks, setLoading) => {
        setLoading(true);
        try {
            const data = await api.getTasks();
            setTasks(data);
        } catch (error) {
            console.error('Erro ao buscar tarefas:', error);
        } finally {
            setLoading(false);
        }
    },

    replaceTasks: async (tasks, setTasks, setLoading) => {
        try {
            for (const task of tasks) {
                await api.queueTask(task.id);
            }
            await taskService.loadTasks(setTasks, setLoading);
        } catch (error) {
            console.error('Erro ao enfileirar tarefas:', error);
        }
    }
}

export default taskService