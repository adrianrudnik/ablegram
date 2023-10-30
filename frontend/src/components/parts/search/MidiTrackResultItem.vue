<template>
  <div>
    <SearchResultCard header="Midi Track" :title="result.displayName" :tags="props.result.tags">
      <PropertyList>
        <PropertyListItem v-if="result.pathFolder" title="Ordner" icon="pi pi-folder">
          {{ result.pathFolder }}
        </PropertyListItem>

        <PropertyListItem v-if="result.tempo" title="BPM" icon="pi pi-clock">
          {{ result.tempo }}
        </PropertyListItem>

        <PropertyListItem
          v-if="result.audioTrackCount || result.midiTrackCount"
          title="Spuren"
          icon="pi pi-bars"
        >
          {{ result.midiTrackCount }} MIDI, {{ result.audioTrackCount }} Audio
        </PropertyListItem>

        <PropertyListItem v-if="result.creator" title="Software" icon="pi pi-desktop">
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
          label="Open folder"
        />

        <Button
          size="small"
          outlined
          icon="pi pi-file-o"
          v-if="result.pathAbsolute"
          @click="openLocalPath(result.pathAbsolute)"
          label="Open file"
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
import DescriptionList from '@/components/structure/DescriptionList.vue'
import DescriptionListItem from '@/components/structure/DescriptionListItem.vue'
import PropertyList from '@/components/structure/PropertyList.vue'
import PropertyListItem from '@/components/structure/PropertyListItem.vue'

const props = defineProps<{
  result: LiveSetResult
}>()
</script>
