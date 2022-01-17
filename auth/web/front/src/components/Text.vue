<script setup>
import { NFormItem, NInput } from 'naive-ui'

const props = defineProps(['modelValue', 'placeholder'])
const emits = defineEmits(['update:modelValue'])

</script>

<template>
    <n-form-item>
        <n-input
            :value="props.modelValue"
            :on-update:value="v => emits('update:modelValue', v)"
            clearable
            :placeholder="props.placeholder"
        />
    </n-form-item>
</template>
