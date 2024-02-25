<template>
    <div class="home">
      <div class="top-card-background">
          <canvas ref="canvas"></canvas>
          <div class="left-cards">
              <div class="status-card">
                  <p class="status-text">Status: <span class="status-ok">Operational</span></p>
                  <div class="status-details">
                      <div class="status-table db-count">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-database" width="24" height="24" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 6m-8 0a8 3 0 1 0 16 0a8 3 0 1 0 -16 0" /><path d="M4 6v6a8 3 0 0 0 16 0v-6" /><path d="M4 12v6a8 3 0 0 0 16 0v-6" /></svg>
                          <p>100 DBs</p>
                      </div>
                      <div class="status-table user-count">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-password-user" width="24" height="24" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 17v4" /><path d="M10 20l4 -2" /><path d="M10 18l4 2" /><path d="M5 17v4" /><path d="M3 20l4 -2" /><path d="M3 18l4 2" /><path d="M19 17v4" /><path d="M17 20l4 -2" /><path d="M17 18l4 2" /><path d="M9 6a3 3 0 1 0 6 0a3 3 0 0 0 -6 0" /><path d="M7 14a2 2 0 0 1 2 -2h6a2 2 0 0 1 2 2" /></svg>
                          <p>1000 Users</p>
                      </div>
                      <div class="status-table db-size">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-server-2" width="24" height="24" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M3 4m0 3a3 3 0 0 1 3 -3h12a3 3 0 0 1 3 3v2a3 3 0 0 1 -3 3h-12a3 3 0 0 1 -3 -3z" /><path d="M3 12m0 3a3 3 0 0 1 3 -3h12a3 3 0 0 1 3 3v2a3 3 0 0 1 -3 3h-12a3 3 0 0 1 -3 -3z" /><path d="M7 8l0 .01" /><path d="M7 16l0 .01" /><path d="M11 8h6" /><path d="M11 16h6" /></svg>
                          <p>9999GB</p>
                      </div>
                  </div>
              </div>
          </div>
          <div class="right-cards">
              <p>CPU: 100%</p>
              <p>MEMORY: 9999GB</p>
          </div>
      </div>
      <HelloWorld msg="Welcome to PgHarmony Control Panel"/>
    </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component"
import HelloWorld from "@/components/HelloWorld.vue" // @ is an alias to /src

@Options({
    components: {
        HelloWorld
    }
})
export default class HomeView extends Vue {
    mounted () {
        const canvas = this.$refs.canvas as HTMLCanvasElement | null
        if (canvas && canvas.parentElement) {
            const ctx = canvas.getContext("2d")
            if (ctx) {
                // キャンバスのサイズを設定
                canvas.width = canvas.parentElement.offsetWidth
                canvas.height = canvas.parentElement.offsetHeight

                // ドット背景を描画
                for (let x = 0; x <= canvas.width; x += 10) {
                    for (let y = 0; y <= canvas.height; y += 10) {
                        const opacity = x / canvas.width // 透明度を計算
                        ctx.fillStyle = `rgba(0, 0, 0, ${opacity})` // 透明度を設定
                        ctx.fillRect(x, y, 1, 1) // 1pxの点を描画
                    }
                }
            }
        }
    }
}
</script>

<style>
.top-card-background:after {
    z-index: 1;
}

.top-card-background {
    background: linear-gradient(to right, rgba(255, 255, 255, 0), rgba(255, 255, 255, 0.9)),
                url("https://s3.sda1.net/nyan/contents/f6c11900-10cc-4b37-9e04-91c903219335.jpg");
    background-size: cover;
    height: 300px;
    width: 100%;
    border-radius: 10px;
    object-fit: cover;
    z-index: 0;
    position: relative;
}

.top-card-background canvas {
    position: absolute;
    top: 0;
    left: 0;
    z-index: 1;
    height: 99%;
    width: 99%;
    padding-top: 6px;
}

.left-cards {
    position: absolute;
    top: 0;
    left: 0;
    z-index: 2;
    display: flex;
    justify-content: left;
    align-items: center;
    height: 100%;
    width: 100%;

    > .status-card {
        > .status-details {
            display: flex;
            justify-content: space-around;
            margin-top: 32px;

            > .status-table {
                text-align: center;

                > p {
                    margin: 0;
                }
            }
        }
    }
}

.right-cards {
    position: absolute;
    top: 80px;
    right: 0;
    z-index: 2;
    display: block;
    width: 40%;
    text-align: left;
    max-width: 200px;
}

.status-card {
    background-color: #fdfdfdbd;
    backdrop-filter: blur(20px);
    border-radius: 10px;
    box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.1);
    width: 30%;
    height: 60%;
    margin: 16px;
    text-align: left;

    > .status-text {
        font-size: 1.3em;
        padding-left: 16px;

        > .status-ok {
            color: #00ae4c;
        }
    }
}
</style>
