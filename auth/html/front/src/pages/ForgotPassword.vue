<script setup>
import { ref, reactive } from 'vue'
import LayoutVue from '../components/Layout.vue';
import VerifyCodeVue from '../components/VerifyCode.vue'
import LinkVue from '../components/Link.vue';
import LinksVue from '../components/Links.vue';
import FormVue from '../components/Form.vue';
import PasswordVue from '../components/Password.vue';

const loading = ref(false)
const errorText = ref('')

const model = reactive({
    number: null,
    code: null,
    password: ''
})

const onSubmit = async () => {
    console.log(location.pathname)
    console.log(location.search)
}

</script>

<template>
    <LayoutVue title="找回密码" :loading="loading" :error="errorText">
        <FormVue button="修改" @submit="onSubmit">
            <VerifyCodeVue
                v-model:number="model.number"
                v-model:code="model.code"
            />
            <PasswordVue v-model="model.password" />
        </FormVue>
        <LinksVue justify="flex-end">
            <LinkVue title="返回登录" to="/sign_in" />
        </LinksVue>
    </LayoutVue>
</template>
