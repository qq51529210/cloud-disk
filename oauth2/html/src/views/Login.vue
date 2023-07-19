<template>
    <v-app>
        <v-main>
            <v-container fluid
                         class="fill-height">
                <v-row align="center">
                    <v-col>
                        <v-row justify="center">
                            <v-hover v-slot="{ isHovering, props }">
                                <v-card :elevation="isHovering ? 16 : 8"
                                        v-bind="props"
                                        width="400">
                                    <v-card-title class="mb-4 mt-4 text-center">
                                        用户登录
                                    </v-card-title>
                                    <v-card-text>
                                        <v-form ref="form"
                                                v-model="validateResult"
                                                :disabled="requesting"
                                                class="mb-4"
                                                @submit.prevent="submit">
                                            <v-text-field v-model="formData.account"
                                                          class="mb-2"
                                                          variant="outlined"
                                                          clearable
                                                          label="账号"
                                                          :counter="30"
                                                          :rules="formRule.account"></v-text-field>

                                            <v-text-field v-model="formData.password"
                                                          class="mb-2"
                                                          variant="outlined"
                                                          clearable
                                                          label="密码"
                                                          :rules="formRule.password"></v-text-field>
                                            <v-btn block
                                                   variant="flat"
                                                   color="primary"
                                                   size="x-large"
                                                   type="submit"
                                                   :loading="requesting">确定</v-btn>
                                        </v-form>
                                    </v-card-text>
                                </v-card>
                            </v-hover>
                        </v-row>
                    </v-col>
                </v-row>
            </v-container>
        </v-main>
    </v-app>
</template>

<script lang="ts" setup>
import { reactive, ref, inject } from 'vue'
import { post as login, model as loginModel } from '@/api/login'
import { VForm } from 'vuetify/lib/components/index.mjs';
import { redirectURI } from '@/const'
// 
const showTip: any = inject('showTip')
// 数据
const formData: loginModel = reactive({
    account: 'test-user',
    password: '123123',
})
// 验证
const formRule: any = {
    account: [
        (v: string) => !!v || '不能为空'
    ],
    password: [
        (v: string) => !!v || '不能为空'
    ],
}
// 请求中
const requesting = ref(false)
const validateResult = ref(false)
// 提交
const submit = async () => {
    // 验证失败
    if (!validateResult.value) {
        return
    }
    // 请求
    requesting.value = true
    let err = await login(formData)
    requesting.value = false
    if (err) {
        showTip(err.phrase, err.detail)
        return
    }
    // 重定向
    let query = new URLSearchParams(window.location.search);
    let redirect = query.get(redirectURI)
    if (redirect) {
        window.location.replace(redirect)
    }
}

</script>
