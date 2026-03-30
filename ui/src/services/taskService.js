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
    },

    createTask: async (taskForm, setCreatedTask) => {
        taskForm.preventDefault();
        const formData = new FormData(taskForm.target);
        const task = {
            title: formData.get('title'),
            description: formData.get('description'),
            degree_of_difficulty: parseInt(formData.get('difficulty')),
            degree_of_importance: parseInt(formData.get('importance')),
            deadline: formData.get('deadline') === '' ? null : `${formData.get('deadline')}T23:59:00Z`,
            repeatable: formData.get('repeatable') === 'on' ? true : false
        };
        
        try {
            const newTask = await api.createTask(task);
            setCreatedTask(false);
            return newTask;
        } catch (error) {
            console.error('Erro ao criar tarefa:', error);  
        }
        
    },

    setTaskDone: async (id) => {
        try {
            await api.setDone(id);
        } catch (error) {
            console.error('Erro ao marcar tarefa como concluída:', error);
        }
    }
}

export default taskService