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

const props = defineProps<{
  user: User
}>()

const tooltip = computed(() => {
  const v = props.user.display_name
  return props.user.ip ? `${v} connecting from ${props.user.ip}` : v
})
</script>
