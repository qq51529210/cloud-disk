<script setup>
import LayoutVue from '../components/Layout.vue';
import { NCard, NTabs, NTabPane, NForm, NFormItem, NInput, NInputNumber, NButton, NInputGroup } from 'naive-ui'
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import QRCodeVue3 from 'qrcode-vue3'

const router = useRouter()
const sendingSMS = ref(0)
const qrCodeSize = 320

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
    <LayoutVue title="登录">
        <n-tabs
            default-value="account"
            justify-content="space-evenly"
            size="large"
        >
            <n-tab-pane name="account" tab="账号">
                <n-form size="large" :show-label="false">
                    <n-form-item>
                        <n-input placeholder="账号" clearable />
                    </n-form-item>
                    <n-form-item>
                        <n-input
                            type="password"
                            show-password-on="click"
                            clearable
                            placeholder="密码"
                        />
                    </n-form-item>
                </n-form>
                <n-button size="large" type="primary" block>登录</n-button>
            </n-tab-pane>
            <n-tab-pane name="sms" tab="验证码">
                <n-form size="large" :show-label="false">
                    <n-form-item>
                        <n-input-number
                            placeholder="手机号码"
                            clearable
                            :show-button="false"
                            style="width:100%;"
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
                <n-button size="large" type="primary" block>登录</n-button>
            </n-tab-pane>
            <n-tab-pane name="qrcode" tab="二维码">
                <div class="qrcode">
                    <QRCodeVue3
                        value="test qrcode"
                        :width="qrCodeSize"
                        :height="qrCodeSize"
                        :backgroundOptions="{ color: '#ffffff' }"
                        :qrOptions="{ typeNumber: 0, mode: 'Byte', errorCorrectionLevel: 'H' }"
                        :imageOptions="{ hideBackgroundDots: true, imageSize: 0.4, margin: 0 }"
                        :dotsOptions="{
                            type: 'extra-rounded',
                            color: '#26249a',
                            gradient: {
                                type: 'linear',
                                rotation: 0,
                                colorStops: [
                                    { offset: 0, color: '#26249a' },
                                    { offset: 1, color: '#26249a' },
                                ],
                            },
                        }"
                        :cornersSquareOptions="{ type: 'square', color: '#000000' }"
                        :cornersDotOptions="{ type: undefined, color: '#000000' }"
                    />
                </div>
            </n-tab-pane>
        </n-tabs>
        <div class="link">
            <n-button
                text
                tag="a"
                type="primary"
                @click="router.push('/sign_up')"
            >注册新账号</n-button>
            <n-button
                text
                tag="a"
                type="primary"
                @click="router.push('/forgot_password')"
            >找回密码</n-button>
        </div>
    </LayoutVue>
</template>

<style scoped>
.link {
    margin-top: 14px;
    display: flex;
    justify-content: space-between;
    align-content: center;
}
.qrcode {
    display: flex;
    justify-content: center;
    align-content: center;
}
</style>
