<template>
  <div class="grid nested-grid mx-1 sm:mx-2 md:mx-4 lg:mx-8 text-gray-400">
    <div class="col-12 md:col-6 lg:col-4" v-for="metric in metrics" :key="metric['k']">
      <div class="grid text-sm">
        <div class="col-5 md:col-3 text-right p-1">{{ n(metric['v']) }}</div>
        <div class="col-7 md:col-9 p-1">{{ t('index.type.' + metric['k']) }}</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useStatStore } from '@/stores/stats'
import { useI18n } from 'vue-i18n'
import { computed } from 'vue'
import { orderBy } from 'lodash'

const { counters } = useStatStore()

const metrics = computed(() => orderBy(Object.entries(counters).map(([k, v]) => ({ k, v })), ['v'], ['desc']))

const { t, n } = useI18n()
</script>
