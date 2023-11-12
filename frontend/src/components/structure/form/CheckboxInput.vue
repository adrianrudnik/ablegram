<template>
  <div class="field">
    <div class="flex align-items-center">
      <Checkbox
        v-model="value"
        :binary="true"
        :class="{ 'p-invalid': !!errorMessage }"
        :inputId="id"
        :name="id"
        v-bind="$attrs"
      />

      <label :for="id" class="ml-2">{{ label }}</label>
    </div>
    <small v-if="errorMessage" :id="`${id}-help`" class="p-error">{{ errorMessage }}</small>
  </div>
</template>

<script lang="ts" setup>
import Checkbox from 'primevue/checkbox'
import { useField } from 'vee-validate'
import { toRef } from 'vue'

const props = defineProps<{
  name: string
  label: string
}>()

const id = 'f_' + props.name

const { errorMessage, value } = useField(toRef(props, 'name'))

defineExpose({
  errorMessage
})
</script>
