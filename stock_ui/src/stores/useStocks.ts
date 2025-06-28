// src/store/useStocks.ts
import { defineStore } from 'pinia'
import axios from 'axios'

const API_BASE = 'http://localhost:8080/api'

export const useStocks = defineStore('stocks', {
  state: () => ({
    stocks: [] as any[],
    loading: false,
    page: 1,
    filters: {
      ticker: '',
      company: '',
      month: '', 
      profile: '',
    },
  }),
  actions: {
    async fetchStocks() {
      this.loading = true
      try {
        const params: any = { page: this.page }

        if (this.filters.ticker) {
          params.ticker = this.filters.ticker
        }

        if (this.filters.company) {
          params.company = this.filters.company
        }

        if (this.filters.month) {
          params.month = this.filters.month
        }

        const res = await axios.get(`${API_BASE}/stocks`, { params })
        this.stocks = res.data
      } catch (err) {
        console.error('Error fetching stocks', err)
        this.stocks = []
      } finally {
        this.loading = false
      }
    },

    async fetchRecommendations() {
      if (!this.filters.profile) return
      this.loading = true
      try {
        const res = await axios.get(`${API_BASE}/recommendations`, {
          params: {
            profile: this.filters.profile,
            page: this.page,
          },
        })
        this.stocks = res.data
      } catch (err) {
        console.error('Error fetching recommendations', err)
        this.stocks = []
      } finally {
        this.loading = false
      }
    },

    setPage(newPage: number) {
      this.page = newPage
      if (this.filters.profile) {
        this.fetchRecommendations()
      } else {
        this.fetchStocks()
      }
    },
  },
})
