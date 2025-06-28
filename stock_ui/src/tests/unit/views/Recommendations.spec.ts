import { mount } from '@vue/test-utils'
import Recommendations from '@/views/RecommendationsPage.vue'
import { useStocks } from '@/stores/useStocks'
import { createPinia } from 'pinia'
import { describe, it, expect } from 'vitest'

describe('Recommendations.vue', () => {
  it('muestra el estado de carga correctamente', async () => {
    const pinia = createPinia()
    const store = useStocks(pinia)
    store.loading = true
    
    const wrapper = mount(Recommendations, {
      global: {
        plugins: [pinia],
        stubs: ['ProfileSelector', 'Pagination']
      }
    })
    
    expect(wrapper.find('.animate-spin').exists()).toBe(true)
  })

  it('muestra los stocks correctamente', async () => {
    const pinia = createPinia()
    const store = useStocks(pinia)
    store.stocks = [
      { ticker: 'AAPL', company: 'Apple', action: 'Comprar' }
    ]
    store.loading = false
    
    const wrapper = mount(Recommendations, {
      global: {
        plugins: [pinia],
        stubs: ['ProfileSelector', 'Pagination']
      }
    })
    
    await wrapper.vm.$nextTick()
    expect(wrapper.text()).toContain('AAPL')
    expect(wrapper.text()).toContain('Apple')
  })
})