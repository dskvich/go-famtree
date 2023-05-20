import { writable } from 'svelte/store';
import api from '../helpers/api';

const people = writable([]);

const getPeopleByTree = async (treeId) => {
    try {
        const response = await api.get(`/trees/${treeId}/people`);
        people.set(response.data);
    } catch (error) {
        console.error('Failed to fetch people in a tree:', error);
    }
};

export default {
    subscribe: people.subscribe,
    getPeopleByTree,
};