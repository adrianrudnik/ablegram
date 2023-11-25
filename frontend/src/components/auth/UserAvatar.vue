<template>
  <Avatar
    icon="pi pi-user"
    class="mr-2"
    :style="{ 'background-color': colorizeUser(props.user).color, color: 'white' }"
    shape="circle"
    v-tooltip.left="tooltip"
  />
</template>

<script setup lang="ts">
import Avatar from 'primevue/avatar'
import type { User } from '@/stores/users'
import { computed } from 'vue'
import { colorizeUser } from '@/stores/users'
import { useSessionStore } from '@/stores/session'
import { useI18n } from 'vue-i18n'

const props = defineProps<{ user: User }>()
const { t } = useI18n()

const sessionStore = useSessionStore()

const tooltip = computed(() => {
  if (props.user.id === sessionStore.id) {
    return t('user-avatar.you')
  }

  const v = props.user.display_name
  return props.user.ip ? `${v} connecting from ${props.user.ip}` : v
})
</script>
