<script setup>
import { NDivider, NSpace, NSpin, NCard } from 'naive-ui'
import ErrorVue from '../components/Error.vue';
const props = defineProps(['title', 'loading', 'error']);

</script>

<template>
    <div style="height: 20px;"></div>
    <n-divider>
        <h2>{{ props.title }}</h2>
    </n-divider>
    <div style="display: flex; justify-content: center;">
        <n-space vertical>
            <ErrorVue :text="props.error" />
            <n-spin :show="props.loading">
                <n-card style="width: 330px;" :bordered="true">
                    <slot />
                </n-card>
            </n-spin>
        </n-space>
    </div>
</template>
