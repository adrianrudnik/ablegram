<template>
  <pre>{{ counters }}</pre>
  <!--  <MetaStep>-->
  <!--    <MetaStepItem>-->
  <!--      <template #number>-->
  <!--        {{ validFiles }} /-->
  <!--        <span class="text-red-500" v-tooltip.right="t('metric.invalid_files', invalidFiles)">-->
  <!--          {{ invalidFiles }}-->
  <!--        </span>-->
  <!--      </template>-->
  <!--      <template #label>{{ t('metric.valid_files', validFiles) }}</template>-->
  <!--    </MetaStepItem>-->

  <!--    <MetaStepItem v-for="({ k, v }, idx) in stats" :key="idx">-->
  <!--      <template #number>{{ v }}</template>-->
  <!--      <template #label>{{ t('metric.' + k, v) }}</template>-->
  <!--    </MetaStepItem>-->
  <!--  </MetaStep>-->
</template>

<script setup lang="ts">
import { useStatStore } from '@/stores/stats'
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import MetaStep from '@/components/structure/MetaStep.vue'
import MetaStepItem from '@/components/structure/MetaStepItem.vue'

const { t } = useI18n()

const counters = computed(() => useStatStore().counters)
</script>
