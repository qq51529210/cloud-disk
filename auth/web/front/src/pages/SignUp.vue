<script setup>
import LayoutVue from '../components/Layout.vue';
import { NCard, NTabs, NTabPane, NForm, NFormItem, NInput, NInputNumber, NButton, NInputGroup } from 'naive-ui'
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import QRCodeVue3 from 'qrcode-vue3'

const router = useRouter()
const sendingSMS = ref(0)

const sendSMSBtnTitle = computed(() => {
    if (sendingSMS.value) {
        return `${sendingSMS.value}秒`
    }
    return `验证码`
})

const onSendSMS = () => {
    if (sendingSMS.value) {
        return
    }
    sendingSMS.value = 10
    let timer = setInterval(() => {
        sendingSMS.value--
        if (!sendingSMS.value) {
            clearInterval(timer)
        }
    }, 1000);
}

</script>

<template>
    <LayoutVue title="注册">
        <n-form size="large" :show-label="false">
            <n-form-item>
                <n-input-number
                    placeholder="手机号码就是账号"
                    clearable
                    :show-button="false"
                    style="width:100%;"
                />
            </n-form-item>
            <n-form-item>
                <n-input
                    type="password"
                    show-password-on="click"
                    clearable
                    placeholder="密码"
                />
            </n-form-item>
            <n-form-item>
                <n-input-group>
                    <n-input-number
                        clearable
                        placeholder="验证码"
                        :show-button="false"
                        style="width:100%;"
                    />
                    <n-button
                        :loading="sendingSMS"
                        @click="onSendSMS"
                    >{{ sendSMSBtnTitle }}</n-button>
                </n-input-group>
            </n-form-item>
        </n-form>
        <n-button size="large" type="primary" block>注册</n-button>
        <div class="link">
            <n-button
                text
                tag="a"
                type="primary"
                @click="router.push('/sign_in')"
            >登录已有账号</n-button>
        </div>
    </LayoutVue>
</template>

<style scoped>
.link {
    margin-top: 14px;
    display: flex;
    justify-content: flex-end;
    align-content: center;
}
</style>
