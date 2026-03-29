import { useState } from 'react'
import { GetButton, PostButton } from './components/Button'
import { TaskFullview, TasksContainer, CreateTaskForm } from './components/Task'
import taskService from './services/taskService'

import './styles/App.css'

function App() {
  const [tasks, setTasks] = useState([])
  const [selectedTask, setSelectedTask] = useState(null);
  const [createTask, setCreateTask] = useState(false);
  const [loading, setLoading] = useState(false);

  return (
    <>
      <GetButton onClick={() => taskService.replaceTasks(tasks, setTasks, setLoading)}/>
      <PostButton onClick={() => setCreateTask(true)} />

      <div className="content">
        {loading ? (
          <p>Carregando tarefas...</p>
        ) : (
          <TasksContainer tasks={tasks} setTasks={setTasks} onTaskClick={setSelectedTask} onSetDone={taskService.setTaskDone} setLoading={setLoading} />
        )}
      </div>


      {/* Modal overlays section */}
      
      <TaskFullview 
        task={selectedTask} 
        onClose={() => setSelectedTask(null)} 
      />
      
      <CreateTaskForm 
        create={createTask}
        onClose={() => setCreateTask(false)}
        onCreate={(e) => taskService.createTask(e, setCreateTask)}
      />
    </>
  )
}

export default App
