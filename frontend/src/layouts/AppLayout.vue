<template>
  <div class="AppLayout">
    <h1 class="text-3xl flex gap-2 font-medium">
      <ProgressLogo />

      Ablegram
    </h1>

    <MenuBar class="mb-3" />

    <RouterView />
  </div>

  <ConfirmDialog />
  <DynamicDialog />
</template>

<script setup lang="ts">
import MenuBar from '@/components/parts/MenuBar.vue'
import ConfirmDialog from 'primevue/confirmdialog'
import ProgressLogo from '@/components/parts/ProgressLogo.vue'
import DynamicDialog from 'primevue/dynamicdialog'
</script>

<style lang="scss">
.AppLayout {
  padding-left: 1rem;
  padding-right: 1rem;
}
</style>
