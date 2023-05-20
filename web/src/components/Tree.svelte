<script>
    import { onMount } from 'svelte';
    import peopleStore from '../stores/peopleStore';

    export let treeId;
    export let rootId;
    export let level = 0; // Default level is 0 (root person)

    let people = [];

    onMount(async () => {
        people = await peopleStore.getPeopleByTree(treeId);
        console.log(people)
    });
</script>

<div class="family-tree">
    {#if people.length > 0}
        <div class="tree left-tree">
            <!-- Render the mother's tree -->
            {#each people as person}
                {#if person.id === rootId}
                    <div class="person" data-id={person.id}>
                        <span class="name">{person.name}</span>
                        {#if level > 0}
                            {#each person.children as childId}
                                {#each people as childPerson}
                                    {#if childPerson.id === childId}
                                        <svelte:self {treeId} {rootId} level={level - 1} />
                                    {/if}
                                {/each}
                            {/each}
                        {/if}
                    </div>
                {/if}
            {/each}
        </div>

        <div class="tree center-tree">
            <!-- Render the root person in the center -->
            {#each people as person}
                {#if person.id === rootId}
                    <div class="person" data-id={person.id}>
                        <span class="name">{person.name}</span>
                        {#if level > 0}
                            {#each person.children as childId}
                                {#each people as childPerson}
                                    {#if childPerson.id === childId}
                                        <svelte:self {treeId} {rootId} level={level - 1} />
                                    {/if}
                                {/each}
                            {/each}
                        {/if}
                    </div>
                {/if}
            {/each}
        </div>

        <div class="tree right-tree">
            <!-- Render the father's tree -->
            {#each people as person}
                {#if person.id === rootId}
                    <div class="person" data-id={person.id}>
                        <span class="name">{person.name}</span>
                        {#if level > 0}
                            {#each person.children as childId}
                                {#each people as childPerson}
                                    {#if childPerson.id === childId}
                                        <svelte:self {treeId} {rootId} level={level - 1} />
                                    {/if}
                                {/each}
                            {/each}
                        {/if}
                    </div>
                {/if}
            {/each}
        </div>
    {/if}
</div>

<style>
    .family-tree {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
    }

    .tree {
        flex: 1;
        display: flex;
        flex-direction: column;
        align-items: center;
        margin: 0 20px;
    }

    .person {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin: 20px;
        padding: 10px;
        border: 1px solid #ddd;
        border-radius: 5px;
    }

    .name {
        font-weight: bold;
        margin-bottom: 5px;
    }

    .center-tree {
        background-color: lightblue;
    }

    .left-tree {
        text-align: right;
    }

    .right-tree {
        text-align: left;
    }
</style>
