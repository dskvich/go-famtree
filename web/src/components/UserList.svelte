<script>
    import {onMount} from 'svelte';
    import userStore from '../stores/userStore';
    import TreeList from './TreeList.svelte';
    import UserForm from './UserForm.svelte';

    let editUser = null;
    let selectedUser = null;

    const handleSelect = (user) => {
        selectedUser = user;
    };

    const handleEdit = (user) => {
        editUser = user;
    };

    const handleDelete = async (id) => {
        await userStore.deleteUser(id);
    };

    onMount(() => {
        userStore.fetchUsers();
    });
</script>

<style>
    @import 'bulma/css/bulma.min.css';
</style>

<div className="container">
    <div className="grid">
        <div className="grid-item">
            <h1>Users</h1>

            <UserForm {editUser}/>

            <table className="table">
                <thead>
                <tr>
                    <th>Name</th>
                    <th>Login</th>
                    <th>Email</th>
                    <th>Actions</th>
                </tr>
                </thead>
                <tbody>
                {#each $userStore as user (user.id)}
                    <tr on:click={() => handleSelect(user)}>
                        <td>{user.name}</td>
                        <td>{user.login}</td>
                        <td>{user.email}</td>
                        <td>
                            <button class="edit-button" on:click={(e) => {e.stopPropagation(); handleEdit(user)}}>Edit</button>
                            <button class="delete-button" on:click={(e) => {e.stopPropagation(); handleDelete(user.id)}}>Delete</button>
                        </td>
                    </tr>
                {/each}
                </tbody>
            </table>
        </div>

        <div class="grid-item">
            {#if selectedUser}
                <TreeList {selectedUser}/>
            {/if}
        </div>
    </div>
</div>

