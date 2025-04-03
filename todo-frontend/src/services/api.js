import { API_CONFIG } from '../config';

const API_BASE_URL = API_CONFIG.BASE_URL;

export const TodoService = {
  getAllTodos: async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/todos`);
      if (!response.ok) throw new Error('Failed to fetch todos');
      const data = await response.json();
      return data || []; // Return empty array if data is null
    } catch (error) {
      console.error('Error fetching todos:', error);
      return []; // Return empty array on error
    }
  },
  
  createTodo: async (title) => {
    const response = await fetch(`${API_BASE_URL}/todos`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title }),
    });
    if (!response.ok) throw new Error('Failed to create todo');
    return response.json();
  },
  
  updateTodo: async (id, updates) => {
    const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(updates),
    });
    if (!response.ok) throw new Error('Failed to update todo');
    return response.json();
  },
  
  deleteTodo: async (id) => {
    const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
      method: 'DELETE',
    });
    if (!response.ok) throw new Error('Failed to delete todo');
    return true;
  }
}; 