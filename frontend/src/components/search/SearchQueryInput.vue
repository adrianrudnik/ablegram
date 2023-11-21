<template>
  <div class="flex flex-column relative">
    <!-- total result counter above the query input bar -->
    <i18n-t
      keypath="search-query-input.total-hits"
      tag="div"
      class="ResultCounter"
      v-if="!isClean"
      :plural="totalHits"
    >
      <template v-slot:count>{{ n(totalHits) }}</template>
    </i18n-t>

    <InputText
      v-model="currentQueryString"
      class="w-full"
      :placeholder="t('search-query-input.placeholder')"
      v-focus
    />
  </div>
</template>

<script setup lang="ts">
import InputText from 'primevue/inputtext'
import { computed } from 'vue'
import { watchThrottled } from '@vueuse/core'
import { useStatStore } from '@/stores/stats'
import { useSearchStore } from '@/stores/search'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import { createQueryInstanceWithDefaults } from '@/plugins/search/query'

const { t, n } = useI18n()

const statStore = useStatStore()
const searchStore = useSearchStore()

const { currentQueryString, isClean, totalHits } = storeToRefs(searchStore)

watchThrottled(
  computed(() => currentQueryString.value),
  async (q) => {
    try {
      await useSearchStore().search({
        ...createQueryInstanceWithDefaults(),
        query: {
          query: q
        },
        sort: ['-_score', '_id']
      })
    } finally {
      statStore.isSearching = false
    }
  },
  { throttle: 500 }
)
</script>

<style scoped lang="scss">
.ResultCounter {
  font-size: 0.8rem;
  font-weight: lighter;
  color: var(--gray-600);
  top: -1rem;
  right: 0.25rem;
  position: absolute;
  z-index: 10;
}
</style>
