<template>
    <var-input :placeholder="placeholder" v-model="value">
        <template #prepend-icon>
            <var-icon name="close-circle-outline" />
        </template>
        <template #append-icon>
            <var-icon name="minus" />
        </template>
    </var-input>
</template>

<script>
export default {
    props: {
        modelValue: String,
        placeholder: String,
    },
    emits: ['update:modelValue']
    setup() {
    }
}
</script>