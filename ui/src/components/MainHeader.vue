<template>
    <header class= "mb-auto">
        <nav class= "navbar">
            <div class="left-box">
                <span class= "page-title navbar-brand mb-0 h1">{{ title }}</span>
                <span class="beta">Beta</span>
            </div>
            <div class="right-box">
                <svg @click="openModal" xmlns="http://www.w3.org/2000/svg" class="header-action-icon icon-tabler icon-tabler-power" width="24" height="24" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M7 6a7.75 7.75 0 1 0 10 0" /><path d="M12 4l0 8" /></svg>
                <svg @click="showTodo" xmlns="http://www.w3.org/2000/svg" class="header-action-icon icon-tabler icon-tabler-user-circle top-user-icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 12m-9 0a9 9 0 1 0 18 0a9 9 0 1 0 -18 0" /><path d="M12 10m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0" /><path d="M6.168 18.849a4 4 0 0 1 3.832 -2.849h4a4 4 0 0 1 3.834 2.855" /></svg>
            </div>
        </nav>
    </header>
    <transition name="fade-modal">
        <ModalDialog v-if="showModal" :title="modalTitle" :text="modalText" :title-icon="IconType.License" @close="closeModal"/>
    </transition>
</template>

<script>
import ModalDialog from "./ModalDialog.vue"
import { IconType } from "@/types/IconType"

export default {
    name: "TheHeader",
    computed: {
        IconType () {
            return IconType
        }
    },
    components: {
        ModalDialog
    },
    data () {
        return {
            title: "PgHarmony",
            showModal: false,
            modalText: "Sorry. This message shouldn't be here.",
            modalTitle: "Fatal Error!"
        }
    },
    methods: {
        openModal () {
            this.modalText = "©️2024 nexryai All rights reserved.\n\n" +
                "This software is licensed under the MIT License.\n" +
                "This software also uses some third party open source software. Please see GitHub for details."
            this.modalTitle = "Legal Notice"
            this.showModal = true
        },
        showTodo () {
            this.modalText = "This feature is not implemented yet."
            this.modalTitle = "Harmony Board"
            this.showModal = true
        },
        closeModal () {
            console.log("closeModal")
            this.showModal = false
        }
    }
}
</script>

<style scoped>
.navbar {
    padding: 10px 10px 17px 10px;
    text-align: left;
    display: flex;
    justify-content: space-between;
    border-bottom: 1px solid #eaeaea;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    background-color: #ffffff54;
    backdrop-filter: blur(18px);
    z-index: 1000;
}

.left-box {
    margin: 5px 0 0 6px;

    > .page-title {
        font-size: 1.5em;
        color: #4a4a4a;
    }

    > .beta {
        font-size: 0.8em;
        color: #f8f9fa;
        background-color: #007bff;
        padding: 0.25em 0.5em;
        border-radius: 0.25em;
        margin-left: 10px;
    }
}

.right-box {
    margin-right: 32px;

    > .top-user-icon {
        margin-top: 8px;
        margin-left: 19px;
    }

    > .header-action-icon:hover {
        cursor: pointer;
        color: #007bff;
    }
}

</style>
