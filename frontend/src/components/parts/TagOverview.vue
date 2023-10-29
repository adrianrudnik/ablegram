<template>
  <SectionHeadline :title="t('tag-overview.explain.title')">
    <template #description>
      <p>{{ t('tag-overview.explain.description') }}</p>
    </template>

    <SearchTag v-if="exampleTag" :tag="exampleTag" show-count :disable-translation="true" />
  </SectionHeadline>

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

    <div class="mb-3" v-if="fileYears">
      <SearchTag :tag="tag" v-for="tag in fileYears" :key="tag.id" show-count />
    </div>

    <div class="mb-3" v-if="fileMonths">
      <SearchTag :tag="tag" v-for="tag in fileMonths" :key="tag.id" show-count />
    </div>

    <div class="mb-3" v-if="fileQuarters">
      <SearchTag :tag="tag" v-for="tag in fileQuarters" :key="tag.id" show-count />
    </div>

    <div class="mb-3" v-if="fileWeekdays">
      <SearchTag :tag="tag" v-for="tag in fileWeekdays" :key="tag.id" show-count />
    </div>

    <div class="mb-3" v-if="fileWeekNumbers">
      <SearchTag :tag="tag" v-for="tag in fileWeekNumbers" :key="tag.id" show-count />
    </div>

    <div class="mb-3">
      <SearchTag :tag="tag" v-for="tag in westernZodiacs" :key="tag.id" show-count />
    </div>

    <div class="mb-3">
      <SearchTag :tag="tag" v-for="tag in chineseZodiacs" :key="tag.id" show-count />
    </div>
  </SectionHeadline>
</template>

<script setup lang="ts">
import { createTagFromString, TagCategory, useTagStore } from '@/stores/tags'
import { computed } from 'vue'
import SearchTag from '@/components/structure/SearchTag.vue'
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import sortBy from 'lodash/sortBy'
import orderBy from 'lodash/orderBy'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const exampleTag = createTagFromString(
  'sys:example:topic:detail:value',
  t('tags.sys:example:topic:detail:count')
)

const entries = computed(() => useTagStore().entries)

const softwareTags = computed(() =>
  orderBy(
    entries.value.filter((t) => t.category === TagCategory.Software),
    ['value'],
    ['desc']
  )
)

const locationTags = computed(() =>
  sortBy(
    entries.value.filter((t) => t.category === TagCategory.Location),
    'id'
  )
)

const liveSetTags = computed(() =>
  sortBy(
    entries.value.filter((t) => t.category === TagCategory.Tracks),
    'id'
  )
)

const tempoTags = computed(() =>
  sortBy(
    entries.value.filter((t) => t.category === TagCategory.Tempo),
    'value'
  )
)

const fileYears = computed(() =>
  sortBy(
    entries.value.filter((t) => t.category === TagCategory.Year),
    ['id', 'value']
  )
)

const fileMonths = computed(() =>
  sortBy(
    entries.value.filter((t) => t.category === TagCategory.Month),
    ['id', 'value']
  )
)

const fileQuarters = computed(() =>
  sortBy(
    entries.value.filter((t) => t.category === TagCategory.Quarter),
    ['id', 'value']
  )
)

const fileWeekdays = computed(() =>
  sortBy(
    entries.value.filter((t) => t.category === TagCategory.Weekday),
    ['id', 'value']
  )
)

const fileWeekNumbers = computed(() =>
  sortBy(
    entries.value.filter((t) => t.category === TagCategory.WeekNumber),
    ['id', 'value']
  )
)

const westernZodiacs = computed(() =>
  sortBy(
    entries.value.filter((t) => t.category === TagCategory.WesternZodiac),
    'value'
  )
)

const chineseZodiacs = computed(() =>
  sortBy(
    entries.value.filter((t) => t.category === TagCategory.ChineseZodiac),
    'value'
  )
)
</script>
