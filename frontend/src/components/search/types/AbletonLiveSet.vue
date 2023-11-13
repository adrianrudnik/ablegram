<template>
  <SearchResultCard :result="result" :expanded="expanded">
    <SearchCardAttribute>
      <i18n-t
        keypath="common.label.midi-track.nc"
        tag="span"
        :plural="props.result.midiTrackCount ?? 0"
      >
        <template v-slot:count>{{ n(props.result.midiTrackCount ?? 0) }}</template> </i18n-t
      >.
      <i18n-t
        keypath="common.label.audio-track.nc"
        tag="span"
        :plural="props.result.audioTrackCount ?? 0"
      >
        <template v-slot:count>{{ n(props.result.audioTrackCount ?? 0) }}</template> </i18n-t
      >.
    </SearchCardAttribute>

    <div v-if="expanded">
      <SearchCardAttribute
        v-if="props.result.scaleRootNote"
        :v="props.result.scaleRootNote"
        k="Scale root note"
      />
      <SearchCardAttribute
        v-if="props.result.scaleName"
        :v="props.result.scaleName"
        k="Scale name"
      />
      <SearchCardAttribute v-if="props.result.bpm" :k="t('common.label.beats-per-minute.s')">
        <i18n-t keypath="common.label.beats-per-minute.nc" tag="span" :plural="props.result.bpm">
          <template v-slot:count>
            {{ n(props.result.bpm) }}
          </template>
        </i18n-t>
      </SearchCardAttribute>
    </div>
  </SearchResultCard>
</template>

<script setup lang="ts">
import SearchResultCard from '@/components/search/SearchResultCard.vue'
import SearchCardAttribute from '@/components/search/SearchCardAttribute.vue'
import type { AbletonLiveSetResult } from '@/plugins/search/result/AbletonLiveSetResult'
import { useI18n } from 'vue-i18n'

const { t, n } = useI18n()

const props = withDefaults(
  defineProps<{
    result: AbletonLiveSetResult
    expanded?: boolean
  }>(),
  {
    expanded: false
  }
)
</script>
