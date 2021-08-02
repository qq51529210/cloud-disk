import { createI18n } from 'vue-i18n'

const i18n = createI18n({
    locale: navigator.language,
    globalInjection: true,
    fallbackLocale: 'en-US',
    messages: {
        'zh-CN': {
            htmlTitle: '注册',
            tabPhone: '手机注册',
            tabEmail: '邮箱注册',
        },
        'en-US': {
            htmlTitle: 'Sign Up',
            tabPhone: 'Cellular Phone Sign Up',
            tabEmail: 'Email Sign In',
        }
    }
})

export default i18n
