<script setup>
import { NForm, NButton } from 'naive-ui'

const props = defineProps(['button'])
const emits = defineEmits(['submit'])

</script>
<template>
    <n-form size="large" :show-label="false">
        <slot />
    </n-form>
    <n-button
        size="large"
        type="primary"
        block
        @click="emits('submit')"
    >{{ props.button }}</n-button>
</template>
