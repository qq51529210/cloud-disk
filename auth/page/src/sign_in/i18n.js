import { createI18n } from 'vue-i18n'

const i18n = createI18n({
    locale: navigator.language,
    globalInjection: true,
    fallbackLocale: 'en-US',
    messages: {
        'zh-CN': {
            htmlTitle: '登录',
            tabAccount: '账号登录',
            tabSMS: '短信登录',
            account: '账号',
            password: '密码',
            showPassword: '显示密码',
            forgetPassword: '忘记密码',
            newAccount: '注册新账号',
            phone: '手机号',
            smsCode: '短信验证码',
            getSMSCode: '获取验证码',
            signIn: '登录',
            otherAccount: '其他账号登录',
        },
        'en-US': {
            htmlTitle: 'Sign In',
            tabAccount: 'Account Sign In',
            tabSMS: 'SMS Sign In',
            account: 'Account',
            password: 'Password',
            showPassword: 'Show Password',
            phone: 'Cellular Phone',
            smsCode: 'SMS Verification Code',
            getSMSCode: 'Verification Code',
            forgetPassword: 'Forget Password?',
            newAccount: 'New Account',
            signIn: 'Sign In',
            otherAccount: 'Other Account Sign In',
        }
    }
})

export default i18n
