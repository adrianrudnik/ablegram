<template>
  <div class="MetaBar surface-section px-0">
    <ul class="list-none p-0 m-0 surface-900 flex overflow-y-hidden overflow-x-auto border-round">
      <li
        class="relative py-3 pl-4 pr-3 flex flex-column justify-content-center"
        v-for="({ k, v }, idx) in stats"
        :key="idx"
        :class="{ 'pl-6': idx > 0 }"
      >
        <div
          class="absolute left-0 top-0 z-1"
          style="
            border-left: 20px solid var(--surface-900);
            border-top: 45px solid transparent;
            border-bottom: 45px solid transparent;
            width: 0;
            height: 0;
          "
        ></div>
        <div class="text-xl font-medium text-white mb-1">{{ v.value }}</div>
        <span class="text-white white-space-nowrap">{{ t('stat.' + k, v.value) }}</span>
        <div
          v-if="idx > 0"
          class="absolute top-0"
          style="
            left: 1px;
            border-left: 20px solid var(--surface-300);
            border-top: 45px solid transparent;
            border-bottom: 45px solid transparent;
            width: 0;
            height: 0;
          "
        ></div>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { useStatStore } from '@/stores/stats'
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useFilesStore } from '@/stores/files'

const { t } = useI18n()

const filesCount = computed(() => useFilesStore().count)
const midiTracks = computed(() => useStatStore().midiTracks)
const audioTracks = computed(() => useStatStore().audioTracks)

const stats = computed(() => {
  return [
    { k: 'files', v: filesCount },
    { k: 'midi-tracks', v: midiTracks },
    { k: 'audio-tracks', v: audioTracks }
  ]
})
</script>
