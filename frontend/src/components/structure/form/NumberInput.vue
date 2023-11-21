<template>
  <div class="field">
    <div class="flex align-items-center">
      <label :for="id" class="ml-2">{{ label }}</label>

      <InputNumber
        v-model="value"
        :binary="true"
        :class="{ 'p-invalid': !!errorMessage }"
        :inputId="id"
        :name="id"
        v-bind="$attrs"
      />
    </div>
    <small v-if="errorMessage" :id="`${id}-help`" class="p-error">{{ errorMessage }}</small>
  </div>
</template>

<script lang="ts" setup>
import InputNumber from 'primevue/inputnumber'
import { useField } from 'vee-validate'
import { toRef } from 'vue'
import { createIdFrom } from '@/plugins/id'

const props = defineProps<{
  name: string
  label: string
  value?: number
}>()

const id = createIdFrom(props.name)

const { errorMessage, value } = useField(toRef(props, 'name'))

defineExpose({
  errorMessage
})
</script>
