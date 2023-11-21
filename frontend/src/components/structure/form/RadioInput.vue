<template>
  <div class="flex align-items-center">
    <RadioButton
      v-model="value"
      :inputId="id"
      :name="name"
      :value="radioValue"
      v-bind="$attrs"
      :class="{ 'p-invalid': !!errorMessage }"
    />
    <label :for="id" class="ml-2">{{ label }}</label>
  </div>
</template>

<script setup lang="ts">
import RadioButton from 'primevue/radiobutton'
import { toRef } from 'vue'
import { useField } from 'vee-validate'
import { createIdFrom } from '@/plugins/id'

const props = defineProps<{
  name: string
  label: string
  radioValue?: string
}>()

const id = createIdFrom(props.name + '_' + (props.radioValue?.toString() ?? 'null'))

const { errorMessage, value } = useField(toRef(props, 'name'))

defineExpose({
  errorMessage
})
</script>
