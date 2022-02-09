<script setup>
import { ref, computed } from 'vue'
import { NFormItem, NInputNumber, NButton, NInputGroup } from 'naive-ui'
import { getPhoneCode } from '../service';

const props = defineProps(['number', 'code'])
const emits = defineEmits(['update:number', 'update:code'])

const seconds = ref(0)

const numberFeedback = ref('')
const codeFeedback = ref('')

const title = computed(() => {
    if (seconds.value) {
        return `${seconds.value}秒`
    }
    return `验证码`
})

const onSubmit = () => {
    if (seconds.value || !validateNumber()) {
        return
    }
    seconds.value = 60
    let timer = setInterval(() => {
        seconds.value--
        if (!seconds.value) {
            clearInterval(timer)
        }
    }, 1000);
    //
    getPhoneCode({
        number: props.number + ''
    })
}

const validateNumber = () => {
    if (!props.number) {
        numberFeedback.value = '不能为空'
        return false
    }
    if (!/^\d{11}$/.test(props.number)) {
        numberFeedback.value = '格式不正确'
        return false
    }
    numberFeedback.value = ''
    return true
}

const validateCode = () => {
    if (!props.code) {
        codeFeedback.value = '不能为空'
        return false
    }
    if (!/^\d{6}$/.test(props.code)) {
        codeFeedback.value = '格式不正确'
        return false
    }
    codeFeedback.value = ''
    return true
}

const validate = () => {
    return validateNumber() && validateCode()
}

defineExpose({ validate })

</script>

<template>
    <n-form-item
        :feedback="numberFeedback"
        :validation-status="numberFeedback ? 'error' : undefined"
    >
        <n-input-number
            :value="props.number"
            :on-update:value="v => emits('update:number', v)"
            placeholder="手机号码"
            clearable
            :show-button="false"
            style="width:100%;"
        />
    </n-form-item>
    <n-form-item
        :feedback="codeFeedback"
        :validation-status="codeFeedback ? 'error' : undefined"
    >
        <n-input-group>
            <n-input-number
                :value="props.code"
                :on-update:value="v => emits('update:code', v)"
                clearable
                placeholder="验证码"
                :show-button="false"
                style="width:100%;"
            />
            <n-button :loading="seconds !== 0" @click="onSubmit">{{ title }}</n-button>
        </n-input-group>
    </n-form-item>
</template>
