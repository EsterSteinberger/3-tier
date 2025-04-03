import React from 'react';

const TodoItem = ({ todo, toggleTodo, deleteTodo }) => {
  if (!todo) {
    return <div>Invalid todo item</div>;
  }

  return (
    <div className={`todo-item ${todo.completed ? 'completed' : ''}`}>
      <input
        type="checkbox"
        checked={todo.completed}
        onChange={() => toggleTodo(todo.id, todo.completed)}
        className="todo-checkbox"
      />
      <span className="todo-title">{todo.title}</span>
      <button 
        onClick={() => deleteTodo(todo.id)} 
        className="todo-delete-button"
      >
        Delete
      </button>
    </div>
  );
};

export default TodoItem; 