
function GetButton({ onClick }) {
    return (
        <button className="btn btn-primary" onClick={onClick}>
            Change Tasks
        </button>
    )
}

function PostButton({ onClick }) {
    return (
        <button className="btn btn-secondary" onClick={onClick}>
            New
        </button>
    )
}


export {GetButton, PostButton}