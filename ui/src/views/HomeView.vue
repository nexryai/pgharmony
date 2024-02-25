<template>
    <div class="home">
      <div class="top-card-background">
          <div class="left-cards">
              <div class="status-card">
                  <p class="status-text">Status: <span class="status-ok">Operational</span></p>
              </div>
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
export default class HomeView extends Vue {}
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
}

.left-cards {
    display: flex;
    justify-content: left;
    align-items: center;
    height: 100%;
    width: 100%;
    z-index: 1;
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
