<script setup>
import { NFormItem, NInput } from 'naive-ui'

const props = defineProps(['modelValue', 'feedback'])
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
            type="password"
            show-password-on="click"
            clearable
            placeholder="密码"
        />
    </n-form-item>
</template>
