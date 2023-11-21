<template>
  <FormGrid @submit="onFormSubmit" class="mb-3">
    <FormRow>
      <TextInput name="id" :label="t('collector-settings.form.target.filesystem.id.title')" />
      <p>{{ t('collector-settings.form.target.filesystem.id.help') }}</p>
    </FormRow>

    <FormRow>
      <TextInput name="uri" :label="t('collector-settings.form.target.filesystem.uri.title')" />
      <p>{{ t('collector-settings.form.target.filesystem.uri.help') }}</p>
    </FormRow>

    <FormRow>
      <FormRadioGroup :title="t('collector-settings.form.target.filesystem.performance.title')">
        <RadioInput
          name="parser_performance"
          :label="t('collector-settings.form.target.filesystem.performance.value.low')"
          radio-value="low"
        />
        <RadioInput
          name="parser_performance"
          :label="t('collector-settings.form.target.filesystem.performance.value.default')"
          radio-value="default"
        />
        <RadioInput
          name="parser_performance"
          :label="t('collector-settings.form.target.filesystem.performance.value.high')"
          radio-value="high"
        />
      </FormRadioGroup>
    </FormRow>

    <FormRow>
      <NumberInput name="parser_delay" label="P Worker Delay" />
      <p>{{ t('collector-settings.form.target.filesystem.delay.help') }}</p>
    </FormRow>

    <FormRow>
      <CheckboxInput
        name="exclude_system_folders"
        :label="t('behavior-settings.form.open_browser_on_start.title')"
      />
    </FormRow>

    <FormRow>
      <CheckboxInput
        name="exclude_dot_folders"
        :label="t('behavior-settings.form.open_browser_on_start.title')"
      />
    </FormRow>

    <SubmitButton :label="t('button.save')" :loading="isSubmitting" v-model="isSaved" />
  </FormGrid>
</template>

<script setup lang="ts">
import TextInput from '@/components/structure/form/TextInput.vue'
import FormRadioGroup from '@/components/structure/form/FormRadioGroup.vue'
import CheckboxInput from '@/components/structure/form/CheckboxInput.vue'
import NumberInput from '@/components/structure/form/NumberInput.vue'
import FormRow from '@/components/structure/form/FormRow.vue'
import RadioInput from '@/components/structure/form/RadioInput.vue'
import FormGrid from '@/components/structure/form/FormGrid.vue'
import SubmitButton from '@/components/structure/form/SubmitButton.vue'
import { useForm } from 'vee-validate'
import type { CollectorConfig } from '@/stores/config'
import { useConfigStore } from '@/stores/config'
import { boolean, number, object, string } from 'yup'
import { useI18n } from 'vue-i18n'
import { ref } from 'vue'

const { t } = useI18n()

const isSaved = ref<boolean>(false)

const { handleSubmit, isSubmitting, values } = useForm<CollectorConfig>({
  initialValues: useConfigStore().current.collector,
  validationSchema: object().shape({
    id: string()
      .required()
      .matches(/^[\w-]+$/)
      .label(t('collector-settings.form.target.filesystem.id.title')),
    type: string().required(),
    uri: string().required().label(t('collector-settings.form.target.filesystem.uri.title')),
    parser_performance: string()
      .required()
      .label(t('collector-settings.form.target.filesystem.performance.title')),
    parser_delay: number()
      .required()
      .label(t('collector-settings.form.target.filesystem.delay.title')),
    exclude_system_folders: boolean(),
    exclude_dot_folders: boolean()
  })
})

const onFormSubmit = handleSubmit(async (v) => {
  // useConfigStore().current = await fetchApi<Config>('/api/config/behaviour', {
  //   method: 'PUT',
  //   body: JSON.stringify({
  //     autostart_webservice: v.autostart_webservice,
  //     open_browser_on_start: v.open_browser_on_start,
  //     show_service_gui: v.show_service_gui
  //   })
  // })

  isSaved.value = true
})
</script>
