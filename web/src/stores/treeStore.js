import { writable } from "svelte/store";
import api from '../helpers/api';

const trees = writable([]);

const getTreesForUser = async (userId) => {
    try {
        const response = await api.get(`/users/${userId}/trees`);
        trees.set(response.data);
    } catch (error) {
        console.error('Failed to fetch trees:', error);
    }
};

const createTreeForUser = async (userId, tree) => {
    try {
        const response = await api.post(`/users/${userId}/trees`, tree);
        trees.update(currentTrees => [...currentTrees, response.data]);
    } catch (error) {
        console.error('Failed to create tree:', error);
    }
};

const updateTreeById = async (userId, treeId, tree) => {
    try {
        const response = await api.put(`/users/${userId}/trees/${treeId}`, tree);
        trees.update(currentTrees => currentTrees.map(t => t.id === treeId ? {...t, ...tree} : u));
    } catch (error) {
        console.error('Failed to update tree:', error);
    }
};

const deleteTreeById = async (userId, treeId) => {
    try {
        await api.delete(`/users/${userId}/trees/${treeId}`);
        trees.update(currentTrees => currentTrees.filter(t => t.id !== treeId));
    } catch (error) {
        console.error('Failed to delete tree:', error);
    }
};

export default {
    subscribe: trees.subscribe,
    getTreesForUser,
    createTreeForUser,
    updateTreeById,
    deleteTreeById,
};