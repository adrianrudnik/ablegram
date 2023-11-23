<template>
  <div class="AppLayout mt-3">
    <div class="flex gap-2 font-medium justify-content-between mb-3">
      <div class="flex gap-2 justify-content-center">
        <ProgressLogo />
        <h1 class="text-3xl p-0 m-0">Ablegram</h1>
      </div>
      <div class="flex gap-3 mt-1 align-items-center OutLinks">
        <a href="https://www.ablegram.app/" target="_blank">
          <i class="pi pi-fw pi-book text-xl"></i>
        </a>

        <a href="https://github.com/adrianrudnik/ablegram" target="_blank">
          <i class="pi pi-fw pi-github text-xl"></i>
        </a>

        <Avatar
          icon="pi pi-user"
          class="mr-2"
          :class="{ 'bg-black-alpha-90 text-white': isAdmin, 'text-black-alpha-90': isGuest }"
          @click="openUserPanel"
        />
        <OverlayPanel ref="userPanel">
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

          <Button :label="t('user-avatar.logout')" @click="goodbye" v-if="isAdmin" />
        </OverlayPanel>
        <ContextMenu ref="menu" :model="items" class="text-white background-black" />
      </div>
    </div>

    <MenuBar class="mb-3" />

    <RouterView />

    <ConfirmDialog />
    <DynamicDialog />
  </div>
</template>

<script setup lang="ts">
import MenuBar from '@/components/parts/MenuBar.vue'
import ProgressLogo from '@/components/parts/ProgressLogo.vue'
import DynamicDialog from 'primevue/dynamicdialog'
import ConfirmDialog from 'primevue/confirmdialog'
import ContextMenu from 'primevue/contextmenu'
import OverlayPanel from 'primevue/overlaypanel'
import Avatar from 'primevue/avatar'
import Button from 'primevue/button'
import { useSessionStore } from '@/stores/session'
import { storeToRefs } from 'pinia'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const { username, ip, isGuest, isAdmin } = storeToRefs(useSessionStore())
const { goodbye } = useSessionStore()

const menu = ref()
const items = ref([
  { label: 'Copy', icon: 'pi pi-copy' },
  { label: 'Rename', icon: 'pi pi-file-edit' }
])

const userPanel = ref()

const openUserPanel = (event: Event) => {
  userPanel.value.toggle(event)
}
</script>

<style lang="scss">
.AppLayout {
  .OutLinks {
    a {
      color: black;
    }
  }
  padding-left: 1rem;
  padding-right: 1rem;
}
</style>
