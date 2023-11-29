<template>
  <AvatarGroup v-if="entries.length > 1">
    <ClientAvatar
      :client="client"
      v-for="client in entries.slice(0, visibleMemberCount)"
      :key="client.id"
    />
    <Avatar
      v-if="entries.length > visibleMemberCount"
      :label="'+' + (entries.length - visibleMemberCount)"
      :pt="{ label: { class: 'text-sm' } }"
      shape="circle"
    />
  </AvatarGroup>
</template>

<script setup lang="ts">
import { useUserClientStore } from '@/stores/users'
import ClientAvatar from '@/components/parts/ClientAvatar.vue'
import Avatar from 'primevue/avatar'
import AvatarGroup from 'primevue/avatargroup'
import { computed } from 'vue'
import { breakpointsPrimeFlex, useBreakpoints } from '@vueuse/core'

const { entries } = useUserClientStore()

const { sm, md, lg } = useBreakpoints(breakpointsPrimeFlex)

const visibleMemberCount = computed(() => {
  if (lg.value) {
    return 10
  } else if (md.value) {
    return 7
  } else if (sm.value) {
    return 5
  } else {
    return 1
  }
})
</script>
