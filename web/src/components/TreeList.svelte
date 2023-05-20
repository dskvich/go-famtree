<script>
    import { onMount } from 'svelte';
    import treeStore from '../stores/treeStore';
    import Tree from './Tree.svelte'

    export let selectedUser;
    let selectedTree = null

    let newTree = {
        name: '',
        user_id: selectedUser.id
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        await treeStore.createTreeForUser(selectedUser.id, newTree);
        newTree.name = '';
    };

    const handleSelect = (tree) => {
        selectedTree = tree;
    };

    onMount(() => {
        treeStore.getTreesForUser(selectedUser.id);
    });

    $: selectedUser, treeStore.getTreesForUser(selectedUser.id);
</script>

<div class="tree-container">
    <h1>Trees for {selectedUser.name} </h1>

    <form on:submit|preventDefault={handleSubmit}>
        <label for="treeName">Tree Name:</label>
        <input id="treeName" bind:value={newTree.name} required />
        <button type="submit">Add Tree</button>
    </form>

    {#each $treeStore as tree (tree.id)}
        <div class="tree" on:click={() => handleSelect(tree)}>
            <h2>{tree.name}</h2>
        </div>

        <Tree treeId={tree.id} rootId={tree.root_id}/>
    {/each}
</div>