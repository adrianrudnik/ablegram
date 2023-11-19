<template>
  <div class="IndexProp">
    <div class="Label">
      Type: <code class="Name">{{ type }}</code>
    </div>

    <div class="Label">
      Mapping: <code class="Mapping">{{ mapping }}</code>
    </div>

    <div class="Label" v-if="version && version.length > 0">
        Version: <code class="Version" v-for="v in version" :key="v">{{ v }}</code>
    </div>

    <slot />
  </div>
</template>

<script lang="ts" setup>
defineProps<{
  name: string,
  type: string,
  mapping: string,
  version?: string[],
}>()
</script>

<style lang="scss">
.IndexProp {
  .Label {
    display: flex;
    align-items: baseline;
    margin-top: .3rem;
  }

  .Name, .Mapping, .Version {
    margin-left: .5rem;
  }
}
</style>
