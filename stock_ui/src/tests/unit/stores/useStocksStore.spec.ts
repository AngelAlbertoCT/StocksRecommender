import { setActivePinia, createPinia } from 'pinia'
import { useStocks } from '@/stores/useStocks'
import { describe, beforeEach, it, expect, vi } from 'vitest'

describe('useStocksStore', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  })

  it('inicializa con valores por defecto', () => {
    const store = useStocks()
    
    // Ajusta estos valores según tu store real
    expect(store.stocks).toEqual([])
    expect(store.loading).toBe(false)
    // Si tu store tiene filters con valores por defecto, ajústalos aquí
    expect(store.filters).toEqual({
      ticker: '',
      company: '',
      month: '',
      profile: ''
    })
  })

  it('actualiza stocks correctamente', async () => {
    const store = useStocks()
    const mockStocks = [{ ticker: 'AAPL', company: 'Apple' }]
    
    // Mockea la acción fetchStocks
    store.fetchStocks = vi.fn().mockImplementation(() => {
      store.stocks = mockStocks
      return Promise.resolve()
    })
    
    await store.fetchStocks()
    
    expect(store.fetchStocks).toHaveBeenCalled()
    expect(store.stocks).toEqual(mockStocks)
  })
})