<template>
    <div v-if="showModal" class="modal">
        <div class="modal-title-box">
            <div class="title-background">
                <canvas ref="canvas"></canvas>
            </div>
            <h2 v-if="isRedDialog" class="modal-title-text modal-title-red-text">
                <HbRenderIcon :icon="titleIcon" class="modal-title-icon modal-title-red-icon"/>
                {{ title }}
            </h2>
            <h2 v-else class="modal-title-text">
                <HbRenderIcon :icon="titleIcon" class="modal-title-icon"/>
                {{ title }}
            </h2>
        </div>
        <span class="close" @click="closeModal">
            <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-x" width="24" height="24" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M18 6l-12 12" /><path d="M6 6l12 12" /></svg>
        </span>
        <div class="modal-content">
            <p>{{ text }}</p>
        </div>
    </div>
</template>

<script>
import HbRenderIcon from "@/components/core/HbRenderIcon.vue"
import { IconType } from "@/types/IconType"

export default {
    components: {
        HbRenderIcon
    },
    props: {
        title: String,
        titleIcon: IconType,
        text: String,
        isRedDialog: Boolean
    },
    data () {
        return {
            showModal: true
        }
    },
    methods: {
        closeModal () {
            this.$emit("close")
        }
    },
    mounted () {
        // modalTitleの背景を描画
        const canvas = this.$refs.canvas
        const ctx = canvas.getContext("2d")

        // キャンバスのサイズを設定
        canvas.width = canvas.parentElement.offsetWidth
        canvas.height = canvas.parentElement.offsetHeight

        // グラデーションを作成
        const gradient = ctx.createRadialGradient(
            0, 0, 0, // 内側の円の中心と半径
            0, 0, Math.max(canvas.width - 10000, canvas.height - 200) // 外側の円の中心と半径
        )

        if (this.isRedDialog) {
            gradient.addColorStop(0, "rgba(255, 0, 0, 0.45)")
            gradient.addColorStop(1, "rgba(255, 0, 0, 0)")
        } else {
            gradient.addColorStop(0, "rgba(0, 0, 0, 0.2)")
            gradient.addColorStop(1, "rgba(0, 0, 0, 0)")
        }

        // グラデーションを描画
        ctx.fillStyle = gradient
        ctx.fillRect(0, 0, canvas.width, canvas.height)

        // ドット背景を描画
        for (let x = 0; x <= 200; x += 10) {
            for (let y = 0; y <= 200; y += 10) {
                // ドットの位置に応じたグラデーションの透明度を計算
                const distance = Math.sqrt(x * x + y * y)
                const maxDistance = Math.sqrt(200 * 200 + 200 * 200)
                const opacity = 0.2 * (1 - distance / maxDistance)
                ctx.fillStyle = `rgba(0, 0, 0, ${opacity})` // 透明度を設定

                ctx.fillRect(x, y, 1.2, 1.2) // ドットを描画
            }
        }
    }
}
</script>

<style scoped>
.fade-modal-enter-active, .fade-modal-leave-active {
    > .modal {
        transition: opacity 0.2s, background-color 0.2s !important;
    }

    > .modal-content {
        transition: opacity 0.2s, transform 0.2s !important;
    }
}
.fade-modal-enter-from, .fade-modal-leave-to {
    > .modal {
        opacity: 0;
        background-color: rgba(225, 225, 225, 0);
    }

    > .modal-content {
        pointer-events: none;
        opacity: 0;
        transform: scale(0.9);
    }
}

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
    display: flex;

    > .modal-title-text {
        position: absolute;
        font-size: 1.25em;
        margin: 20px;
        z-index: 2;

        > .modal-title-icon {
            width: auto;
            height: 28px;
            vertical-align: top;
        }
    }

    > .modal-title-red-text {
        color: red;

        > .modal-title-red-icon {
            fill: red;
        }
    }
}

.modal-content {
    margin: 15% auto;
    padding: 20px;
    width: 80%;
}

.close {
    position: absolute;
    right: 32px;
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

.title-background {
    position: absolute;
    top: 0;
    left: 0;
    z-index: 1;
    height: 400px;
    width: 300px;
}
</style>
