<script setup>
import { NAlert } from 'naive-ui'

const props = defineProps(['text'])
</script>
<template>
    <n-alert
        v-if="props.text"
        :title="props.text"
        type="error"
        closable
        style="margin-bottom: 24px;"
    />
</template>
