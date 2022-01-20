<script setup>
import { NDivider, NCard } from 'naive-ui'
const props = defineProps(['title']);

</script>

<template>
    <div style="height: 20px;"></div>
    <n-divider>
        <h2>{{ props.title }}</h2>
    </n-divider>
    <div style="display: flex; justify-content: center;">
        <n-card style="width: 330px;" :bordered="true">
            <slot />
        </n-card>
    </div>
</template>
