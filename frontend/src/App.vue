<template>
  <RouterView />
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import { computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { preferredLocale } from 'preferred-locale'

onMounted(() => {
  // Ensure the html tag contains the correct language
  document.documentElement.lang = computed(() =>
    preferredLocale('en', Object.keys(useI18n().messages.value))
  ).value
})
</script>

<style>
html,
body {
  padding: 0;
  margin: 0;
}
</style>
