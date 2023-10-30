<template>
  <div>
    <SearchResultCard
      :header="t('common.label.audio-track.s')"
      :title="props.result.displayName"
      :tags="props.result.tags"
    >
      <PropertyList>
        <PropertyListItem
          v-if="result.filename"
          :title="t('common.label.fs-filename.s')"
          icon="pi pi-file"
        >
          {{ result.filename }}
        </PropertyListItem>

        <PropertyListItem
          v-if="result.pathFolder"
          :title="t('common.label.fs-folder.s')"
          icon="pi pi-folder"
        >
          {{ result.pathFolder }}
        </PropertyListItem>

        <PropertyListItem
          v-if="result.annotation"
          :title="t('common.label.annotation.s')"
          icon="pi pi-user-edit"
        >
          {{ result.annotation }}
        </PropertyListItem>

        <PropertyListItem
          v-if="resolvedColor"
          :title="t('common.label.color.s')"
          icon="pi pi-palette"
        >
          <div class="ColorProperty" :style="'background-color: ' + resolvedColor"></div>
          {{ t('color.' + result.color) }}
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
import SearchResultCard from '@/components/structure/SearchResultCard.vue'
import Button from 'primevue/button'
import { openLocalPath } from '@/plugins/api'
import PropertyList from '@/components/structure/PropertyList.vue'
import PropertyListItem from '@/components/structure/PropertyListItem.vue'
import { useI18n } from 'vue-i18n'
import type { AudioTrackResult } from '@/plugins/search/result/result_audio_track'
import { resolveColorByIndex } from '@/plugins/colors'

const { t } = useI18n()

const props = defineProps<{
  result: AudioTrackResult
}>()

const resolvedColor = resolveColorByIndex(props.result.color)
</script>

<style lang="scss">
.ColorProperty {
  display: inline-block;
  height: 16px;
  width: 16px;
  background-color: black;
}
</style>
