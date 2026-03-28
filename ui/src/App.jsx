import { useState, useEffect } from 'react'
import { GetButton, PostButton } from './components/Button'
import { TaskFullview, TasksContainer } from './components/Task'
import api from './services/api'

import './styles/App.css'

const loadTasks = async (setTasks, setLoading) => {
  setLoading(true);
  try {
    const data = await api.getTasks();
    console.log("Tarefas carregadas:", data);
    setTasks(data);
  } catch (error) {
    console.error('Erro ao buscar tarefas:', error);
  } finally {
    setLoading(false);
  }
};

function App() {
  const [tasks, setTasks] = useState([])
  const [selectedTask, setSelectedTask] = useState(null);
  const [loading, setLoading] = useState(false);


  useEffect(() => {
    loadTasks(setTasks, setLoading);
  }, []); 

  

  return (
    <>
      <GetButton />
      <PostButton />

      <div className="content">
        {loading ? (
          <p>Carregando tarefas...</p>
        ) : (
          <TasksContainer tasks={tasks} onTaskClick={setSelectedTask} />
        )}
      </div>

      {selectedTask && (
        <TaskFullview 
          task={selectedTask} 
          onClose={() => setSelectedTask(null)} 
        />
      )}
    </>
  )
}

export default App
