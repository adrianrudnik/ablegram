<template>
  <div class="NumberInput field">
    <label :for="id">{{ label }}</label>
    <InputNumber
      v-model="value"
      :binary="true"
      class="w-full"
      :class="{ 'p-invalid': !!errorMessage }"
      :inputId="id"
      :name="id"
      v-bind="$attrs"
    />
    <small v-if="props.help">{{ props.help }}</small>
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
  help?: string
}>()

const id = createIdFrom(props.name)

const { errorMessage, value } = useField<number>(toRef(props, 'name'))

defineExpose({
  errorMessage
})
</script>

<style lang="scss">
.NumberInput {
  display: flex;
  flex-direction: column;
  label > small {
    display: block;
    margin-top: 0.5rem;
  }
}
</style>
