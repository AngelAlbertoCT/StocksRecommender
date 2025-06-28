<template>
  <div class="bg-white px-6 py-4 rounded-t-xl border-b border-gray-200">
    <h2 class="text-lg font-medium text-gray-900 mb-4">Filtros de búsqueda</h2>
    <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-4">
      <div>
        <label for="ticker" class="block text-sm font-medium text-gray-700 mb-1">Ticker</label>
        <input
          id="ticker"
          v-model="ticker"
          type="text"
          placeholder="Ej: AAPL, MSFT"
          class="w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500"
          @input="applyFilters"
        />
      </div>
      
      <div>
        <label for="company" class="block text-sm font-medium text-gray-700 mb-1">Empresa</label>
        <input
          id="company"
          v-model="company"
          type="text"
          placeholder="Nombre de empresa"
          class="w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500"
          @input="applyFilters"
        />
      </div>
      
      <div>
        <label for="month" class="block text-sm font-medium text-gray-700 mb-1">Mes</label>
        <select
          id="month"
          v-model="month"
          class="w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500"
          @change="applyFilters"
        >
          <option value="">Todos los meses</option>
          <option value="4">Abril</option>
          <option value="5">Mayo</option>
          <option value="6">Junio</option>
        </select>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useStocks } from '../stores/useStocks'

const store = useStocks()

const ticker = ref(store.filters.ticker)
const company = ref(store.filters.company)
const month = ref(store.filters.month)

function applyFilters() {
  store.filters.ticker = ticker.value.trim()
  store.filters.company = company.value.trim()
  store.filters.month = month.value
  store.page = 1
  store.fetchStocks()
}
</script>