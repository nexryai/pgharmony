<template>
    <div v-if="!icon">
        <p class="debug-message">Error: ERR_HbRenderIcon_INVALID_ICONTYPE</p>
    </div>
    <svg v-else-if="icon === IconType.Copyright" xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-copyright" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 12m-9 0a9 9 0 1 0 18 0a9 9 0 1 0 -18 0" /><path d="M14 9.75a3.016 3.016 0 0 0 -4.163 .173a2.993 2.993 0 0 0 0 4.154a3.016 3.016 0 0 0 4.163 .173" /></svg>
    <svg v-else-if="icon === IconType.Save" xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-device-floppy" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M6 4h10l4 4v10a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12a2 2 0 0 1 2 -2" /><path d="M12 14m-2 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0" /><path d="M14 4l0 4l-6 0l0 -4" /></svg>
    <svg v-else-if="icon === IconType.X" xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-x" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M18 6l-12 12" /><path d="M6 6l12 12" /></svg>
</template>

<script lang="ts">
import { IconType } from "@/types/IconType"

export default {
    props: {
        icon: {
            type: IconType,
            required: true
        }
    },
    data () {
        return {
            IconType
        }
    }
}
</script>

<style scoped>
.debug-message {
    display: none;
}
</style>
