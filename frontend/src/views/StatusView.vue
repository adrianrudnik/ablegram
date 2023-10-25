<template>
  <DescriptionList title="App status">
    <DescriptionListItem :title="t('status.version.label')">
      {{ t('status.version.content', { version: version, hash: commitHash }) }}
    </DescriptionListItem>

    <DescriptionListItem :title="t('status.websocket.label')">
      <i v-if="websocketStatus === 'OPEN'" class="pi pi-fw pi-check-circle text-green-500"></i>
      <i v-else-if="websocketStatus === 'CONNECTING'" class="pi pi-fw pi-spin pi-spinner"></i>
      <i v-else-if="websocketStatus === 'CLOSED'" class="pi pi-fw pi-times-circle text-red-500"></i>

      {{ t('status.websocket.status.' + websocketStatus, { url: websocketUrl}) }}
    </DescriptionListItem>
  </DescriptionList>
</template>

<script setup lang="ts">
import DescriptionList from '@/components/structure/DescriptionList.vue'
import DescriptionListItem from '@/components/structure/DescriptionListItem.vue'
import { useUiStore } from '@/stores/ui'
import { useI18n } from 'vue-i18n'
import { websocket } from '@/plugins/websocket'

const { t } = useI18n()

const version = useUiStore().version
const commitHash = useUiStore().versionCommitHash
const websocketStatus = websocket.status
const websocketUrl = import.meta.env.VITE_WEBSOCKET_URL
</script>
