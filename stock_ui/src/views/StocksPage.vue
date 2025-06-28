<template>
  <section class="bg-gray-50 min-h-screen py-8 px-4 sm:px-6 lg:px-8">
    <div class="max-w-7xl mx-auto">
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">Análisis de Acciones</h1>
        <p class="text-gray-600">Recomendaciones actualizadas del mercado</p>
      </div>

      <div class="bg-white rounded-xl shadow-md overflow-hidden mb-8">
        <SearchFilters />
      </div>

      <div v-if="store.loading" class="flex justify-center items-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
      </div>

      <div v-else>
        <div v-if="store.stocks.length > 0" class="bg-white rounded-xl shadow-md overflow-hidden">
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Ticker</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Empresa</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Broker</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Acción</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Rating</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Objetivo</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr 
                  v-for="stock in store.stocks" 
                  :key="stock.ticker + stock.time"
                  class="hover:bg-gray-50 transition-colors duration-150"
                >
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="font-medium text-blue-600">{{ stock.ticker }}</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-gray-900">{{ stock.company }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-gray-500">{{ stock.brokerage }}</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span 
                      :class="{
                        'bg-green-100 text-green-800': stock.action === 'Comprar',
                        'bg-red-100 text-red-800': stock.action === 'Vender',
                        'bg-yellow-100 text-yellow-800': stock.action === 'Mantener'
                      }"
                      class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                    >
                      {{ stock.action }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-gray-500">
                    <span class="text-gray-400">{{ stock.rating_from }}</span> → 
                    <span class="font-medium">{{ stock.rating_to }}</span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-gray-500">
                    <span class="text-gray-400">{{ stock.target_from }}</span> → 
                    <span class="font-medium">{{ stock.target_to }}</span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="bg-gray-50 px-6 py-3 border-t border-gray-200">
            <Pagination />
          </div>
        </div>

        <div v-else class="text-center py-12">
          <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-gray-100 mb-4">
            <svg class="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <h3 class="text-lg font-medium text-gray-900 mb-1">No se encontraron acciones</h3>
          <p class="text-gray-500">Intenta ajustar los filtros de búsqueda</p>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { onMounted, watch } from 'vue'
import { useStocks } from '../stores/useStocks'
import SearchFilters from '../components/SearchFilters.vue'
import Pagination from '../components/Pagination.vue'

const store = useStocks()

onMounted(() => {
  store.fetchStocks()
})

watch(
  () => store.filters,
  () => {
    store.page = 1
    store.fetchStocks()
  },
  { deep: true }
)

watch(
  () => store.page,
  () => {
    store.fetchStocks()
  }
)
</script>