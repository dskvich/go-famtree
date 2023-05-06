<!-- App.svelte -->
<script>
    import { onMount } from 'svelte';

    let usersMap = {};
    let id = '';
    let login = '';
    let name = '';

    $: saveBtnText = id ? 'Modify' : 'Add';

    $: users = Object.values(usersMap)

    // Get the data from the api, after the page is mounted.
    onMount(async () => {
        const res = await getUsers();
        console.log(res)
        usersMap = res;
    });

    function clearState() {
        id = login = name = '';
    }

    const getUsers = async () => {
        try {
            const res = await fetch('/api/v1/users');
            if (res.ok) {
                const users = await res.json();
                const usersMap = users.reduce(function(map, user) {
                    map[user.id] = user;
                    return map;
                }, {});
                return usersMap;
            } else {
                const msg = await res.text();
                console.error(res.statusCode, 'Get users: ' + msg);
            }
        } catch (e) {
            console.error(e.message);
        }
    }

    const deleteUser = async (id) => {
        try {
            const options = {method: 'DELETE'};
            const res = await fetch(`/api/v1/users/${id}`, options);
            if (!res.ok) throw new Error('failed to delete dog with id ' + id);
            delete usersMap[id];
            usersMap = usersMap;
        } catch (e) {
            console.error(e.message);
        }
    }

    function editUser(dog) {
        ({login, name} = dog);
        id = dog.id;
    }

    const saveUser = async () => {
        // If `id` is set, we are updating a dog.
        // Otherwise we are creating a new dog.
        const user = {login, name, role: "USER"};
        if (id) user.id = id;

        try {
            const options = {
                method: id ? 'PUT' : 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(user)
            };

            console.log(options)

            const path = id ? `/api/v1/users/${id}` : '/api/v1/users';

            const res = await fetch(path, options);
            const result = await res.json();

            if (!res.ok) {
                throw new Error(result.message || result.statusText);
            }

            usersMap[result.id] = result;
            usersMap = usersMap;

            clearState();
        } catch (e) {
            console.error(e.message);
        }
    }
</script>
<style>
</style>
<div class="App">
    <h1>Users</h1>

    <div class="photos">
        <form>
            <input type="text" bind:value={login}/>
            <input type="text" bind:value={name}/>
            <button on:click|preventDefault={saveUser}>
                {saveBtnText}
            </button>
            {#if id}
                <button on:click|preventDefault={clearState}>Cancel</button>
            {/if}
        </form>
        <ul>
            {#each users as user}
                <li>
                    {user.login} {user.name}
                    <button class="icon-btn" on:click={() => editUser(user)}>
                        &#x270E;
                    </button>
                    <button on:click={() => deleteUser(user.id)}>
                        &#x1F5D1;
                    </button>
                </li>
            {:else}
                <p>loading...</p>
            {/each}
        </ul>
    </div>
</div>
