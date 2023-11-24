<template>
  <SectionHeadline :title="t('tag-overview.title')">
    <template #description>
      <p>{{ t('tag-overview.intro') }}</p>
      <p>{{ t('tag-overview.filtering') }}</p>
    </template>

    <InputText v-model="filter" placeholder="Filter" class="w-full mb-3" />

    <SearchTag
      class="cursor-pointer mb-1 mr-1"
      :tag="tag"
      v-for="tag in entries"
      :key="tag.id"
      show-count
      @click="copyTag(tag)"
    />
  </SectionHeadline>
</template>

<script setup lang="ts">
import { useTagStore } from '@/stores/tags'
import { computed, ref } from 'vue'
import SearchTag from '@/components/structure/SearchTag.vue'
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import InputText from 'primevue/inputtext'
import { useI18n } from 'vue-i18n'
import orderBy from 'lodash/orderBy'
import type { Tag } from '@/stores/tags'
import { useToast } from 'primevue/usetoast'
import { ActiveFilterMode, ActiveFilterType, useActiveFiltersStore } from '@/stores/search-filters'

const { t } = useI18n()

const filter = ref('')

const toast = useToast()

const copyTag = (tag: Tag) => {
  const v = 'tags:"' + tag.id + '"'

  useActiveFiltersStore().update({
    id: tag.id,
    type: ActiveFilterType.TAG,
    mode: ActiveFilterMode.SHOULD,
    content: tag,
    query: 'tags:"' + tag.id + '"'
  })

  toast.add({
    severity: 'info',
    summary: t('toast.tag-added-to-filter.title'),
    detail: t('toast.tag-added-to-filter.detail', { tag: tag.id }),
    life: 3000
  })
}

const entries = computed(() => {
  if (filter.value.trim() != '') {
    return useTagStore().entries.filter((v) => v.search.includes(filter.value.toLowerCase()))
  } else {
    return orderBy(useTagStore().entries, ['count'], ['desc'])
  }
})
</script>
