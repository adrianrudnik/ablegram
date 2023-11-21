<template>
  <FormGrid @submit="onFormSubmit" class="mb-3">
    <FormRow>
      <TextInput
        name="id"
        :label="t('collector-target-form.form.target.filesystem.id.title')"
        :help="t('collector-target-form.form.target.filesystem.id.help')"
      />
    </FormRow>

    <FormRow>
      <TextInput
        name="uri"
        :label="t('collector-target-form.form.target.filesystem.uri.title')"
        :help="t('collector-target-form.form.target.filesystem.uri.help')"
      />
    </FormRow>

    <FormRow>
      <RadioInputGroup
        :title="t('collector-target-form.form.target.filesystem.performance.title')"
        :help="t('collector-target-form.form.target.filesystem.performance.help')"
      >
        <RadioInput
          name="parser_performance"
          :label="t('collector-target-form.form.target.filesystem.performance.value.low.title')"
          :help="t('collector-target-form.form.target.filesystem.performance.value.low.help')"
          radio-value="low"
        />
        <RadioInput
          name="parser_performance"
          :label="t('collector-target-form.form.target.filesystem.performance.value.default.title')"
          :help="t('collector-target-form.form.target.filesystem.performance.value.default.help')"
          radio-value="default"
        />
        <RadioInput
          name="parser_performance"
          :label="t('collector-target-form.form.target.filesystem.performance.value.high.title')"
          :help="t('collector-target-form.form.target.filesystem.performance.value.high.help')"
          radio-value="high"
        />
      </RadioInputGroup>
    </FormRow>

    <FormRow>
      <NumberInput
        name="parser_delay"
        :label="t('collector-target-form.form.target.filesystem.delay.title')"
        :help="t('collector-target-form.form.target.filesystem.delay.help')"
      />
      <p>{{ t('collector-settings.form.target.filesystem.delay.help') }}</p>
    </FormRow>

    <FormRow>
      <CheckboxInput
        name="exclude_system_folders"
        :label="t('collector-target-form.form.target.filesystem.exclude_system_folders.title')"
        :help="t('collector-target-form.form.target.filesystem.exclude_system_folders.help')"
      />
    </FormRow>

    <FormRow>
      <CheckboxInput
        name="exclude_dot_folders"
        :label="t('collector-target-form.form.target.filesystem.exclude_dot_folders.title')"
        :help="t('collector-target-form.form.target.filesystem.exclude_dot_folders.help')"
      />
    </FormRow>

    <SubmitButton :label="t('button.save')" :loading="isSubmitting" v-model="isSaved" />
  </FormGrid>
</template>

<script setup lang="ts">
import TextInput from '@/components/structure/form/TextInput.vue'
import RadioInputGroup from '@/components/structure/form/RadioInputGroup.vue'
import CheckboxInput from '@/components/structure/form/CheckboxInput.vue'
import NumberInput from '@/components/structure/form/NumberInput.vue'
import FormRow from '@/components/structure/form/FormRow.vue'
import RadioInput from '@/components/structure/form/RadioInput.vue'
import FormGrid from '@/components/structure/form/FormGrid.vue'
import SubmitButton from '@/components/structure/form/SubmitButton.vue'
import { useForm } from 'vee-validate'
import type { CollectorConfig, CollectorTargetConfig } from '@/stores/config'
import { useConfigStore, defaultCollectorTargetConfig } from '@/stores/config'
import { boolean, number, object, string } from 'yup'
import { useI18n } from 'vue-i18n'
import { ref, inject, onMounted, toRaw } from 'vue'
import type { Ref } from 'vue'
import { fetchApi } from '@/plugins/api'
import type { DynamicDialogInstance } from 'primevue/dynamicdialogoptions'

const { t } = useI18n()

const configStore = useConfigStore()

const isSaved = ref<boolean>(false)

// Reference injected by PrimeVue dialog service
const dialogRef = inject<Ref<DynamicDialogInstance>>('dialogRef')

// We extract the original ID, in case we rename something.
const id = toRaw(dialogRef?.value.data?.target.id ?? '')

// The given target, or an empty one with defaults.
const target = toRaw(dialogRef?.value.data?.target ?? defaultCollectorTargetConfig())

const { handleSubmit, isSubmitting, values } = useForm<CollectorTargetConfig>({
  initialValues: target,
  validationSchema: object().shape({
    id: string()
      .required()
      .matches(/^[a-zA-Z0-9_-]+$/)
      .label(t('collector-target-form.form.target.filesystem.id.title')),
    uri: string().required().label(t('collector-target-form.form.target.filesystem.uri.title')),
    parser_performance: string()
      .required()
      .label(t('collector-target-form.form.target.filesystem.performance.title')),
    parser_delay: number()
      .required()
      .min(0)
      .max(60 * 1000)
      .label(t('collector-target-form.form.target.filesystem.delay.title')),
    exclude_system_folders: boolean(),
    exclude_dot_folders: boolean()
  })
})

const onFormSubmit = handleSubmit(async (v) => {
  // Store the target in the config store.
  await fetchApi('/api/config/collector/targets/' + v.id, {
    method: 'PUT',
    body: JSON.stringify({
      ...v,
      type: 'filesystem'
    })
  })

  // If the ID is set and is different, we should delete the old source
  if (id !== '' && id !== v.id) {
    await fetchApi('/api/config/collector/targets/' + id, {
      method: 'DELETE'
    })
  }

  await configStore.load()

  isSaved.value = true

  dialogRef?.value.close()
})
</script>
