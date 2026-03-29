
function GetButton({ onClick }) {
    return (
        <button className="btn btn-primary" onClick={onClick}>
            Get Tasks
        </button>
    )
}

function PostButton({ onClick }) {
    return (
        <button className="btn btn-secondary" onClick={onClick}>
            Create Task
        </button>
    )
}


export {GetButton, PostButton}