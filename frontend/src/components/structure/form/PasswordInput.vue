<template>
  <div class="TextInput field">
    <label :for="id">{{ label }}</label>
    <Password
      v-model="value"
      :binary="true"
      class="w-full"
      :class="{ 'p-invalid': !!errorMessage }"
      :inputId="id"
      toggle-mask
      :name="id"
      :feedback="feedback"
      :input-props="{ ...$attrs, autocomplete: autocomplete }"
    />
    <small v-if="props.help">{{ props.help }}</small>
    <small v-if="errorMessage" :id="`${id}-help`" class="p-error">{{ errorMessage }}</small>
  </div>
</template>

<script lang="ts" setup>
import Password from 'primevue/password'
import { useField } from 'vee-validate'
import { toRef } from 'vue'
import { createIdFrom } from '@/plugins/id'

const props = withDefaults(
  defineProps<{
    name: string
    label: string
    value?: string
    help?: string
    feedback?: boolean
    autocomplete?: string
  }>(),
  {
    feedback: false,
    autocomplete: 'current-password'
  }
)

const id = createIdFrom(props.name)

const { errorMessage, value } = useField<string>(toRef(props, 'name'))

defineExpose({
  errorMessage
})
</script>

<style lang="scss">
.TextInput {
  display: flex;
  flex-direction: column;

  label > small {
    display: block;
    margin-top: 0.5rem;
  }
}
</style>
