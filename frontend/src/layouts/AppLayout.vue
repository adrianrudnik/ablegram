<template>
  <div class="AppLayout">
    <h1 class="text-3xl flex gap-2 font-medium">
      <RouterLink
        :to="{ name: 'search' }"
        class="inline-flex h-2rem w-2rem"
        v-if="!indicatorActive"
      >
        <img
          src="@/assets/media/logo.svg"
          class="inline-flex border-round-md border-1 border-transparent h-2rem mb-2"
          alt="Ablegram logo"
        />
      </RouterLink>
      <span v-if="indicatorActive" class="h-2rem w-2rem inline-flex">
        <ProgressSpinner
          class="h-auto"
          :pt="{ circle: { style: 'stroke: black !important; paint-order: stroke;' } }"
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
import { computed, ref, watch } from 'vue'
import { useStatStore } from '@/stores/stats'
import ProgressSpinner from 'primevue/progressspinner'
import { useI18n } from 'vue-i18n'
import { watchDebounced } from '@vueuse/core'

const { t } = useI18n()

// Create a slow motion progress indicator for the human eye.
// Any activity needs to be debounced to make the indicator visible long
// enough to perceive it. Otherwise, we would have a realtime flicker party.
// We also need to work around watchers reacting to assignments, not value changes
const indicatorActive = ref(false)
const inProgress = computed(() => useStatStore().inProgress)

watch(inProgress, () => {
  // Active the visible indicator, if not already active
  if (!indicatorActive.value) {
    indicatorActive.value = true
    return
  }
})

watchDebounced(
  inProgress,
  () => {
    // If the indicator was active, re-evaluate if it still should be active
    if (indicatorActive.value) {
      // ...and if not, turn it off
      indicatorActive.value = inProgress.value
    }
  },
  { debounce: 1000 }
)
</script>

<style lang="scss">
.AppLayout {
  padding-left: 1rem;
  padding-right: 1rem;
}
</style>

<i18n lang="yaml"></i18n>
