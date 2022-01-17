<script setup>
import { NButton } from 'naive-ui'

const props = defineProps(['title'])
const emits = defineEmits(['submit'])
</script>
<template>
    <n-button
        size="large"
        type="primary"
        block
        @click="emits('submit')"
    >{{ props.title }}</n-button>
</template>
