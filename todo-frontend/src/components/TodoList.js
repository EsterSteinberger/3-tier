import React, { useState, useEffect } from 'react';
import TodoItem from './TodoItem';
import TodoForm from './TodoForm';
import { TodoService } from '../services/api';

const TodoList = () => {
  const [todos, setTodos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Fetch all todos from the API
  const fetchTodos = async () => {
    try {
      setLoading(true);
      const data = await TodoService.getAllTodos();
      setTodos(data);
      setError(null);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  // Add a new todo
  const addTodo = async (title) => {
    try {
      const newTodo = await TodoService.createTodo(title);
      setTodos([...todos, newTodo]);
    } catch (err) {
      setError(err.message);
    }
  };

  // Toggle todo completion status
  const toggleTodo = async (id, completed) => {
    try {
      const updatedTodo = await TodoService.updateTodo(id, { completed: !completed });
      setTodos(
        todos.map((todo) => (todo.id === id ? updatedTodo : todo))
      );
    } catch (err) {
      setError(err.message);
    }
  };

  // Delete a todo
  const deleteTodo = async (id) => {
    try {
      await TodoService.deleteTodo(id);
      setTodos(todos.filter((todo) => todo.id !== id));
    } catch (err) {
      setError(err.message);
    }
  };

  // Fetch todos when component mounts
  useEffect(() => {
    fetchTodos();
  }, []);

  if (loading && todos.length === 0) {
    return <div className="loading">Loading...</div>;
  }

  if (error) {
    return <div className="error">Error: {error}</div>;
  }

  return (
    <div className="todo-list">
      <h1>TODO List</h1>
      <TodoForm addTodo={addTodo} />
      {todos ? (
        todos.length === 0 ? (
          <p className="empty-list">No todos yet. Add one above!</p>
        ) : (
          todos.map((todo) => (
            todo && <TodoItem
              key={todo.id}
              todo={todo}
              toggleTodo={toggleTodo}
              deleteTodo={deleteTodo}
            />
          ))
        )
      ) : (
        <p>Loading todos...</p>
      )}
    </div>
  );
};

export default TodoList; 