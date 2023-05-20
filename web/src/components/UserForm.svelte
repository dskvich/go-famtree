<script>
    import userStore from '../stores/userStore';

    let formUser = {
        login: '',
        name: '',
        email: '',
        password: ''
    };
    export let editUser;

    const handleSubmit = async (event) => {
        event.preventDefault();
        if (editUser) {
            await userStore.updateUser(editUser.id, formUser);
        } else {
            await userStore.createUser(formUser);
        }
        resetForm();
    };

    const resetForm = () => {
        formUser = {
            login: '',
            name: '',
            email: '',
            password: ''
        };
        editUser = null;
    };

    const populateForm = (user) => {
        formUser = { ...user };
    }

    $: editUser, populateForm(editUser);
</script>

<h2>{editUser ? 'Edit' : 'Create'} User</h2>

<form class="user-form" on:submit|preventDefault={handleSubmit}>
    <div class="form-row">
        <div class="form-field">
            <label for="login">Login:</label>
            <input id="login" bind:value={formUser.login} required />
        </div>
        <div class="form-field">
            <label for="name">Name:</label>
            <input id="name" bind:value={formUser.name} required />
        </div>
        <div class="form-field">
            <label for="email">Email:</label>
            <input id="email" type="email" bind:value={formUser.email} required />
        </div>
        <div class="form-field">
            <label for="password">Password:</label>
            <input id="password" type="password" bind:value={formUser.password} />
        </div>
        <div class="form-field">
            <button type="submit">{editUser ? 'Update' : 'Create'} User</button>
        </div>
    </div>
</form>

<style>

</style>
