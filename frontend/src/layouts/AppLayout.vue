<template>
  <div class="AppLayout">
    <h1 class="text-3xl flex gap-2">
      <RouterLink :to="{ name: 'search' }" class="inline-flex h-2rem w-2rem" v-if="!processing">
        <img
          src="@/assets/media/logo.svg"
          class="inline-flex border-round-md border-1 border-transparent h-2rem mb-2"
          alt="Ablegram logo"
        />
      </RouterLink>
      <span v-if="processing" class="h-2rem w-2rem inline-flex">
        <ProgressSpinner
          class="h-auto"
          :pt="{ circle: { style: 'stroke: green !important; paint-order: stroke;' } }"
          fill="black"
          stroke-width="10"
          v-tooltip="t('app.status.processing')"
        />
      </span>

      Ablegram
    </h1>

    <MenuBar class="mb-3" />

    <RouterView />
  </div>

  <ConfirmDialog />
</template>

<script setup lang="ts">
import MenuBar from '@/components/parts/MenuBar.vue'
import ConfirmDialog from 'primevue/confirmdialog'
import { computed } from 'vue'
import { useStatStore } from '@/stores/stats'
import ProgressSpinner from 'primevue/progressspinner'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const processing = computed(() => useStatStore().isProcessing)
// const processing = true
</script>

<style lang="scss">
.AppLayout {
  padding-left: 1rem;
  padding-right: 1rem;
}
</style>

<i18n lang="yaml"></i18n>
