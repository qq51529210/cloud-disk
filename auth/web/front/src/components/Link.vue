<script setup>
import { NButton } from 'naive-ui'
import { useRouter } from 'vue-router'

const props = defineProps(['title', 'to'])
const router = useRouter()

const onSubmit = () => {
    router.push(props.to)
}

</script>

<template>
    <n-button text tag="a" type="primary" @click="onSubmit">{{ props.title }}</n-button>
</template>
