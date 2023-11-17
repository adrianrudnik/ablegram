<template>
  <div class="AppStatus" v-if="about">
    <SectionHeadline :title="t('app-status.title')">
      <template #description>
        {{ t('app-status.description') }}
      </template>
    </SectionHeadline>

    <DescriptionList>
      <DescriptionListItem :title="t('app-status.version.label')">
        {{
          t('app-status.version.content', {
            version: about.version,
            hash: about.commit,
            date: about.date
          })
        }}
      </DescriptionListItem>

      <DescriptionListItem :title="t('app-status.websocket.label')">
        <i v-if="websocketStatus === 'OPEN'" class="pi pi-fw pi-check-circle text-green-500"></i>
        <i v-else-if="websocketStatus === 'CONNECTING'" class="pi pi-fw pi-spin pi-spinner"></i>
        <i
          v-else-if="websocketStatus === 'CLOSED'"
          class="pi pi-fw pi-times-circle text-red-500"
        ></i>

        {{ t('app-status.websocket.status.' + websocketStatus, { url: websocketUrl }) }}
      </DescriptionListItem>
    </DescriptionList>
  </div>
</template>

<script setup lang="ts">
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import DescriptionList from '@/components/structure/DescriptionList.vue'
import DescriptionListItem from '@/components/structure/DescriptionListItem.vue'
import { useI18n } from 'vue-i18n'
import { getWebsocketUrl, websocket } from '@/websocket'
import { onMounted, ref } from 'vue'
import { fetchApi } from '@/plugins/api'

const websocketStatus = websocket.status
const websocketUrl = getWebsocketUrl()

const { t } = useI18n()

type AboutResponse = {
  commit: string
  date: string
  version: string
}

const about = ref<AboutResponse>()

onMounted(async () => {
  about.value = await fetchApi<AboutResponse>('/about')
})
</script>
