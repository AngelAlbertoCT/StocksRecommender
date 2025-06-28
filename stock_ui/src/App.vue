<template>
  <div class="min-h-screen bg-gray-50 flex flex-col">
    <!-- Header/Navigation -->
    <header class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16 items-center">
          <div class="flex-shrink-0 flex items-center">
            <h1 class="text-xl font-bold text-gray-900">Stock Analyzer</h1>
          </div>
          <nav class="flex space-x-8">
            <RouterLink 
              to="/stocks" 
              class="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
              active-class="text-blue-600 border-blue-500"
            >
              Análisis de Acciones
            </RouterLink>
            <RouterLink 
              to="/recommendations" 
              class="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
              active-class="text-blue-600 border-blue-500"
            >
              Recomendaciones
            </RouterLink>
          </nav>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="flex-grow">
      <RouterView />
    </main>

    <!-- Footer -->
    <footer class="bg-white border-t border-gray-200">
      <div class="max-w-7xl mx-auto py-4 px-4 sm:px-6 lg:px-8">
        <p class="text-center text-sm text-gray-500">
          &copy; 2023 Stock Analyzer. Todos los derechos reservados.
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>