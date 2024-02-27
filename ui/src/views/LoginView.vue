<template>
    <div class="login-header">
        <h1>Welcome back</h1>
    </div>

    <div class="login-box">
        <div class="user-icon">
            <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-user-circle" width="24" height="24" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 12m-9 0a9 9 0 1 0 18 0a9 9 0 1 0 -18 0" /><path d="M12 10m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0" /><path d="M6.168 18.849a4 4 0 0 1 3.832 -2.849h4a4 4 0 0 1 3.834 2.855" /></svg>
        </div>
        <div class="login-form">
            <p>Please enter your credentials</p>
            <input type="text" placeholder="Username">
            <input type="password" placeholder="Password">
            <button>Login</button>
        </div>
    </div>
</template>

<style scoped>
.login-header {
    text-align: left;
    background: linear-gradient(to left, rgba(255, 255, 255, 0), rgba(255, 255, 255, 0.9)),
    url("https://s3.sda1.net/nyan/contents/f6c11900-10cc-4b37-9e04-91c903219335.jpg");
    background-size: cover;
    margin: 10%;
    border-radius: 10px;

    > h1 {
        padding: 10px;
    }
}

.login-box {
    margin: 15%;
    text-align: center;

    > .user-icon {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 10px;
    }

    > .login-form {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 10px;

        > p {
            font-size: 1.0em;
            margin-bottom: 10px;
        }

        > input {
            width: 200px;
            padding: 10px;
            margin: 5px;
            border-radius: 5px;
            border: 1px solid #dbdbdb;
        }

        > button {
            width: 100px;
            padding: 10px;
            margin: 5px;
            border-radius: 5px;
            border: 1px solid #dbdbdb;
            background-color: #f0f0f0;
            cursor: pointer;
        }
    }
}
</style>
