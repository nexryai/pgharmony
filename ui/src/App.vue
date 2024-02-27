<template>
    <MainHeader />
    <router-view class="app-main-view"/>
    <div class="beta-warning">
        <span>Dev build. Not indicative of version will be released.</span>
    </div>
    <div class="footer">
        <span class="footer-copyright">©2024 nexryai All rights reserved.</span>
    </div>
</template>

<script lang="ts">
import MainHeader from "@/components/MainHeader.vue"

export default {
    name: "App",
    components: {
        MainHeader
    }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}

.app-main-view {
  padding-top: 80px;
}

.beta-warning {
    position: fixed;
    padding: 0 10px;
    bottom: 0;
    left: 0;
    background-color: #0000004f;
    color: #000;
    font-size: .8em;
    z-index: 10000;
}

.footer {
    margin-top: 20px;
    text-align: right;
    width: 100%;
}

.footer-copyright {
    font-size: .8em;
    color: #000;
    padding: 10px;
}

.fade-modal-enter-active, .fade-modal-leave-active {
    transition: opacity 0.2s, background-color 0.2s !important;
}
.fade-modal-enter-from, .fade-modal-leave-to {
    opacity: 0;
    background-color: rgba(225, 225, 225, 0);
}
</style>
