<template>
  <div class="bg-white px-6 py-4 rounded-t-xl border-b border-gray-200">
    <h2 class="text-lg font-medium text-gray-900 mb-4">Selecciona tu perfil de inversión</h2>
    <div class="flex flex-wrap gap-3">
      <button
        v-for="option in profiles"
        :key="option.value"
        @click="selectProfile(option.value)"
        :class="[
          'px-4 py-2 rounded-md text-sm font-medium transition-colors duration-200',
          store.filters.profile === option.value
            ? 'bg-indigo-600 text-white shadow-sm hover:bg-indigo-700'
            : 'bg-white text-gray-700 border border-gray-300 hover:bg-gray-50'
        ]"
      >
        {{ option.label }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useStocks } from '../stores/useStocks'

const store = useStocks()

const profiles = [
  { value: 'conservative', label: 'Conservador' },
  { value: 'moderate', label: 'Moderado' },
  { value: 'aggressive', label: 'Agresivo' },
]

function selectProfile(value: string) {
  store.filters.profile = value
  store.page = 1
  store.fetchRecommendations()
}
</script>