function TaskPreview({ title, description }) {
  return (
    <div>
      <h2>{title}</h2>
      <p>{description}</p>
    </div>
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

function TasksContainer({ tasks, onTaskClick }) {
    return (
        <div>
            {tasks.map((task) => (
                <div key={task.id} onClick={() => onTaskClick(task)} style={{ cursor: 'pointer' }}>
                    <TaskPreview title={task.title} description={task.description} />
                </div>
            ))}
        </div>
    )
}

export { TaskPreview, TaskFullview, TasksContainer }