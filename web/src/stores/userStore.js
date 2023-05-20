import { writable } from 'svelte/store';
import api from '../helpers/api';

const users = writable([]);

const fetchUsers = async () => {
    try {
        const response = await api.get('/users');
        users.set(response.data);
    } catch (error) {
        console.error('Failed to fetch users:', error);
    }
};

const createUser = async (user) => {
    try {
        const response = await api.post('/users', user);
        users.update(currentUsers => [...currentUsers, response.data]);
    } catch (error) {
        console.error('Failed to create user:', error);
    }
};

const updateUser = async (id, user) => {
    try {
        await api.put(`/users/${id}`, user);
        users.update(currentUsers => currentUsers.map(u => u.id === id ? {...u, ...user} : u));
    } catch (error) {
        console.error('Failed to update user:', error);
    }
};

const deleteUser = async (id) => {
    try {
        await api.delete(`/users/${id}`);
        users.update(currentUsers => currentUsers.filter(u => u.id !== id));
    } catch (error) {
        console.error('Failed to delete user:', error);
    }
};

export default {
    subscribe: users.subscribe,
    fetchUsers,
    createUser,
    updateUser,
    deleteUser,
};
