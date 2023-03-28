// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    app: {
        head: {
            title: 'Scretch Image',
            link: [
                {rel: 'apple-touch-icon', sizes: '180x180', href: '/apple-touch-icon.png'},
                {rel: 'icon', type: 'image/png', sizes: '32x32', href: '/favicon-32x32.png'},
                {rel: 'icon', type: 'image/png', sizes: '16x16', href: '/favicon-16x16.png'},
                {rel: 'manifest', href: '/site.webmanifest'},
            ]
        }
    },
    modules: [
        '@nuxtjs/i18n',
    ],
    i18n: {
        locales: [
            {
                code: 'en',
                name: 'English',
                file: 'en.js',
            },
            {
                code: 'ru',
                name: 'Русский',
                file: 'ru.js',
            },
            {
                code: 'es',
                name: 'Español',
                file: 'es.js',
            },
            {
                code: 'fr',
                name: 'Français',
                file: 'fr.js',
            },
            {
                code: 'ka',
                name: 'ქართული',
                file: 'ka.js',
            },
        ],
        vueI18n: {
            fallbackLocale: 'en',
        },
        defaultLocale: 'en',
        lazy: true,
        langDir: 'locales/',
        detectBrowserLanguage: {
            useCookie: true,
            cookieKey: 'i18n_redirected',
            redirectOn: 'root',
        }
    },
    plugins: [
        '~/plugins/vue-toastification.client.js',
    ],
    build : {
        transpile: ['vue-toastification']
    },
    css: [
        '~/assets/css/main.css',
        'vue-toastification/dist/index.css',
    ],
    postcss: {
        plugins: {
            tailwindcss: {},
            autoprefixer: {},
        },
    },
})
