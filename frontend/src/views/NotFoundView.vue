<template>
  <div class="AppLayout">
    <h1 class="">{{ t('error.404.title') }}</h1>
    <RouterLink :to="{ name: 'search' }">{{ t('error.404.back-to-search') }}</RouterLink>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>
