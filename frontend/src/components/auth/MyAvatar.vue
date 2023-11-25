<template>
  <Avatar
    icon="pi pi-user"
    class="mr-2"
    :class="{ 'bg-black-alpha-90 text-white': isAdmin, 'text-black-alpha-90': isGuest }"
    @click="openUserPanel"
  />
  <OverlayPanel ref="userPanel" class="w-full md:w-24rem">
    <div class="mb-3">
      <p class="font-semibold">
        {{ username }} [{{ isAdmin ? t('role.admin') : t('role.guest') }}]
      </p>
      <i18n-t keypath="user-avatar.from-ip" tag="p">
        <template v-slot:ip>
          <code class="text-sm p-1 bg-black-alpha-10">{{ ip }}</code>
        </template>
      </i18n-t>
    </div>

    <LoginWithPasswordForm v-if="isGuest" />
    <LogoutForm v-if="!isGuest" />
  </OverlayPanel>
</template>

<script setup lang="ts">
import OverlayPanel from 'primevue/overlaypanel'
import LoginWithPasswordForm from '@/components/auth/LoginWithPasswordForm.vue'
import LogoutForm from '@/components/auth/LogoutForm.vue'
import Avatar from 'primevue/avatar'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import { useSessionStore } from '@/stores/session'
import { ref } from 'vue'

const { t } = useI18n()
const { username, ip, isGuest, isAdmin } = storeToRefs(useSessionStore())

const userPanel = ref()

const openUserPanel = (event: Event) => {
  userPanel.value.toggle(event)
}
</script>
