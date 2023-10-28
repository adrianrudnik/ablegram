<template>
  <SectionHeadline :title="t('tag-overview.software.title')">
    <template #description>
      <p>{{ t('tag-overview.software.description') }}</p>
    </template>

    <SearchTag :tag="tag" v-for="tag in softwareTags" :key="tag.id" show-count />
  </SectionHeadline>

  <SectionHeadline :title="t('tag-overview.location.title')">
    <template #description>
      <p>{{ t('tag-overview.location.description') }}</p>
    </template>

    <SearchTag :tag="tag" v-for="tag in locationTags" :key="tag.id" show-count />
  </SectionHeadline>

  <SectionHeadline :title="t('tag-overview.live-set.title')">
    <template #description>
      <p>{{ t('tag-overview.live-set.description') }}</p>
    </template>

    <div class="mb-3">
      <SearchTag :tag="tag" v-for="tag in liveSetTags" :key="tag.id" show-count />
    </div>

    <SearchTag :tag="tag" v-for="tag in tempoTags" :key="tag.id" show-count />
  </SectionHeadline>

  <SectionHeadline :title="t('tag-overview.file.title')">
    <template #description>
      <p>{{ t('tag-overview.file.description') }}</p>
    </template>

    <div class="mb-3" v-if="filesYear">
      <SearchTag :tag="tag" v-for="tag in filesYear" :key="tag.id" show-count />
    </div>

    <div class="mb-3" v-if="filesMonth">
      <SearchTag :tag="tag" v-for="tag in filesMonth" :key="tag.id" show-count />
    </div>

    <div class="mb-3" v-if="filesQuarter">
      <SearchTag :tag="tag" v-for="tag in filesQuarter" :key="tag.id" show-count />
    </div>

    <div class="mb-3" v-if="filesWeekday">
      <SearchTag :tag="tag" v-for="tag in filesWeekday" :key="tag.id" show-count />
    </div>

    <div class="mb-3" v-if="filesWeekNo">
      <SearchTag :tag="tag" v-for="tag in filesWeekNo" :key="tag.id" show-count />
    </div>

    <div class="mb-3">
      <SearchTag :tag="tag" v-for="tag in wZodiacs" :key="tag.id" show-count />
    </div>

    <div class="mb-3">
      <SearchTag :tag="tag" v-for="tag in cZodiacs" :key="tag.id" show-count />
    </div>
  </SectionHeadline>
</template>

<script setup lang="ts">
import { useTagStore } from '@/stores/tags'
import { computed } from 'vue'
import SearchTag from '@/components/structure/SearchTag.vue'
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import sortBy from 'lodash/sortBy'
import orderBy from 'lodash/orderBy'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const softwareTags = computed(() =>
  orderBy(
    useTagStore().entries.filter((tag) => tag.id.startsWith('sys:ableton:')),
    ['value'],
    ['desc']
  )
)

const locationTags = computed(() =>
  sortBy(
    useTagStore().entries.filter((tag) => tag.id.startsWith('sys:location:')),
    'id'
  )
)

const liveSetTags = computed(() =>
  sortBy(
    useTagStore().entries.filter(
      (tag) => tag.id.startsWith('sys:live-set:') && !tag.id.includes(':tempo')
    ),
    'id'
  )
)

const tempoTags = computed(() =>
  sortBy(
    useTagStore().entries.filter((tag) => tag.id.startsWith('sys:live-set:tempo')),
    'value'
  )
)

const filesYear = computed(() =>
  sortBy(
    useTagStore().entries.filter((tag) => tag.id.includes('time-year:')),
    ['id', 'value']
  )
)

const filesMonth = computed(() =>
  sortBy(
    useTagStore().entries.filter((tag) => tag.id.includes('time-month:')),
    ['id', 'value']
  )
)

const filesQuarter = computed(() =>
  sortBy(
    useTagStore().entries.filter((tag) => tag.id.includes('time-quarter:')),
    ['id', 'value']
  )
)

const filesWeekday = computed(() =>
  sortBy(
    useTagStore().entries.filter((tag) => tag.id.includes('time-weekday:')),
    ['id', 'value']
  )
)

const filesWeekNo = computed(() =>
  sortBy(
    useTagStore().entries.filter((tag) => tag.id.includes('time-weekno:')),
    ['id', 'value']
  )
)

const wZodiacs = computed(() =>
  sortBy(
    useTagStore().entries.filter((tag) => tag.id.startsWith('sys:file:zodiac-western:')),
    'value'
  )
)

const cZodiacs = computed(() =>
  sortBy(
    useTagStore().entries.filter((tag) => tag.id.startsWith('sys:file:zodiac-chinese:')),
    'value'
  )
)
</script>
