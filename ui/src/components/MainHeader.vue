<template>
    <header class= "mb-auto">
        <nav class= "navbar">
            <div class="left-box">
                <span class= "page-title navbar-brand mb-0 h1">{{ title }}</span>
                <span class="beta">Beta</span>
            </div>
            <div class="right-box">
                <svg xmlns="http://www.w3.org/2000/svg" class="header-action-icon icon-tabler icon-tabler-power" width="24" height="24" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M7 6a7.75 7.75 0 1 0 10 0" /><path d="M12 4l0 8" /></svg>
                <svg xmlns="http://www.w3.org/2000/svg" class="header-action-icon icon-tabler icon-tabler-user-circle top-user-icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 12m-9 0a9 9 0 1 0 18 0a9 9 0 1 0 -18 0" /><path d="M12 10m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0" /><path d="M6.168 18.849a4 4 0 0 1 3.832 -2.849h4a4 4 0 0 1 3.834 2.855" /></svg>
            </div>
        </nav>
    </header>
</template>

<script>
export default {
    name: "TheHeader",
    data () {
        return {
            title: "PgHarmony"
        }
    }
}
</script>

<style>
.navbar {
    padding: 10px 10px 17px 10px;
    text-align: left;
    display: flex;
    justify-content: space-between;
    border-bottom: 1px solid #eaeaea;
    position: fixed;
    top: 0;
    width: 100%;
    background-color: #ffffff54;
    backdrop-filter: blur(18px);
    z-index: 1000;
}

.page-title {
    font-size: 1.5em;
    color: #4a4a4a;
}

.beta {
    font-size: 0.8em;
    color: #f8f9fa;
    background-color: #007bff;
    padding: 0.25em 0.5em;
    border-radius: 0.25em;
    margin-left: 10px;
}

.right-box {
    margin-right: 32px;
}

.top-user-icon {
    margin-top: 8px;
    margin-left: 19px;
}

.header-action-icon:hover {
    cursor: pointer;
    color: #007bff;
}
</style>
