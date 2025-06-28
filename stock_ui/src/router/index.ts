import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router' // Importación específica para tipos
import StocksPage from '@/views/StocksPage.vue'
import RecommendationsPage from '@/views/RecommendationsPage.vue'

const routes: RouteRecordRaw[] = [ // También puedes usar esta sintaxis más corta
  {
    path: '/stocks',
    name: 'Stocks',
    component: StocksPage
  },
  {
    path: '/recommendations',
    name: 'Recommendations',
    component: RecommendationsPage,
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router