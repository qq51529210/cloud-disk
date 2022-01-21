<script setup>
import { ref, reactive } from 'vue'
import LayoutVue from '../components/Layout.vue';
import VerifyCodeVue from '../components/VerifyCode.vue'
import LinkVue from '../components/Link.vue';
import LinksVue from '../components/Links.vue';
import FormVue from '../components/Form.vue';
import PasswordVue from '../components/Password.vue';
import * as service from '../service'
import { useRouter } from 'vue-router'

const router = useRouter()

const loading = ref(false)
const errorText = ref('')

const model = reactive({
    number: null,
    code: null,
    password: ''
})

const onSubmit = async () => {
    loading.value = true
    let res = await service.signUp(model)
    loading.value = false
    if (res.error) {
        errorText.value = res.error
        return
    }
    router.push('/sign_in' + location.search)
}

</script>

<template>
    <LayoutVue title="注册" :loading="loading" :error="errorText">
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
