import { writable } from "svelte/store";
import api from '../helpers/api';

const trees = writable([]);

const getTreesForUser = async (userId) => {
    try {
        const response = await api.get('/trees', { params: { userId } });
        trees.set(response.data);
    } catch (error) {
        console.error('Failed to fetch trees:', error);
    }
};

const createTreeForUser = async (tree) => {
    try {
        const response = await api.post(`/trees`, tree);
        trees.update(currentTrees => [...currentTrees, response.data]);
    } catch (error) {
        console.error('Failed to create tree:', error);
    }
};

const updateTreeById = async (treeId, tree) => {
    try {
        const response = await api.put(`/trees/${treeId}`, tree);
        trees.update(currentTrees => currentTrees.map(t => t.id === treeId ? {...t, ...tree} : u));
    } catch (error) {
        console.error('Failed to update tree:', error);
    }
};

const deleteTreeById = async (treeId) => {
    try {
        await api.delete(`/trees/${treeId}`);
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