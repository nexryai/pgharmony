<template>
    <transition name="fade">
        <div v-if="showModal" class="modal">
            <div class="modal-title-box">
                <h2 class="modal-title-text">Modal Dialog</h2>
            </div>
            <div class="modal-content">
                <span class="close" @click="closeModal">&times;</span>
                <p>{{ text }}</p> <!-- propsで受け取ったテキストを表示します -->
            </div>
        </div>
    </transition>
</template>

<script>
export default {
    props: ["text"], // 親コンポーネントからテキストを受け取るためのpropsを定義します
    data () {
        return {
            showModal: true
        }
    },
    methods: {
        closeModal () {
            this.$emit("close")
        }
    }
}
</script>

<style>
.modal {
    display: block;
    position: fixed;
    z-index: 9999;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto;
    background-color: rgba(225, 225, 225, 0.4);
    backdrop-filter: blur(18px);
}

.modal-title-box {
    text-align: left;
    margin: 15px;
}

.modal-title-text {
    font-size: 14px;
}

.modal-content {
    margin: 15% auto;
    padding: 20px;
    width: 80%;
}

.close {
    color: #aaa;
    float: right;
    font-size: 28px;
    font-weight: bold;
}

.close:hover,
.close:focus {
    color: black;
    text-decoration: none;
    cursor: pointer;
}

.fade-enter-active, .fade-leave-active {
    transition: opacity .5s;
}
.fade-enter, .fade-leave-to {
    opacity: 0;
}
</style>
