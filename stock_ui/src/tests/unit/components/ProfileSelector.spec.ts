import { mount } from '@vue/test-utils'
import ProfileSelector from '@/components/ProfileSelector.vue'
import { useStocks } from '@/stores/useStocks'
import { createPinia } from 'pinia'
import { describe, it, expect, vi } from 'vitest'

describe('ProfileSelector.vue', () => {
  it('interactúa correctamente con el store', async () => {
    const pinia = createPinia()
    const wrapper = mount(ProfileSelector, {
      global: {
        plugins: [pinia]
      }
    })

    const store = useStocks()
    store.fetchRecommendations = vi.fn()
    
    const buttons = wrapper.findAll('button')
    await buttons[1].trigger('click') // Segundo botón
    
    expect(store.filters.profile).toBe('moderate') // o 'moderado' según tu caso
    expect(store.fetchRecommendations).toHaveBeenCalled()
  })
})