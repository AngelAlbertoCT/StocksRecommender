<template>
  <div class="flex items-center justify-between px-4 py-3 bg-white border-t border-gray-200 sm:px-6 rounded-b-xl">
    <div class="flex-1 flex justify-between sm:hidden">
      <button
        @click="prevPage"
        :disabled="store.page <= 1"
        class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md"
        :class="store.page <= 1 ? 'bg-gray-100 text-gray-400 cursor-not-allowed' : 'bg-white text-gray-700 hover:bg-gray-50'"
      >
        Anterior
      </button>
      <button
        @click="nextPage"
        :disabled="store.stocks.length < 10"
        class="ml-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md"
        :class="store.stocks.length < 10 ? 'bg-gray-100 text-gray-400 cursor-not-allowed' : 'bg-white text-gray-700 hover:bg-gray-50'"
      >
        Siguiente
      </button>
    </div>
    <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
      <div>
        <p class="text-sm text-gray-700">
          Página <span class="font-medium">{{ store.page }}</span>
        </p>
      </div>
      <div>
        <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
          <button
            @click="prevPage"
            :disabled="store.page <= 1"
            class="relative inline-flex items-center px-4 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium"
            :class="store.page <= 1 ? 'text-gray-300 cursor-not-allowed' : 'text-gray-700 hover:bg-gray-50'"
          >
            <span class="sr-only">Anterior</span>
            <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
              <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
            </svg>
          </button>
          <button
            @click="nextPage"
            :disabled="store.stocks.length < 10"
            class="relative inline-flex items-center px-4 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium"
            :class="store.stocks.length < 10 ? 'text-gray-300 cursor-not-allowed' : 'text-gray-700 hover:bg-gray-50'"
          >
            <span class="sr-only">Siguiente</span>
            <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
              <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
            </svg>
          </button>
        </nav>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useStocks } from '../stores/useStocks'

const store = useStocks()

function prevPage() {
  if (store.page > 1) {
    store.page--
    if (store.filters.profile) {
      store.fetchRecommendations()
    } else {
      store.fetchStocks()
    }
  }
}

function nextPage() {
  if (store.stocks.length === 10) {
    store.page++
    if (store.filters.profile) {
      store.fetchRecommendations()
    } else {
      store.fetchStocks()
    }
  }
}
</script>