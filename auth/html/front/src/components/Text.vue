<script setup>
import { NFormItem, NInput } from 'naive-ui'

const props = defineProps(['modelValue', 'placeholder', 'feedback'])
const emits = defineEmits(['update:modelValue'])

</script>

<template>
    <n-form-item
        :feedback="props.feedback"
        :validation-status="props.feedback ? 'error' : undefined"
    >
        <n-input
            :value="props.modelValue"
            :on-update:value="v => emits('update:modelValue', v)"
            clearable
            :placeholder="props.placeholder"
        />
    </n-form-item>
</template>
