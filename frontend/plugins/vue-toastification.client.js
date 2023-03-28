import Toast from 'vue-toastification/dist/index.mjs'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(Toast, {
    transition: "Vue-Toastification__fade",
    maxToasts: 20,
    newestOnTop: true,
    position: "top-right",
    timeout: 3000,
  })
})
