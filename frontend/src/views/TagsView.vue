<template>

  <SectionHeadline :title="t('tag-overview.type.title')">
    <template #description>
      <p>{{ t('tag-overview.type.description') }}</p>
    </template>

    <InputText v-model="filter" placeholder="Filter" class="w-full mb-3" />

    <SearchTag :tag="tag" v-for="tag in entries" :key="tag.id" show-count />
  </SectionHeadline>
</template>

<script setup lang="ts">
import { useTagStore } from '@/stores/tags'
import {computed, ref} from 'vue'
import SearchTag from '@/components/structure/SearchTag.vue'
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import InputText from "primevue/inputtext";
import sortBy from 'lodash/sortBy'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const filter = ref('')

const entries = computed(() => sortBy(useTagStore().entries.filter(v => v.search.includes(filter.value.toLowerCase())), 'count:desc'))
</script>
