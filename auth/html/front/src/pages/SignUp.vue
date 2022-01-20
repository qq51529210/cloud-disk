<script setup>
import { reactive } from 'vue'
import LayoutVue from '../components/Layout.vue';
import VerifyCodeVue from '../components/VerifyCode.vue'
import LinkVue from '../components/Link.vue';
import LinksVue from '../components/Links.vue';
import FormVue from '../components/Form.vue';
import PasswordVue from '../components/Password.vue';


const model = reactive({
    number: null,
    code: null,
    password: ''
})

const onSubmit = async () => {
    console.log(number.value, code.value, password.value)
}


</script>

<template>
    <LayoutVue title="注册">
        <FormVue button="注册" @submit="onSubmit">
            <VerifyCodeVue
                v-model:number="model.number"
                v-model:code="model.code"
            />
            <PasswordVue v-model="model.password" />
        </FormVue>
        <LinksVue justify="flex-end">
            <LinkVue title="登录已有账号" to="/sign_in" />
        </LinksVue>
    </LayoutVue>
</template>
