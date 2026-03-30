function TaskPreview({ task, tasks, setTasks, setLoading, onTaskClick, onSetDone }) {
  return (
    <>
      <div onClick={() => onTaskClick(task)} style={{ cursor: 'pointer' }}>
        <h2>{task.title}</h2>
        <p>{task.description}</p>
      </div>

      <button onClick={() => {
        onSetDone(task.id);
        setLoading(true);
        
        setTasks(tasks.filter(t => t.id !== task.id));
        setLoading(false);
      }}>Set Done</button>
    </>
  )
}

function TaskFullview({ task, onClose }) {
  if (!task) return null;

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <button onClick={onClose} style={{ float: 'right' }}>X</button>
        <h2>{task.title} (Visualização Completa)</h2>
        <p><strong>Descrição:</strong> {task.description}</p>
        <p><strong>Detalhes:</strong> Informações extras carregadas aqui...</p>
      </div>
    </div>
  );
}

function CreateTaskForm({ create, onCreate, onClose}) {
  if (!create) return null;
  
  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <button onClick={onClose} style={{ float: 'right' }}>X</button>
        <h2>Create Task</h2>
        <form onSubmit={onCreate}>

          <div>
            <label>Title:</label>
            <input type="text" name="title" required />
          </div>

          <div>
            <label>Description:</label>
            <textarea name="description" required></textarea>
          </div>

          <div>
            <label>Degree of Difficulty</label>
            <select name="difficulty" required>
              <option value="1">1</option>
              <option value="2">2</option>
              <option value="3">3</option>
              <option value="4">4</option>
              <option value="5">5</option>
            </select>
          </div>

          <div>
            <label>Degree of Importance</label>
            <select name="importance" required>
              <option value="1">1</option>
              <option value="2">2</option>
              <option value="3">3</option>
              <option value="4">4</option>
              <option value="5">5</option>
            </select>
          </div>

          <div>
            <label>Deadline:</label>
            <input type="date" name="deadline" />
          </div>

          <div>
            <label>Repeatable:</label>
            <input type="checkbox" name="repeatable" />
          </div>

          <button type="submit">Enter</button>
        </form>
      </div>
    </div>
  )
}

function TasksContainer({ tasks, setTasks, onTaskClick, onSetDone, setLoading }) {
    return (
        <div>
            {tasks.map((task) => (
              <TaskPreview task={task} tasks={tasks} setTasks={setTasks} onTaskClick={onTaskClick} onSetDone={onSetDone} setLoading={setLoading} />
            ))}
        </div>
    )
}

export { TaskPreview, TaskFullview, TasksContainer, CreateTaskForm }