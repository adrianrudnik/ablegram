<template>
  <Button
    :disabled="showSuccessState || props.loading"
    :icon="showSuccessState ? 'pi pi-check' : loading ? 'pi pi-spin pi-spinner' : undefined"
    :label="statusLabel"
    :severity="showSuccessState ? 'success' : undefined"
    type="submit"
    v-bind="$attrs"
    class="w-auto"
  />
</template>

<script lang="ts" setup>
import Button from 'primevue/button'
import { computed, ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const emit = defineEmits(['update:modelValue'])

const props = withDefaults(
  defineProps<{
    modelValue: boolean
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
  if (showSuccessState.value) return t('button.state-success')
  if (props.loading) return t('button.state-loading')
  return props.label ?? t('button.submit')
})

// Construct some reactive state to show the success state, for a second, then hide it
// and make the button usable again.
const showSuccessState = computed({
  get: () => props.modelValue,
  set: (newValue) => emit('update:modelValue', newValue)
})

watch(
  () => props.modelValue,
  (newValue) => {
    if (newValue) {
      showSuccessState.value = newValue
      setTimeout(() => {
        showSuccessState.value = false
      }, 1000)
    }
  }
)
</script>
