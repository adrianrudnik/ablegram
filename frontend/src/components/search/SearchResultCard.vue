<template>
  <Card class="SearchResultCard">
    <template #header v-if="color">
      <div class="ColorBox" :style="'background-color: ' + color"></div>
    </template>

    <template #title>
      <div class="title" v-if="props.result.displayName">{{ props.result.displayName }}</div>
      <div class="title font-italic text-gray-400" v-else>{{ t('common.label.no-name') }}</div>
      <div class="type">{{ t('index.type.' + props.result.type) }}</div>
    </template>

    <template #content>
      <div class="filename">{{ props.result.filename }}</div>
      <div class="folder" v-if="variant === 'default'">{{ props.result.pathFolder }}</div>
      <div class="user-memo" v-if="userMemo">
        <MoreText :text="userMemo" :expanded="variant === 'default'" />
      </div>

      <slot />

      <div v-if="variant === 'default' && props.result.tags" class="mt-3">
        <div class="tags">
          <TagRow :tags="props.result.tags" />
        </div>
      </div>
    </template>
  </Card>
</template>

<script setup lang="ts">
import Card from 'primevue/card'
import type { HitFieldset } from '@/plugins/search/result'
import { resolveAbletonColorByIndex } from '@/plugins/colors'
import { useI18n } from 'vue-i18n'
import MoreText from '@/components/structure/MoreText.vue'
import TagRow from '@/components/parts/search/TagRow.vue'

const { t } = useI18n()

const props = defineProps<{
  result: HitFieldset
  variant?: 'compact' | 'default'
}>()

const color = 'color' in props.result ? resolveAbletonColorByIndex(props.result.color) : undefined
const userMemo = 'annotation' in props.result ? props.result.annotation : undefined
</script>

<style lang="scss">
.SearchResultCard {
  cursor: pointer;
  flex-grow: 1;
  position: relative;
  border: 2px solid black;
  box-shadow: unset;

  &:hover {
    background-color: var(--surface-100);
  }

  .p-card-body {
    padding: 1rem 2rem 1rem 1rem;
  }

  .p-card-content {
    > div {
      margin-bottom: 0.4rem;
    }

    .user-memo {
      border-left: 4px solid var(--surface-300);
      padding-left: 6px;
    }
  }

  .p-card-footer {
    .p-button {
      scale: (0.8);
    }
  }

  .p-card-title {
    display: flex;
    flex-direction: column;

    .title {
      font-size: medium;
      margin-bottom: 1px;
    }

    .type {
      font-size: small;
      font-weight: normal;
    }
  }

  .ColorBox {
    position: absolute;
    top: 0;
    right: 0;

    display: inline-block;
    height: 100%;
    width: 16px;
    background-color: black;
    border-radius: 4px;
    border-top-left-radius: 0 !important;
    border-bottom-left-radius: 0 !important;
  }
}
</style>
