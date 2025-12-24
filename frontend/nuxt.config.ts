// https://nuxt.com/docs/api/configuration/nuxt-config

export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  nitro: {
    watchOptions: {
      ignored: ['**/sql_dump/**']
    }
  },

  runtimeConfig: {
    public: {
      apiUserBaseUrl: process.env.NUXT_USER_API_BASE_URL,
      apiOrderBaseUrl: process.env.NUXT_ORDER_API_BASE_URL,
      apiProductBaseUrl: process.env.NUXT_PRODUCT_API_BASE_URL,
      apiPaymentBaseUrl: process.env.NUXT_PAYMENT_API_BASE_URL,
      apiNotificationBaseUrl: process.env.NUXT_NOTIFICATION_API_BASE_URL,
    }
  },

  modules: [
    '@nuxt/content',
    '@nuxt/eslint',
    '@nuxt/fonts',
    '@nuxt/icon',
    '@nuxt/image',
    '@nuxt/ui',
    '@pinia/nuxt',
  ],
  vite: {
    plugins: [
    ],
  },
  css: [
    '@tabler/icons-webfont/dist/tabler-icons.min.css',
    'swiper/css',
    'swiper/css/navigation',
    'swiper/css/pagination',
    'swiper/css/thumbs',
    'dropzone/dist/dropzone.css',
    'quill/dist/quill.snow.css',
    'nouislider/dist/nouislider.min.css',
    '@/assets/css/theme.min.css',
  ],
  build: {
    transpile: ['swiper']
  },
  app: {
    head: {
      title: 'Sayur Project',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/images/favicon/favicon.ico' },
        { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
        { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' },
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap' },
        { rel: 'stylesheet', href:'/libs/tiny-slider/dist/tiny-slider.css'},
        { rel: 'stylesheet', href:'/libs/swiper/swiper-bundle.min.css'},
        { rel: 'stylesheet', href:'/libs/simplebar/dist/simplebar.min.css'},
      ],
      script: [
        {
          src: "https://app.sandbox.midtrans.com/snap/snap.js", 'data-client-key': process.env.MIDTRANS_CLIENT_KEY,
        }
      ]
    },
  }
})