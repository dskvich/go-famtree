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
        getUsers();
    });

    function clearState() {
        id = login = name = '';
    }

    const getUsers = async () => {
        try {
            const resp = await fetch('/api/v1/users');
            if (!resp.ok) {
                const msg = await resp.text();
                console.error(resp.statusCode, 'getting users: ' + msg);
            }

            const result = await resp.json();
            usersMap = result.reduce(function(map, user) {
                map[user.id] = user;
                return map;
            }, {});
            usersMap = usersMap;
        } catch (e) {
            console.error(e.message);
        }
    }

    const deleteUser = async (id) => {
        try {
            const resp = await fetch(`/api/v1/users/${id}`, {method: 'DELETE'});
            if (!resp.ok) {
                const msg = await resp.text();
                console.error(resp.statusCode, 'deleting users: ' + msg);
            }

            delete usersMap[id];
            usersMap = usersMap;
        } catch (e) {
            console.error(e.message);
        }
    }

    function editUser(user) {
        ({id, login, name} = user);
    }

    const saveUser = async () => {
        try {
            const user = {id, login, name};
            const options = {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(user)
            };

            const resp = await fetch('/api/v1/users', options);
            if (!resp.ok){
                const msg = await resp.text();
                console.error(resp.statusCode, 'saving users: ' + msg);
            }

            const result = await resp.json();
            usersMap[result.id] = result;
            usersMap = usersMap;

            clearState();
        } catch (e) {
            console.error(e.message);
        }
    }

    const updateUser = async () => {
        try {
            const user = {login, name};
            const options = {
                method: 'PUT',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(user)
            };

            const resp = await fetch(`/api/v1/users/${id}`, options);
            if (!resp.ok) {
                const msg = await resp.text();
                console.error(resp.statusCode, 'updating users: ' + msg);
            }

            usersMap[id] = {id, login, name};
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
            <button on:click|preventDefault={id?updateUser:saveUser}>
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
