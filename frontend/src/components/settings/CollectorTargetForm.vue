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
      <NumberInput name="parser_delay" label="P Worker Delay" />
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
import { ref } from 'vue'

const { t } = useI18n()

const isSaved = ref<boolean>(false)

const props = withDefaults(
  defineProps<{
    target?: CollectorTargetConfig
  }>(),
  {
    target: () => defaultCollectorTargetConfig()
  }
)

const { handleSubmit, isSubmitting, values } = useForm<CollectorTargetConfig>({
  initialValues: props.target,
  validationSchema: object().shape({
    id: string()
      .required()
      .matches(/^[\w_]+$/)
      .label(t('collector-settings.form.target.filesystem.id.title')),
    type: string().required(),
    uri: string().required().label(t('collector-settings.form.target.filesystem.uri.title')),
    parser_performance: string()
      .required()
      .label(t('collector-settings.form.target.filesystem.performance.title')),
    parser_delay: number()
      .required()
      .label(t('collector-settings.form.target.filesystem.delay.title')),
    exclude_system_folders: boolean().default(true),
    exclude_dot_folders: boolean().default(true)
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
