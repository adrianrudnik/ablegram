<template>
  <div>
    <SearchResultCard
      :header="t('common.label.ableton-live-set.s')"
      :title="result.displayName"
      :tags="props.result.tags"
    >
      <PropertyList>
        <PropertyListItem
          v-if="result.pathFolder"
          :title="t('common.label.fs-folder.s')"
          icon="pi pi-folder"
        >
          {{ result.pathFolder }}
        </PropertyListItem>

        <PropertyListItem
          v-if="result.bpm"
          :title="t('common.label.beats-per-minute.s')"
          icon="pi pi-clock"
        >
          {{ result.bpm }}
        </PropertyListItem>

        <PropertyListItem
          v-if="result.audioTrackCount || result.midiTrackCount"
          :title="t('common.label.ableton-track.c', { count: 2 })"
          icon="pi pi-bars"
        >
          {{ t('common.label.ableton-midi-track.nc', { count: result.midiTrackCount }) }}.
          {{ t('common.label.ableton-audio-track.nc', { count: result.audioTrackCount }) }}.
        </PropertyListItem>

        <PropertyListItem
          v-if="result.creator"
          :title="t('common.label.software.s')"
          icon="pi pi-desktop"
        >
          {{ result.creator }}
        </PropertyListItem>
      </PropertyList>

      <template #actions>
        <Button
          size="small"
          outlined
          icon="pi pi pi-folder-open"
          v-if="result.pathFolder"
          @click="openLocalPath(result.pathFolder)"
          :label="t('common.button.open-folder')"
        />

        <Button
          size="small"
          outlined
          icon="pi pi-file-o"
          v-if="result.pathAbsolute"
          @click="openLocalPath(result.pathAbsolute)"
          :label="t('common.button.open-file')"
        />
      </template>
    </SearchResultCard>
  </div>
</template>

<script setup lang="ts">
import type { LiveSetResult } from '@/plugins/search/result/result_live_set'
import SearchResultCard from '@/components/structure/SearchResultCard.vue'
import Button from 'primevue/button'
import { openLocalPath } from '@/plugins/api'
import PropertyList from '@/components/structure/PropertyList.vue'
import PropertyListItem from '@/components/structure/PropertyListItem.vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const props = defineProps<{
  result: LiveSetResult
}>()
</script>
