import { defineStore } from 'pinia'
import { setupStore } from '@/stores/base'
import type { HitFieldset } from '@/plugins/search/result'

export const useSearchResultStore = defineStore('search-results', setupStore<HitFieldset>())
