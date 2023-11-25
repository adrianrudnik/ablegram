<template>
  <Button :label="t('user-avatar.logout')" @click="logout" v-if="isAdmin" />
</template>

<script setup lang="ts">
import Button from 'primevue/button'
import { useI18n } from 'vue-i18n'
import { useSessionStore } from '@/stores/session'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'

const { t } = useI18n()
const { goodbye, reconsider } = useSessionStore()
const router = useRouter()

const { isAdmin } = storeToRefs(useSessionStore())

const logout = async () => {
  await goodbye()
  await reconsider()

  await router.push({ name: 'app' })
}
</script>
