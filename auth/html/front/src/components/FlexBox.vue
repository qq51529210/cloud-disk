<script setup>
import { computed } from 'vue'
const props = defineProps(['justify', 'align'])
const style = computed(() => ({
    'display': 'flex',
    'justify-content': props.justify,
    'align-content': props.align
}))
</script>
<template>
    <div :style="style">
        <slot />
    </div>
</template>
