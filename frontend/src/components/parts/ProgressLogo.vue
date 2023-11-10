<template>
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
</template>

<script setup lang="ts">
import ProgressSpinner from "primevue/progressspinner";
import {computed, ref} from "vue";
import {useStatStore} from "@/stores/stats";
import {watchThrottled} from "@vueuse/core";
import {useI18n} from "vue-i18n";

const {t} = useI18n()

const indicatorActive = ref(false)
const inProgress = computed(() => useStatStore().inProgress)

watchThrottled(inProgress, () => {
  indicatorActive.value = inProgress.value
}, {throttle: 125})
</script>

