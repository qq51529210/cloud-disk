<script setup>
import { ref, computed } from 'vue'
import { NFormItem, NInputNumber, NButton, NInputGroup } from 'naive-ui'

const props = defineProps(['number', 'code'])
const emits = defineEmits(['update:number', 'update:code'])

const seconds = ref(0)

const title = computed(() => {
    if (seconds.value) {
        return `${seconds.value}秒`
    }
    return `验证码`
})

const onSubmit = () => {
    if (seconds.value) {
        return
    }
    seconds.value = 10
    let timer = setInterval(() => {
        seconds.value--
        if (!seconds.value) {
            clearInterval(timer)
        }
    }, 1000);
}

</script>

<template>
    <n-form-item>
        <n-input-number
            :value="props.number"
            :on-update:value="v => emits('update:number', v)"
            placeholder="手机号码"
            clearable
            :show-button="false"
            style="width:100%;"
        />
    </n-form-item>
    <n-form-item>
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
