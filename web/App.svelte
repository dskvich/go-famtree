<!-- App.svelte -->
<script>
    import { onMount } from 'svelte';

    let users = [];

    onMount(async () => {
        const res = await fetch(`/api/users`);
        users = await res.json();
    });
</script>
<style>
</style>
<div class="App">
    <h1>Users</h1>

    <div class="photos">
        <ul>
            {#each users as user}
                <li>
                    {user.first_name} {user.last_name}
                </li>
            {:else}
                <p>loading...</p>
            {/each}
        </ul>
    </div>

</div>
