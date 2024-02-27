<template>
    <div class="title-bar">
        <p class="section-title">{{ title }}</p>
        <hr class="section-title-underline">
    </div>
</template>

<script>
export default {
    props: ["title"]
}
</script>

<style scoped>
.title-bar {
    text-align: left;
    margin: 50px 28px 28px 28px;

    > .section-title {
        font-size: 1.1em;
        color: #6a6a6a;
    }

    > .section-title-underline {
        width: 250px;
        margin-left: 1px;
        margin-top: -10px;
        border: 1px solid #dbdbdb;
    }
}
</style>
