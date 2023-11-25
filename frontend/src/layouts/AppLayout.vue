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

        <a href="https://github.com/adrianrudnik/ablegram" target="_blank" v-if="sm">
          <i class="pi pi-fw pi-github text-xl"></i>
        </a>

        <ActiveUsers />
        <MyAvatar />
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
import MyAvatar from '@/components/auth/MyAvatar.vue'
import ActiveUsers from '@/components/parts/ActiveUsers.vue'
import { breakpointsPrimeFlex, useBreakpoints } from '@vueuse/core'

const { sm } = useBreakpoints(breakpointsPrimeFlex)
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
