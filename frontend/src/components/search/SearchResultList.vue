<template>
  <div class="grid">
    <div
      class="flex align-self-stretch col-12 md:col-6 lg:col-4 xl:col-3"
      v-for="result in results"
      :key="result.id"
    >
      <component
        :is="resolveComponent(result.type)"
        :result="result"
        :expanded="false"
        @click="openFocusDialog(result)"
      />
    </div>
  </div>

  <InfiniteScrollTrigger @trigger="loadMore" class="mb-6" />

  <div class="mx-auto my-5 text-center" v-if="isSearching && !isClean">
    <i class="pi pi-spin pi-spinner" style="font-size: 2rem"></i>
  </div>
</template>

<script lang="ts" setup>
import SearchResultCard from '@/components/search/SearchResultCard.vue'
import { useSearchResultStore } from '@/stores/results'
import { computed, shallowRef } from 'vue'
import type { HitFieldset, ResultType } from '@/plugins/search/result'
import InfiniteScrollTrigger from '@/components/structure/InfiniteScrollTrigger.vue'
import { useSearchStore } from '@/stores/search'
import AbletonMidiTrack from '@/components/search/types/AbletonMidiTrack.vue'
import { storeToRefs } from 'pinia'
import { useStatStore } from '@/stores/stats'
import { useDialog } from 'primevue/usedialog'
import SearchResultFocusDialog from '@/components/search/SearchResultFocusDialog.vue'
import { useI18n } from 'vue-i18n'
import AbletonSampleReference from '@/components/search/types/AbletonSampleReference.vue'
import AbletonMidiClip from '@/components/search/types/AbletonMidiClip.vue'
import AbletonLiveSet from '@/components/search/types/AbletonLiveSet.vue'

const { t } = useI18n()
const { loadMore } = useSearchStore()
const { isClean } = storeToRefs(useSearchStore())
const { isSearching } = storeToRefs(useStatStore())

function resolveComponent(type: ResultType): any {
  switch (type) {
    case 'AbletonLiveSet':
      return AbletonLiveSet
    // case 'AbletonAudioTrack':
    // case 'AbletonReturnTrack':
    // case 'AbletonGroupTrack':
    // case 'AbletonPreHearTrack':
    // case 'AbletonAudioClip':
    // case 'AbletonMixer':
    // case 'AbletonDeviceChain':
    // case 'AbletonScene':
    case 'AbletonMidiTrack':
      return AbletonMidiTrack
    case 'AbletonMidiClip':
      return AbletonMidiClip
    case 'AbletonSampleReference':
      return AbletonSampleReference
    default:
      return SearchResultCard
  }
}

const results = computed(() => useSearchResultStore().entries)
const dialog = useDialog()

const openFocusDialog = (result: HitFieldset) => {
  dialog.open(SearchResultFocusDialog, {
    props: {
      header: t('search-result-focus-dialog.header'),
      modal: true,
      dismissableMask: true,
      style: {
        width: '50vw'
      },
      breakpoints: {
        '1200px': '60vw',
        '960px': '75vw',
        '640px': '90vw'
      }
    },
    data: {
      component: shallowRef(resolveComponent(result.type)),
      result: result,
      expanded: true
    }
  })
}
</script>
