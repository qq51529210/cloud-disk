module.exports = {
    productionSourceMap: false,
    pages: {
        sign_in: 'src/sign_in/main.js',
        sign_up: 'src/sign_up/main.js'
    },
    chainWebpack: config => {
        config.resolve.alias.set('vue-i18n', 'vue-i18n/dist/vue-i18n.cjs.js')
    }
}