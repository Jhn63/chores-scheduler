import { useState, useEffect } from 'react'
import { GetButton, PostButton } from './components/Button'
import { TaskFullview, TasksContainer } from './components/Task'
import taskService from './services/taskService'

import './styles/App.css'

function App() {
  const [tasks, setTasks] = useState([])
  const [selectedTask, setSelectedTask] = useState(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    taskService.loadTasks(setTasks, setLoading);
  }, []); 

  return (
    <>
      <GetButton onClick={() => taskService.replaceTasks(tasks, setTasks, setLoading)}/>
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
