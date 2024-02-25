<template>
    <transition name="fade">
        <div v-if="showModal" class="modal">
            <div class="modal-title-box">
                <div class="title-background">
                    <canvas ref="canvas"></canvas>
                </div>
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
    },
    mounted () {
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
        gradient.addColorStop(0, "rgba(0, 0, 0, 0.2)")
        gradient.addColorStop(1, "rgba(0, 0, 0, 0)")

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
                ctx.fillRect(x, y, 1, 1) // 1pxの点を描画
            }
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
    display: flex;
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

.title-background {
    position: absolute;
    top: 0;
    left: 0;
    z-index: 1;
    height: 400px;
    width: 300px;
}

.modal-title-text {
    position: absolute;
    margin: 20px;
    z-index: 2;
}

.fade-enter-active, .fade-leave-active {
    transition: opacity .5s;
}
.fade-enter, .fade-leave-to {
    opacity: 0;
}
</style>
