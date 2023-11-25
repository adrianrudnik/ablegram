<template>
  <AvatarGroup v-if="entries.length > 1">
    <UserAvatar :user="user" v-for="user in entries.slice(0, visibleMemberCount)" :key="user.id" />
    <Avatar
      v-if="entries.length > visibleMemberCount"
      :label="'+' + (entries.length - visibleMemberCount)"
      :pt="{ label: { class: 'text-sm' } }"
      shape="circle"
    />
  </AvatarGroup>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/users'
import UserAvatar from '@/components/auth/UserAvatar.vue'
import Avatar from 'primevue/avatar'
import AvatarGroup from 'primevue/avatargroup'
import { computed } from 'vue'
import { breakpointsPrimeFlex, useBreakpoints } from '@vueuse/core'

const { entries } = useUserStore()

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
