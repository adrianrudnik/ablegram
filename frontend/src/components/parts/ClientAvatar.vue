<template>
  <Avatar
    icon="pi pi-user"
    class="mr-2"
    :style="{ 'background-color': colorizeUserClient(props.client).color, color: 'white' }"
    shape="circle"
    v-tooltip.left="tooltip"
  />
</template>

<script setup lang="ts">
import Avatar from 'primevue/avatar'
import type { UserClient } from '@/stores/users'
import { computed } from 'vue'
import { colorizeUserClient } from '@/stores/users'
import { useSessionStore } from '@/stores/session'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'

const props = defineProps<{ client: UserClient }>()
const { t } = useI18n()

const { clientId, userId } = storeToRefs(useSessionStore())

const tooltip = computed(() => {
  // Same user, different client
  if (props.client.user_id === userId.value && props.client.id !== clientId.value) {
    return t('user-avatar.you.different-client')
  }

  // Same user, same client
  if (props.client.id === clientId.value) {
    return t('user-avatar.you.same-client')
  }

  // Someone else's client
  const v = props.client.user_display_name
  return props.client.ip ? `${v} connecting from ${props.client.ip}` : v
})
</script>
