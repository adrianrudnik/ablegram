<template>
  <div class="AppLayout">
    <h1 class="text-3xl">Ablegram</h1>

    <MenuBar class="mb-3" />

    <RouterView />
  </div>
</template>

<script setup lang="ts">
import MenuBar from '@/components/parts/MenuBar.vue'
</script>

<style lang="scss">
.AppLayout {
  padding-left: 1rem;
  padding-right: 1rem;
}
</style>

<i18n lang="yaml"></i18n>
