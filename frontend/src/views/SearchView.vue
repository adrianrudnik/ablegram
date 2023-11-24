<template>
  <SearchQueryInput class="z-2 relative" />
  <SearchQueryToolbar class="mb-3 z-1 relative" />
  <ActiveFilterBar class="mb-3 z-0" />
  <SearchResultList />
  <MetricCloud v-if="isClean" />
</template>

<script setup lang="ts">
import SearchResultList from '@/components/search/SearchResultList.vue'
import SearchQueryInput from '@/components/search/SearchQueryInput.vue'
import ActiveFilterBar from '@/components/search/ActiveFilterBar.vue'
import SearchQueryToolbar from '@/components/search/SearchQueryToolbar.vue'
import { nextTick, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useSearchStore } from '@/stores/search'
import { createTagFromString } from '@/stores/tags'
import { ActiveFilterMode, ActiveFilterType, useActiveFiltersStore } from '@/stores/search-filters'
import MetricCloud from '@/components/structure/MetricCloud.vue'

const router = useRouter()
const route = useRoute()

const { currentQueryString, isClean } = storeToRefs(useSearchStore())

const { update: addActiveFilter } = useActiveFiltersStore()

onMounted(() => {
  const q = new URL(window.location.href).searchParams

  let migrate = false

  // We use the next tick, to give the store enough time to install all required watchers.
  nextTick(() => {
    // Transport the query string into the current search object
    if (q.has('q')) {
      currentQueryString.value = q.get('q')!
      migrate = true
    }

    if (q.has('tag')) {
      const tags = q.getAll('tag')

      for (let tagId of tags) {
        console.log('tag', tagId)
        // Extract the binary expression, if available
        let match = ActiveFilterMode.SHOULD
        switch (tagId.charAt(0)) {
          case '+':
            match = ActiveFilterMode.MUST
            tagId = tagId.substring(1)
            break
          case '-':
            match = ActiveFilterMode.NOT
            tagId = tagId.substring(1)
            break
        }

        const tag = createTagFromString(tagId, 0)

        addActiveFilter({
          id: tag.id,
          mode: match,
          content: tag,
          query: 'tags:"' + tag.id + '"',
          type: ActiveFilterType.TAG
        })

        migrate = true
      }
    }

    // If we have something migrated, replace to initial uri with variant without query string parts.
    if (migrate) router.replace(route.path)
  })
})
</script>
