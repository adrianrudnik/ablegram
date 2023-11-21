<template>
  <div class="field">
    <div class="flex flex-column gap-2">
      <label :for="id">{{ label }}</label>

      <InputText
        v-model="value"
        :binary="true"
        class="w-full"
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
import InputText from 'primevue/inputtext'
import { useField } from 'vee-validate'
import { toRef } from 'vue'
import { createIdFrom } from '@/plugins/id'

const props = defineProps<{
  name: string
  label: string
  value?: string
}>()

const id = createIdFrom(props.name)

const { errorMessage, value } = useField(toRef(props, 'name'))

defineExpose({
  errorMessage
})
</script>
