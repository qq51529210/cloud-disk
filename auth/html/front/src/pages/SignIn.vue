<script setup>
import { ref, reactive } from 'vue'
import { NTabs, NTabPane } from 'naive-ui'
import LayoutVue from '../components/Layout.vue';
import VerifyCodeVue from '../components/VerifyCode.vue'
import LinkVue from '../components/Link.vue';
import LinksVue from '../components/Links.vue';
import FormVue from '../components/Form.vue';
import TextVue from '../components/Text.vue';
import PasswordVue from '../components/Password.vue';
import QRCodeVue from '../components/QRCode.vue';
import * as service from '../service'
import { parseUrlQuery } from '../util/parse-url-query'

const loading = ref(false)
const errorText = ref('')

const accountModel = reactive({
    account: '',
    password: '',
})

const phoneModel = reactive({
    number: null,
    code: null,
})

const onAccountSubmit = async () => {
    submit(service.signInAccount(accountModel))
}

const onPhoneSubmit = async () => {
    submit(service.signInPhone(phoneModel))
}

const submit = async (sumbit) => {
    loading.value = true
    errorText.value = ''
    let res = await sumbit
    loading.value = false
    if (res.error) {
        errorText.value = res.error
        // return
    }
    // oauth2
    if (location.pathname === '/oauth2') {
        let queries = parseUrlQuery(location.search)
        console.log(queries)
    }
}

</script>

<template>
    <LayoutVue title="登录" :loading="loading" :error="errorText">
        <n-tabs
            default-value="account"
            justify-content="space-evenly"
            size="large"
        >
            <n-tab-pane name="account" tab="账号">
                <FormVue button="登录" @submit="onAccountSubmit">
                    <TextVue v-model="accountModel.account" placeholder="账号" />
                    <PasswordVue v-model="accountModel.password" />
                </FormVue>
            </n-tab-pane>
            <n-tab-pane name="sms" tab="验证码">
                s
                <FormVue button="登录" @submit="onPhoneSubmit">
                    <VerifyCodeVue
                        v-model:number="phoneModel.number"
                        v-model:code="phoneModel.code"
                    />
                </FormVue>
            </n-tab-pane>
            <n-tab-pane name="qrcode" tab="二维码">
                <QRCodeVue />
            </n-tab-pane>
        </n-tabs>
        <LinksVue justify="space-between">
            <LinkVue title="注册新账号" to="/sign_up" />
            <LinkVue title="找回密码" to="/forgot_password" />
        </LinksVue>
    </LayoutVue>
</template>
