<template>
  <Button
    :disabled="isSuccess || props.loading"
    :icon="isSuccess ? 'pi pi-check' : loading ? 'pi pi-spin pi-spinner' : undefined"
    :label="statusLabel"
    :severity="isSuccess ? 'success' : undefined"
    type="submit"
    v-bind="$attrs"
  />
</template>

<script lang="ts" setup>
import Button from 'primevue/button'
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const props = withDefaults(
  defineProps<{
    label?: string
    success?: boolean | undefined
    loading?: boolean | undefined
  }>(),
  {
    label: undefined,
    success: undefined,
    loading: undefined
  }
)

const statusLabel = computed(() => {
  if (isSuccess.value) return t('button.state-success')
  if (props.loading) return t('button.state-loading')
  return props.label ?? t('button.submit')
})

const isSuccess = computed(() => props.success)
</script>
